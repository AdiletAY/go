package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

func main() {
	u := User{
		ID:   "123",
		Name: "Adilet",
	}

	data, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data))
}
