package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "7823146", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Deere"}})
	movies = append(movies, Movie{ID: "2", Isbn: "7829782", Title: "Movie Two", Director: &Director{Firstname: "Bates", Lastname: "Sir"}})
	movies = append(movies, Movie{ID: "3", Isbn: "2452368", Title: "Movie Three", Director: &Director{Firstname: "Steve", Lastname: "O"}})
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Started server at :8080.\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
