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
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Creating a slice

var movies []Movie

func main() {
	r := mux.NewRouter()


	movies = append(movies, 
		Movie{
		Id: "1",
		Isbn: "1372002",
		Title: "Ravan",
		Director:&Director{
			Firstname: "Prabhas",
			Lastname: "Sonawane",
		}})

	
		movies = append(movies, 
			Movie{
			Id: "2",
			Isbn: "1372003",
			Title: "The Story of Timo",
			Director:&Director{
				Firstname: "Tilu",
				Lastname: "Sonawane",
			}})
	

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at Post 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}

// Res and Req 
func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)

	for index,items := range movies{

		if items.Id==params["id"]{
			movies=append(movies[:index],movies[index+1:]... ) // Pyhton Slice ahai neet bag 
			break

		}
	}
	// Returning the movies 
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)

	for _,items:= range movies{
		if items.Id==params["id"]{
			json.NewEncoder(w).Encode(items)
			return
		}
	}
}

func createMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_= json.NewDecoder(r.Body).Decode(&movie)
	movie.Id=strconv.Itoa(rand.Intn(324601))
	movies=append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter,r *http.Request){

	// set json content type
	w.Header().Set("Content-Type","application-json") 
	//params
	params:= mux.Vars(r)
	// loop over the movies range
	for index,items := range movies{

		if items.Id==params["id"]{
			movies=append(movies[:index],movies[index+1:]... ) // Pyhton Slice ahai neet bag 
			// Adding a movie 
			var movie Movie
			_= json.NewDecoder(r.Body).Decode(&movie)
			movie.Id=params["id"]
			movies=append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}