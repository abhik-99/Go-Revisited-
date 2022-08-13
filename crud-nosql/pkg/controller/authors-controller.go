package controller

import (
	"crud-nosql/pkg/models"
	"crud-nosql/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {

	if authors, err := models.GetAllAuthors(); err == nil {
		utils.SendResponse(w, "application/json", http.StatusOK, authors)
	} else {
		fmt.Println("Error Occured!", err)
		http.Error(w, err.Error(), http.StatusNotFound)
	}

}

func GetAuthorById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if author, err := models.GetAuthorById(id); err == nil {
		utils.SendResponse(w, "application/json", http.StatusOK, author)
	} else {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if result, err := author.CreateAuthor(); err == nil {
		utils.SendResponse(w, "application/json", http.StatusOK, result)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var updatedAuthor models.Author
	if err := json.NewDecoder(r.Body).Decode(&updatedAuthor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	params := mux.Vars(r)
	id := params["id"]
	if oldAuthor, err := models.GetAuthorById(id); err == nil {
		if updatedAuthor.Firstname != "" {
			oldAuthor.Firstname = updatedAuthor.Firstname
		}
		if updatedAuthor.Lastname != "" {
			oldAuthor.Lastname = updatedAuthor.Lastname
		}
		if result, err := models.UpdateAuthor(oldAuthor); err == nil {
			utils.SendResponse(w, "application/json", http.StatusOK, result)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if result, err := models.DeleteAuthorById(id); err == nil {
		utils.SendResponse(w, "application/json", http.StatusOK, result)
	} else {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
}
