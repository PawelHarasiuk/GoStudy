package main

import (
	"RestDemo/routes"
	"fmt"
	"net/http"
)

// TODO
// data base + dependency injection
// dodac mutex
// tests
// poprawic repozytoria ze zamiazst zapisywania wszystkiego naraz robimy pojedyncze query (delete, add, itp)
func main() {
	routes.RegisterRoutes()
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
