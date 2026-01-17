package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pallinder/go-randomdata"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

func (movieData Movie) printMovies() {

	fmt.Println(movieData.ID)
	fmt.Println(movieData.Isbn)
	fmt.Println(movieData.Title)
	fmt.Println(movieData.Director.FirstName)
	fmt.Println(movieData.Director.LastName)
	fmt.Println("-----------------------------------")
}

func NewMovie(movieID, movieIsbn, movieTitle, directorFirstName, directorLastName string) *Movie {
	return &Movie{
		ID:    movieID,
		Isbn:  movieIsbn,
		Title: movieTitle,
		Director: &Director{
			FirstName: directorFirstName,
			LastName:  directorLastName,
		},
	}
}

func addRandomMovies(count int) {

	for count != 0 {
		var dataMovie *Movie
		dataMovie = NewMovie(randomdata.Digits(3), randomdata.Digits(5), randomdata.SillyName(), randomdata.FirstName(2), randomdata.LastName())
		dataMovie.printMovies()
		movies = append(movies, *dataMovie)
		count--
	}

}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	//create a router using mux
	r := mux.NewRouter()

	//creating some random movies to get started with
	addRandomMovies(5)

	//define all the paths and the functions that will run when those paths are called as per the METHOD
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	// //starting the server
	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMovie *Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	movies = append(movies, *newMovie)
	json.NewEncoder(w).Encode(newMovie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var newMovie *Movie
			_ = json.NewDecoder(r.Body).Decode(&newMovie)
			movies = append(movies, *newMovie)
			json.NewEncoder(w).Encode(newMovie)
			return
		}
	}

}
