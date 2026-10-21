package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func lookupUser(id string) (User, bool) {
	if id == "1" {
		return User{ID: "1", Name: "Alice"}, true
	}
	if id == "2" {
		return User{ID: "2", Name: "Bob"}, true
	}
	return User{}, false
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	user, ok := lookupUser(id)
	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	http.HandleFunc("/users/{id}", getUser)
	http.ListenAndServe(":8080", nil)
}
