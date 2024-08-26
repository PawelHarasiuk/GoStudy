package api

import (
	"TodoApp/repositories"
	"TodoApp/types"
	"encoding/json"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"
)

var (
	rep = repositories.PostgresRepository{
		DriverName: "postgres",
		ConnString: "postgresql://pawel:passwd@localhost/todo?sslmode=disable",
	}
	indexPath = "index.html"
)

// is it correct to mix endpoint with displaying template????
func GetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		slog.Error("Wrong request method")
		return
	}
	tmpl := template.Must(template.ParseFiles(indexPath))
	todos, err := rep.GetTasks()
	if err != nil {
		slog.Error(err.Error())
	}
	err = tmpl.Execute(w, todos)
	if err != nil {
		slog.Error(err.Error())
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		slog.Error("Wrong request method")
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("Wrong id param in query")
		return
	}
	err = rep.DeleteTask(id)
	if err != nil {
		slog.Error("Repository error: ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		slog.Error("Wrong request method")
		return
	}
	var newTask types.Todo
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		slog.Error("Error parsing body: ", err)
		return
	}
	err = rep.CreateTask(&newTask)
	if err != nil {
		slog.Error("Repository error: ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func CompleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		slog.Error("Wrong request method")
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("Wrong id param in query")
		return
	}
	err = rep.CompleteTask(id)
	if err != nil {
		slog.Error("Repository error: ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UnCompleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		slog.Error("Wrong request method")
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("Wrong id param in query")
		return
	}
	err = rep.UnCompleteTask(id)
	if err != nil {
		slog.Error("Repository error: ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
