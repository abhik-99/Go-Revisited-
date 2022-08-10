package main

import (
	"crud-bookstore/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	fmt.Println("Starting on PORT 3000")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":3000", r))

}
