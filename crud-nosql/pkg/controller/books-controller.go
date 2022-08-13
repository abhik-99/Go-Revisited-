package controller

import (
	"crud-nosql/pkg/models"
	"crud-nosql/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	book := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(*r, book)
	result := book.CreateBook()
	res, _ := json.Marshal(result) //Supposed to return the ID of the inserted book
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	updatedBook := &models.Book{}
	utils.ParseBody(*r, updatedBook)
	oldBook := models.GetBookById(id)

	if updatedBook.Author != nil {
		oldBook.Author = updatedBook.Author
	}
	if updatedBook.Name != "" {
		oldBook.Name = updatedBook.Name
	}
	if updatedBook.Publication != "" {
		oldBook.Publication = updatedBook.Publication
	}

	result := models.UpdateBook(id, oldBook)
	res, _ := json.Marshal(result)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	result := models.DeleteBookById(id)
	res, _ := json.Marshal(result)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
