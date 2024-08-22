package main

import (
	"RestDemo/client/requests"
	"RestDemo/models"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	parse, err := time.Parse("02.04.2006 15:4:5", "12.02.2000 00:00:00")
	if err != nil {
		return
	}
	var testStudent = models.Student{
		Id:        "1",
		FirstName: "Jaaaaa",
		LastName:  "Smith",
		Birthdate: parse,
	}

	jsonStudent, err := json.Marshal(&testStudent)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(jsonStudent)
	//err := requests.GetStudent(&client, "4532")
	err = requests.DeleteStudent(&client, jsonStudent)
	err = requests.GetStudent(&client, "1")
	if err != nil {
		fmt.Println(err)
	}
}
