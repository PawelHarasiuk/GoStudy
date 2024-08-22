package requests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func GetStudents(client *http.Client) error {
	request, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:8080/students",
		nil,
	)

	if err != nil {
		return err
	}
	do, err := client.Do(request)
	if err != nil {
		return err
	}
	println(do.Status)
	all, err := io.ReadAll(do.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(all))
	return nil
}

func GetStudent(client *http.Client, id string) error {
	url := fmt.Sprintf("http://localhost:8080/students/get?id=%s", id)
	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		return err
	}
	do, err := client.Do(request)
	if err != nil {
		return err
	}
	println(do.Status)
	all, err := io.ReadAll(do.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(all))
	return nil
}

func CreateStudent(client *http.Client, jsonStudent []byte) error {
	url := "http://localhost:8080/students/create"
	body := bytes.NewBuffer(jsonStudent)
	request, err := http.NewRequest(
		http.MethodPost,
		url,
		body,
	)

	if err != nil {
		return err
	}
	do, err := client.Do(request)
	if err != nil {
		return err
	}
	println(do.Status)
	all, err := io.ReadAll(do.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(all))
	return nil
}

func DeleteStudent(client *http.Client, jsonStudent []byte) error {
	url := "http://localhost:8080/students/delete"
	body := bytes.NewBuffer(jsonStudent)

	request, err := http.NewRequest(
		http.MethodDelete,
		url,
		body,
	)

	if err != nil {
		return err
	}
	do, err := client.Do(request)
	if err != nil {
		return err
	}
	println(do.Status)
	all, err := io.ReadAll(do.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(all))
	return nil
}

func UpdateStudent(client *http.Client, jsonStudent []byte) error {
	url := "http://localhost:8080/students/update"
	body := bytes.NewBuffer(jsonStudent)
	request, err := http.NewRequest(
		http.MethodPut,
		url,
		body,
	)

	if err != nil {
		return err
	}
	do, err := client.Do(request)
	if err != nil {
		return err
	}
	println(do.Status)
	all, err := io.ReadAll(do.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(all))
	return nil
}
