package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func makeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	makeItSpeak(Dog{})
	makeItSpeak(Cat{})
}
