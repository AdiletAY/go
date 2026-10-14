package main

import "fmt"

type Counter struct {
	Count int
}

func increment(c *Counter) {
	c.Count++
}

func main() {
	c := Counter{Count: 0}

	increment(&c) // Pass the address of c.

	fmt.Println(c.Count) // 1
}
