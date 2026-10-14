package main

import "fmt"

func example() {
	defer fmt.Println("Cleanup runs")

	fmt.Println("Before panic")
	panic("Something went wrong")
	// Code below panic will not run.
}

func main() {
	example()
	fmt.Println("This will not print")
}
