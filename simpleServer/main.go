package main

import (
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
)

type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	// 3 routes /, /home, /form
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/home", homeHandle)
	http.HandleFunc("/form", formHandle)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

func formHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("form.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			slog.Error(err.Error())
			return
		}
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
			return
		}
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "%s, %s", name, address)
	}
}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "hello!")
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, nil)

	if err != nil {
		slog.Error(err.Error())
		return
	}
}
