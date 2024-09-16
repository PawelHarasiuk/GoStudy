package main

import (
	"github.com/gorilla/mux"
	"log"
	"moviesServer/handlers"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	registerRoutes(router)
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func registerRoutes(router *mux.Router) {
	router.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	router.HandleFunc("/movies/create", handlers.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/update", handlers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/delete", handlers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/movies/{id}", handlers.GetMovie).Methods("GET")
}
