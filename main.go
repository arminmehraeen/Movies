package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

type Response struct {
	Message string `json:"message"`
}

var movies []Movie

var idNumber int = 3

func getMovies(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "get list of movie request = GET")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "delete a movie request = DELETE")
	params := mux.Vars(r)

	isFound := false

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			isFound = true
			break
		}
	}

	if !isFound {
		message := Response{
			Message: "seleted movie not found .",
		}
		json.NewEncoder(w).Encode(message)
	} else {
		json.NewEncoder(w).Encode(movies)
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "get a movie request = GET")
	params := mux.Vars(r)

	isFound := false

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			isFound = true
			break
		}
	}

	if !isFound {
		message := Response{
			Message: "seleted movie not found .",
		}
		json.NewEncoder(w).Encode(message)
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "create movie request = POST")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(idNumber)
	movie.Isbn = "000" + strconv.Itoa(idNumber)
	idNumber++
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "update movie request = PUT")
	params := mux.Vars(r)

	isFound := false

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movie.Isbn = item.Isbn
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			isFound = true
			break
		}
	}

	if !isFound {
		message := Response{
			Message: "seleted movie not found .",
		}
		json.NewEncoder(w).Encode(message)
	}
}

func setHeader(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Print(message + "\n")
}

func generateMovie() {
	movies = append(movies,
		Movie{
			ID: "1", Isbn: "0001", Title: "Movie One", Director: &Director{
				Firstname: "Armin", Lastname: "Mehraeen",
			},
		},
		Movie{
			ID: "2", Isbn: "0002", Title: "Movie Two", Director: &Director{
				Firstname: "Armin", Lastname: "Mehraeen",
			},
		})
}

func main() {
	r := mux.NewRouter()
	generateMovie()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("starting server . \n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
