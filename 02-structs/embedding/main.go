package main

import (
	"fmt"
	"time"
)

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b Base) Age() time.Duration {
	return time.Since(b.CreatedAt)
}

type User struct {
	Base
	Name string
}

func main() {
	u := User{
		Base: Base{
			CreatedAt: time.Now().Add(-24 * time.Hour),
			UpdatedAt: time.Now(),
		},
		Name: "Adilet",
	}

	fmt.Println(u.Name)
	fmt.Println(u.CreatedAt)
	fmt.Println(u.Base.CreatedAt)
	fmt.Println(u.Age())
}
