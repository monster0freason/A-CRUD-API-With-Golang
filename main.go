package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
)


type Movie struct {
	ID       string   `json:"id"`
	ISBN     int      `json:"isbn"`
	Title    string   `json:"title"`
	Director *Director `json:"director"`
}


type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}


var movies []Movie


func getMovies(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(movies)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}


func deleteMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id := params["id"]
    for index, item := range movies {
        if item.ID == id {
            movies = append(movies[:index], movies[index+1:]...)
            json.NewEncoder(w).Encode(movies)
            return
        }
    }
    http.Error(w, "Movie not found", http.StatusNotFound)
}


func getMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id := params["id"]
    for _, item := range movies {
        if item.ID == id {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.Error(w, "Movie not found", http.StatusNotFound)
}


func createMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var movie Movie
    err := json.NewDecoder(r.Body).Decode(&movie)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    movie.ID = strconv.Itoa(rand.Intn(1000000)) 
    movies = append(movies, movie)
    json.NewEncoder(w).Encode(movie)
}


func updateMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id := params["id"]
    for i, item := range movies {
        if item.ID == id {
            movies = append(movies[:i], movies[i+1:]...)
            var movie Movie
            err := json.NewDecoder(r.Body).Decode(&movie)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
            movie.ID = id 
            movies = append(movies, movie)
            json.NewEncoder(w).Encode(movie)
            return
        }
    }
    http.Error(w, "Movie not found", http.StatusNotFound)
}



func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:       "1",
		ISBN:     438227,
		Title:    "Movie 1",
		Director: &Director{FirstName: "John", LastName: "Doe"},
	}, Movie{
		ID:       "2",
		ISBN:     45445,
		Title:    "Movie 2",
		Director: &Director{FirstName: "Steve", LastName: "Smith"},
	})

	
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))

	

}


