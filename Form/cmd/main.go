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
}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	err := loadHtml(w, indexPath, messages)
	if err != nil {
		slog.Error(err.Error())
		return
	}
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
		fmt.Println(r.FormValue("message"))
		fmt.Println(r.FormValue("content"))
	default:
		slog.Error("Wrong request method used in createHandle")
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

func appendMessage(title, messageText string) []map[string]string {
	message := make(map[string]string)
	message[title] = messageText
	return append(messages, message)
}

func initMessages() {
	messages = appendMessage("Message One", "Message one content")
	messages = appendMessage("Message Two", "Message two content")
}
