package main

import "fmt"

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func validate(age int) error {
	if age < 0 {
		return &ValidationError{
			Field: "age",
			Msg:   "must not be negative",
		}
	}

	return nil
}

func main() {
	err := validate(-1)

	if err != nil {
		fmt.Println("Validation failed:", err)
		return
	}

	fmt.Println("It's good")
}
