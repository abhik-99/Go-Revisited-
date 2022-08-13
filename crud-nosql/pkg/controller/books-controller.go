package controller

import (
	"crud-nosql/pkg/models"
	"crud-nosql/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	if books, err := models.GetAllBooks(); err == nil {
		utils.SendResponse(w, "application/json", http.StatusOK, books)
	} else {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if book, err := models.GetBookById(id); err == nil {
		utils.SendResponse(w, "application/json", http.StatusOK, book)
	} else {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(*r, book)

	if result, err := book.CreateBook(); err == nil {
		utils.SendResponse(w, "application/json", http.StatusOK, result)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	updatedBook := &models.Book{}
	utils.ParseBody(*r, updatedBook)
	if oldBook, err := models.GetBookById(id); err == nil {
		if updatedBook.Author != "" {
			oldBook.Author = updatedBook.Author
		}
		if updatedBook.Name != "" {
			oldBook.Name = updatedBook.Name
		}
		if updatedBook.Publication != "" {
			oldBook.Publication = updatedBook.Publication
		}

		if result, err := models.UpdateBook(id, oldBook); err == nil {
			utils.SendResponse(w, "application/json", http.StatusOK, result)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if result, err := models.DeleteBookById(id); err == nil {
		utils.SendResponse(w, "application/json", http.StatusOK, result)
	} else {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
}
