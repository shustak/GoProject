package bookPackage

import (
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type BaseHandler struct {
	db *sql.DB
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

// Index Handler
func (h *BaseHandler)Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	_, err := w.Write([]byte("Welcome to the world of tomorrow!"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetAllBooks return all books from database
func (h *BaseHandler)GetAllBooks(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	//books = make([]Book, 0)
	books, err := GetAllBooks(h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetBookByID find and return bookPackage from database by id
func (h *BaseHandler)GetBookByID(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := params.ByName("id")

	book, err := FindByID(h.db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(book.ID)
	if book.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//CreateBook create a bookPackage
func (h *BaseHandler)CreateBook(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	body := req.Body
	book := NewBook()

	err := json.NewDecoder(body).Decode(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer body.Close()

	err = book.ValidateBook(h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookID, err := book.Save(h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateBook update an existing bookPackage
func (h *BaseHandler)UpdateBook(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ID := params.ByName("id")
	if ID == "" {
		http.Error(w, "ID can not be empty", http.StatusBadRequest)
		return
	}
	var book Book
	err := json.NewDecoder(req.Body).Decode(&book)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = book.ValidateBook(h.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Update(h.db, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteBook delete a bookPackage from database
func (h *BaseHandler)DeleteBook(w http.ResponseWriter, req *http.Request, prm httprouter.Params) {
	ID := prm.ByName("id")

	book, err := FindByID(h.db, ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if book.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = Delete(h.db, ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}