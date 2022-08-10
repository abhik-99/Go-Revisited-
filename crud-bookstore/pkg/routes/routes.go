package routes

import (
	"crud-bookstore/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.DeleteBookById).Methods("DELETE")
	r.HandleFunc("/book/{id}", controllers.UpdateBookById).Methods("PUT")

}
