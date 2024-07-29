package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/moaz/go-BookStoreProject/pkg/models"
	"github.com/moaz/go-BookStoreProject/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	NewBook := models.GetBooks()
	res, _ := json.Marshal(NewBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BookId := vars["bookId"]
	id, err := strconv.ParseInt(BookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	Book, _ := models.GetBooksbyID(id)
	res, _ := json.Marshal(Book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var createBook = &models.Book{}
	utils.BodyParse(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var DeletedBook models.Book
	vars := mux.Vars(r)
	bookID := vars["bookId"]
	newID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	DeletedBook = models.DeleteBook(newID)
	res, _ := json.Marshal(DeletedBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.BodyParse(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	newId, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book, db := models.GetBooksbyID(newId)
	if book.Name != "" {
		book.Name = updateBook.Name
	}
	if book.Author != "" {
		book.Author = updateBook.Author
	}
	if book.Publication != "" {
		book.Publication = updateBook.Publication
	}
	db.Save(&book)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
