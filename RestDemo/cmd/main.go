package main

import (
	"RestDemo/routes"
	"fmt"
	"net/http"
)

// TODO
// dodac mutex
// tests
func main() {
	routes.RegisterRoutes()
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
