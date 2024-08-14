package routes

import (
	"RestDemo/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/students", handlers.GetStudents)
	http.HandleFunc("/students/get", handlers.GetStudent)
	http.HandleFunc("/students/delete", handlers.DeleteStudent)
	http.HandleFunc("/students/create", handlers.CreateStudent)
	http.HandleFunc("/students/update", handlers.UpdateStudent)
}
