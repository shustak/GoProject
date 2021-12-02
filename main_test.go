package main

import (
	"bytes"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.corp.globant.com/a-muliarchik/GoTraining/bookPackage"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

func TestIndex(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	handle := bookPackage.NewBaseHandler(db)
	router := NewRouter(handle)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new request", err)
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	expectedResult := "Welcome to the world of tomorrow!"
	assert.Equal(t, expectedResult, response.Body.String())
}

func TestGetAllBooks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	handle := bookPackage.NewBaseHandler(db)
	router := NewRouter(handle)

	rs := mock.NewRows([]string{"id", "name", "price", "genre", "amount"}).AddRow(125, "TestBook_1", 11.11, 1, 5).AddRow(521, "TestBook_2", 10.10, 2, 1)

	mock.ExpectQuery("SELECT id, name, price, genre, amount FROM books WHERE amount > 0").
		WillDelayFor(time.Second).
		WillReturnRows(rs)

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new request", err)
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	expectedResult := "[{\"id\":125,\"name\":\"TestBook_1\",\"price\":11.11,\"genre\":1,\"amount\":5},{\"id\":521,\"name\":\"TestBook_2\",\"price\":10.1,\"genre\":2,\"amount\":1}]\n"
	assert.Equal(t, expectedResult, response.Body.String())
}

func TestGetBookById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	handle := bookPackage.NewBaseHandler(db)
	router := NewRouter(handle)

	rs := mock.NewRows([]string{"id", "name", "price", "genre", "amount"}).AddRow(15, "TestBook_1", 11.11, 1, 5)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM books WHERE id = ?")).
		WillDelayFor(time.Second).
		WillReturnRows(rs)

	req, err := http.NewRequest("GET", "/books/15", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new request", err)
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	expectedResult := "{\"id\":15,\"name\":\"TestBook_1\",\"price\":11.11,\"genre\":1,\"amount\":5}\n"
	assert.Equal(t, expectedResult, response.Body.String())
}

func TestCreateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	handle := bookPackage.NewBaseHandler(db)
	router := NewRouter(handle)

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO books ( name, price, genre, amount) VALUES (?, ?, ?, ?)")).WillReturnResult(sqlmock.NewResult(555, 1))

	var jsonData = []byte(`{
	  "name":"NewBook",
	  "price": 500.1,
	  "genre": 1,
	  "amount": 5
	}`)

	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new request", err)
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	expectedResult := "555\n"
	assert.Equal(t, expectedResult, response.Body.String())
}

func TestUpdateBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	handle := bookPackage.NewBaseHandler(db)
	router := NewRouter(handle)

	mock.ExpectExec(regexp.QuoteMeta("UPDATE books SET name=?, price=?, genre=?, amount=? WHERE id=?")).WillReturnResult(sqlmock.NewResult(0, 0))

	var jsonData = []byte(`{
		"id": 156,
	  "name":"NewBook",
	  "price": 500.1,
	  "genre": 1,
	  "amount": 5
	}`)
	req, err := http.NewRequest("PUT", "/books/156", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new request", err)
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteBook(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	handle := bookPackage.NewBaseHandler(db)
	router := NewRouter(handle)

	rs := mock.NewRows([]string{"id", "name", "price", "genre", "amount"}).AddRow(351, "TestBook_1", 11.11, 1, 5)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM books WHERE id = ?")).WillReturnRows(rs)
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM books WHERE id = ?")).WillReturnResult(sqlmock.NewResult(0, 0))

	req, err := http.NewRequest("DELETE", "/books/351", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new request", err)
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusNoContent, response.Code)
}

func TestValidateEmptyBookName(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	handle := bookPackage.NewBaseHandler(db)
	router := NewRouter(handle)

	var jsonData = []byte(`{
	  "name":"",
	  "price": 500.1,
	  "genre": 1,
	  "amount": 5
	}`)

	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new request", err)
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusBadRequest, response.Code)

	expectedResult := "Book name can not be empty or more than 100 symbols\n"
	assert.Equal(t, expectedResult, response.Body.String())
}

func TestValidatePriceLessThanZero(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	handle := bookPackage.NewBaseHandler(db)
	router := NewRouter(handle)

	var jsonData = []byte(`{
	  "name":"NewName",
	  "price": -5,
	  "genre": 1,
	  "amount": 5
	}`)

	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new request", err)
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusBadRequest, response.Code)

	expectedResult := "Book price can not be less than zero\n"
	assert.Equal(t, expectedResult, response.Body.String())
}

func TestValidateAmountLessThanZero(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	handle := bookPackage.NewBaseHandler(db)
	router := NewRouter(handle)

	var jsonData = []byte(`{
	  "name":"NewName",
	  "price": 5,
	  "genre": 1,
	  "amount": -10
	}`)

	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new request", err)
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusBadRequest, response.Code)

	expectedResult := "Book amount can not be less than zero\n"
	assert.Equal(t, expectedResult, response.Body.String())
}