package helpers

import "net/http"

func PrepareResponse(w http.ResponseWriter, received string, expected string) bool {
	w.Header().Set("Content-Type", "application/json")
	if received != expected {
		http.Error(w, "wrong method used", 400)
		return false
	}
	return true
}
