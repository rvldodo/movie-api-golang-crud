package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Rating int `json:"rating"`
	Director *Director
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName" `
}

var movies []Movie

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)
	for i, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Title: "Harry Potter", Rating: 8, Director: &Director{FirstName: "J.K", LastName: "Rowling"}})
	movies = append(movies, Movie{ID: "2", Title: "Warrior", Rating: 9, Director: &Director{FirstName: "Tom", LastName: "Hardy"}})
	movies = append(movies, Movie{ID: "3", Title: "Grey Hound", Rating: 8, Director: &Director{FirstName: "Tom", LastName: "Hanks"}})

	r.HandleFunc("/movie", getAllMovies).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server running in port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}