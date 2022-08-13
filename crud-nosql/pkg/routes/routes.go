package routes

import (
	"crud-nosql/pkg/controller"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(r *mux.Router) {
	r.HandleFunc("/book", controller.GetAllBooks).Methods("GET")
	r.HandleFunc("/book/{id}", controller.GetBookById).Methods("GET")
	r.HandleFunc("/book", controller.CreateBook).Methods("POST")
	r.HandleFunc("/book/{id}", controller.UpdateBook).Methods("PATCH")
	r.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")

	r.HandleFunc("/author", controller.GetAllAuthors).Methods("GET")
	r.HandleFunc("/author/{id}", controller.GetAuthorById).Methods("GET")
	r.HandleFunc("/author", controller.CreateAuthor).Methods("POST")
	r.HandleFunc("/author/{id}", controller.UpdateAuthor).Methods("PATCH")
	r.HandleFunc("/author/{id}", controller.DeleteAuthor).Methods("DELETE")

}
