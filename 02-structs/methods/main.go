package main

import "fmt"

type Counter struct {
	count int
}

// Value receiver: changes a copy.
func (c Counter) IncrementCopy() {
	c.count++
}

// Pointer receiver: changes the original.
func (c *Counter) Increment() {
	c.count++
}

type Calculator struct {
	value int
}

func (c *Calculator) Add(x int) {
	c.value += x
}

func (u User) Greeting() string {
	return "Hello, " + u.Name
}

type User struct {
	Name string
}

func main() {
	counter := Counter{}
	counter.IncrementCopy()
	fmt.Println(counter.count) // 0

	counter.Increment()
	fmt.Println(counter.count) // 1

	calc := Calculator{value: 10}
	calc.Add(5)
	fmt.Println(calc.value) // 15

	user := User{Name: "Adilet"}
	fmt.Println(user.Greeting())
}
