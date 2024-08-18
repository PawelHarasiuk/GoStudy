package handlers

import (
	"RestDemo/helpers"
	"RestDemo/models"
	"RestDemo/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

//var repository = repositories.NewRepository(repositories.CsvRepository{
//	Path: "data/dane.csv",
//})

var repository = repositories.NewRepository(repositories.PostgresRepository{
	ConnString: "postgresql://pawel:passwd@localhost/student?sslmode=disable",
})

func GetStudents(w http.ResponseWriter, r *http.Request) {
	if !helpers.PrepareResponse(w, http.MethodGet, r.Method) {
		return
	}
	students, err := repository.RepositoryHandler.GetStudents()
	if err != nil {
		mess := fmt.Sprintf("Error in repository: %v", err)
		http.Error(w, mess, 500)
		return
	}
	err = json.NewEncoder(w).Encode(&students)
	if err != nil {
		http.Error(w, "Error parsing student", 404)
		return
	}
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	if !helpers.PrepareResponse(w, http.MethodGet, r.Method) {
		return
	}
	students, err := repository.RepositoryHandler.GetStudents()
	if err != nil {
		mess := fmt.Sprintf("Error in repository: %v", err)
		http.Error(w, mess, 500)
		return
	}
	seen := false
	var foundStudent models.Student
	id := r.URL.Query().Get("id")
	for _, student := range students {
		if id == student.Id {
			foundStudent = student
			seen = true
			break
		}
	}
	if !seen {
		http.Error(w, "student not found", 500)
		return
	}
	err = json.NewEncoder(w).Encode(&foundStudent)
	if err != nil {
		http.Error(w, "Error parsing student", 404)
		return
	}
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	if !helpers.PrepareResponse(w, http.MethodDelete, r.Method) {
		return
	}
	var oldStudent models.Student
	err := json.NewDecoder(r.Body).Decode(&oldStudent)
	if err != nil {
		http.Error(w, "Error reading student", 404)
		return
	}
	err = repository.RepositoryHandler.DeleteStudent(oldStudent)
	if err != nil {
		mess := fmt.Sprintf("Error in repository: %v", err)
		http.Error(w, mess, 500)
		return
	}
	err = json.NewEncoder(w).Encode(&oldStudent)
	if err != nil {
		http.Error(w, "Error parsing student", 404)
		return
	}
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	if !helpers.PrepareResponse(w, http.MethodPost, r.Method) {
		return
	}
	var newStudent models.Student
	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		http.Error(w, "Error reading student", 404)
		return
	}
	err = repository.RepositoryHandler.CreateStudent(newStudent)
	if err != nil {
		mess := fmt.Sprintf("Error in repository: %v", err)
		http.Error(w, mess, 500)
		return
	}
	err = json.NewEncoder(w).Encode(&newStudent)
	if err != nil {
		http.Error(w, "Error parsing student", 404)
		return
	}
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Used wrong http method", 400)
		return
	}
	var newStudent models.Student
	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		http.Error(w, "Error reading student", 404)
		return
	}
	err = repository.RepositoryHandler.UpdateStudent(newStudent)
	if err != nil {
		mess := fmt.Sprintf("Error in repository: %v", err)
		http.Error(w, mess, 500)
		return
	}
	err = json.NewEncoder(w).Encode(&newStudent)
	if err != nil {
		http.Error(w, "Error parsing student", 404)
		return
	}
}
