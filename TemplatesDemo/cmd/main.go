package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handleExample(w http.ResponseWriter, r *http.Request) {
	fileName := "index.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/login", handleExample)
	http.ListenAndServe("localhost:8080", nil)
}
