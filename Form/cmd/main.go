package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
)

var (
	messages   = make([]map[string]string, 0)
	indexPath  = "index.html"
	createPath = "create.html"
)

func main() {
	initMessages()
	registerRouts()
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func registerRouts() {
	// jak zrobic zeby nie bylo konfliktu miedzy sciezka / i /create
	http.HandleFunc("/home", homeHandle)
	http.HandleFunc("/create", createHandle)
	http.HandleFunc("/delete", deleteMessage)
}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	err := loadHtml(w, indexPath, messages)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(r.FormValue("Delete"))
}

func createHandle(w http.ResponseWriter, r *http.Request) {
	err := loadHtml(w, createPath, nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	switch r.Method {
	case http.MethodGet:
		fmt.Println("get")
	case http.MethodPost:
		title := r.FormValue("title")
		content := r.FormValue("content")
		if len(title) == 0 {
			http.Error(w, "Title is empty", http.StatusBadRequest)
		} else if len(content) == 0 {
			http.Error(w, "Content is empty", http.StatusBadRequest)
		} else {
			messages = appendMessage(title, content)
		}
	default:
		slog.Error("Wrong request method used in createHandle")
	}
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	for i, message := range messages {
		if _, ok := message[title]; ok {
			messages = append(messages[:i], messages[i+1:]...)
			w.WriteHeader(http.StatusOK)
			break
		}
	}
}

func loadHtml(w http.ResponseWriter, path string, data any) error {
	files, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	//is it correct?
	if data != nil {
		err = files.ExecuteTemplate(w, path, messages)
	} else {
		err = files.ExecuteTemplate(w, path, nil)
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func appendMessage(title, content string) []map[string]string {
	message := make(map[string]string)
	message[title] = content
	return append(messages, message)
}

func initMessages() {
	messages = appendMessage("MessageOne", "Message one content")
	messages = appendMessage("MessageTwo", "Message two content")
}
