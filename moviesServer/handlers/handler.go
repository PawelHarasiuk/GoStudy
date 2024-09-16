package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"moviesServer/types"
	"net/http"
	"strconv"
)

var dbMock = []types.Movie{
	{
		Id:   1,
		Name: "Shining",
	},
	{
		Id:   2,
		Name: "Godfather",
	},
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(dbMock)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	stringId := vars["id"]
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return
	}

	for _, movie := range dbMock {
		if movie.Id == id {
			err := json.NewEncoder(w).Encode(movie)
			if err != nil {
				log.Fatal(err)
				return
			}
			return
		}
	}
	http.Error(w, "movie not found", http.StatusNotFound)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie types.Movie

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		return
	}

	dbMock = append(dbMock, movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie types.Movie

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
		return
	}

	for id, movie := range dbMock {
		if movie.Id == id {
			dbMock[id] = types.Movie{
				Id:   movie.Id,
				Name: movie.Name,
			}
			if err != nil {
				log.Fatal(err)
				return
			}
			return
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie types.Movie

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
		return
	}

	dbMock = append(dbMock, movie)
}
