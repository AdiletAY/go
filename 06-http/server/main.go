package main

import (
	"fmt"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // reads {id} from the matched pattern
	fmt.Println("URL: " + r.URL.String())
	w.Write([]byte("User ID: " + id))
}

func main() {
	http.HandleFunc("/users/{id}", getUser)
	http.ListenAndServe(":8080", nil)
}
