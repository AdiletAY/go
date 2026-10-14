package main

import "fmt"

func example() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Caught panic:", r)
		}
	}()

	fmt.Println("Before panic")
	panic("Something went wrong")
	fmt.Println("After panic") // Never runs
}

func main() {
	example()
	fmt.Println("Main continues!")
}
