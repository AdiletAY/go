package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	u := User{
		Name:  "Adilet",
		Age:   25,
		Email: "adilet@example.com",
	}

	fmt.Println(u)
	fmt.Println(u.Name)
	fmt.Println(u.Age)
}
