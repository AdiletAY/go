package main

import (
	"fmt"
	"os"
)

func example() {
	defer fmt.Println("First deferred call")
	defer fmt.Println("Second deferred call")

	fmt.Println("Function running")
}

func main() {
	example()

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Could not open file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File opened")
}
