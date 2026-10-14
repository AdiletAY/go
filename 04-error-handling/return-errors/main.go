package main

import (
	"errors"
	"fmt"
	"time"
)

type Base struct {
	CreatedAt time.Time
}

func (b Base) Age() (time.Duration, error) {
	age := time.Since(b.CreatedAt)

	if age < 0 {
		return 0, errors.New("created time is in the future")
	}

	return age, nil
}

func main() {
	b := Base{
		CreatedAt: time.Now().Add(time.Hour),
	}

	age, err := b.Age()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Age:", age)
}
