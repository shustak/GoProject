package bookPackage

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

// Book struct
type Book struct {
	ID	  	int   `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Genre	int   `json:"genre"`
	Amount	int   `json:"amount"`
}

// NewBook return a pointer to a Book
func NewBook() *Book {
	return new(Book)
}

// GetAllBooks return slice of Book that contains all Books in database
func GetAllBooks(db *sql.DB) ([]Book, error) {
	books := make([]Book, 0)

	rows, err := db.Query(`SELECT id, name, price, genre, amount FROM books WHERE amount > 0;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.ID, &book.Name, &book.Price, &book.Genre, &book.Amount)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	log.Println("Books: ", books)
	return books, nil
}

// FindByID return bookPackage by ID
func FindByID(db *sql.DB, ID string) (*Book, error) {
	if ID == "" {
		return nil, errors.New("ID can not be empty")
	}
	id, err := strconv.Atoi(ID)
	if err != nil {
		return nil, errors.New("ID must be a number")
	}
	if id <= 0 {
		return nil, errors.New("ID can not be less or equal 0")
	}

	rows, err := db.Query("SELECT * FROM books WHERE id = ?", ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var book Book

	if rows.Next() {
		if err := rows.Scan(&book.ID, &book.Name, &book.Price, &book.Genre, &book.Amount);
		err != nil {
			return nil, err
		}
		return &book, nil
	}
	if err = rows.Err();
	err != nil {
		return nil, err
	}

	return nil, errors.New("No data found")
}

// Save bookPackage
func (book *Book) Save(db *sql.DB) (int64, error) {

	bookID := int64(0)

	row, err := db.Exec(`INSERT INTO books ( name, price, genre, amount) VALUES (?, ?, ?, ?)`, book.Name, book.Price, book.Genre, book.Amount)
	if err != nil {
		return bookID, err
	}
	bookID, err = row.LastInsertId()
	if err != nil {
		return bookID, err
	}

	return bookID, nil
}

// ValidateBook check bookPackage parameters
func (book *Book) ValidateBook(db *sql.DB) error {
	if book.Name == "" || len(book.Name) >100{
		return errors.New("Book name can not be empty or more than 100 symbols")
	}

	if book.Price < 0 {
		return errors.New("Book price can not be less than zero")
	}

	if book.Amount < 0 {
		return errors.New("Book amount can not be less than zero")
	}

	return nil
}

// Update existing bookPackage
func Update(db *sql.DB, Book Book) error {
	var err error

	_, err = db.Exec(`UPDATE books SET name=?, price=?, genre=?, amount=? WHERE id=?`, Book.Name, Book.Price, Book.Genre, Book.Amount, Book.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete bookPackage from database
func Delete(db *sql.DB, ID string) error {
	var err error

	_, err = db.Exec(`DELETE FROM books WHERE id = ?`, ID)
	if err != nil {
		return err
	}
	return nil
}