package main

import (
	"crud-nosql/pkg/config"
	"crud-nosql/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	config.Connect()
	defer config.Disconnect()
	routes.RegisterRoutes(r)
	fmt.Println("Starting on PORT 3000")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":3000", r))
}
