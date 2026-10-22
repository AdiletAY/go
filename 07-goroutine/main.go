package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello world from goroutine!"
	}()

	msg := <-ch
	fmt.Println(msg)
}
