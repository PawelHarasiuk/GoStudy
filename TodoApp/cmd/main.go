package main

import (
	"TodoApp/api"
	"log/slog"
	"net/http"
)

func main() {
	registerRoutes()
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		slog.Error(err.Error())
	}
}

func registerRoutes() {
	http.HandleFunc("/todo", api.GetTasks)
	http.HandleFunc("/todo/create", api.CreateTask)
	http.HandleFunc("/todo/delete", api.DeleteTask)
	http.HandleFunc("/todo/complete", api.CompleteTask)
	http.HandleFunc("/todo/uncompleted", api.UnCompleteTask)
	http.HandleFunc("/todo/update", api.UpdateTask)
}
