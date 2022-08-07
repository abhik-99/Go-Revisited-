package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Name     string    `json:"name"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.Id == params["id"] {
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.Id = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, newMovie)

	json.NewEncoder(w).Encode(newMovie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for idx, movie := range movies {
		if movie.Id == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conten-type", "application/json")
	params := mux.Vars(r)
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	for idx, movie := range movies {
		if movie.Id == params["id"] {
			newMovie.Id = movie.Id
			movies = append(movies[:idx], movies[idx+1:]...)
			movies = append(movies, newMovie)
		}
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie", addMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Listening on Port 8000!")
	http.ListenAndServe(":8000", r)
}
