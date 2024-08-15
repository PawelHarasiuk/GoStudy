package handlers

import (
	"RestDemo/models"
	"RestDemo/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

var repository = repositories.NewRepository(repositories.CsvRepository{
	Path: "data/dane.csv",
})
var Students = repository.RepositoryHandler.Load()

func GetStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong request", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&Students)
	if err != nil {
		return
	}
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong request", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	for _, student := range Students {
		if student.Id == id {
			err := json.NewEncoder(w).Encode(student)
			if err != nil {
				return
			}
			return
		}
	}
	http.Error(w, "Student do no exists", 401)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Wrong request", 400)
		return
	}
	defer repository.RepositoryHandler.Save(Students)
	id := r.URL.Query().Get("id")
	for _, student := range Students {
		if student.Id == id {
			delete(Students, student.Id)
			return
		}
	}
	http.Error(w, "No such student", 400)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong request", 400)
		return
	}
	defer repository.RepositoryHandler.Save(Students)

	var newStudent models.Student

	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		fmt.Println(err)
		return
	}
	if _, ok := Students[newStudent.Id]; ok {
		fmt.Println("Student already exists")
		return
	}
	Students[newStudent.Id] = newStudent
	_, err := fmt.Fprintf(w, "Created %v", newStudent)
	if err != nil {
		return
	}
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong request", 400)
		return
	}
	defer repository.RepositoryHandler.Save(Students)

	var newStudent models.Student

	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		fmt.Printf("error decoding json %v\n", err)
		return
	}

	for id, student := range Students {
		if newStudent.Id == id {
			Students[id] = newStudent
			_, err := fmt.Fprintf(w, "updated %v", student)
			if err != nil {
				return
			}
			return
		}
	}

	http.Error(w, "No such student", 400)
}
