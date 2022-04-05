package main

import (
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

var movie []Movie

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Started server at :8080.\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
