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

type Movie struct{
	Id string `json:id`
	Isbn string `json:isbn`
	Title string `json:title`
	Director *Director `json:director`
}

type Director struct{
	Firstname string `json:firstname` 
	Lastname string `json:lastname`
}
var movies[] Movie

func getMovies(w http.ResponseWriter,r* http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter,r* http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item :=range movies{
		if item.Id ==params["id"]{
				movies = append(movies[:index],movies[index+1:]...)
				break
		}	
	}
}

func getMovie(w http.ResponseWriter,r* http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)
	for _,item := range movies{
		if item.Id ==params["id"]{
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovie(w http.ResponseWriter,r* http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter,r* http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item := range movies{
		if item.Id == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}

}
func main(){
	movies = append(movies, Movie{Id: "1", Isbn: "123",Title: "Inception", Director: &Director{Firstname: "kushagra", Lastname: "gupta"}})
	movies = append(movies, Movie{Id: "2",Isbn: "345", Title: "Oppenheimer", Director: &Director{Firstname: "Christopher",Lastname: "Nolan"}})
	r:=mux.NewRouter()
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}