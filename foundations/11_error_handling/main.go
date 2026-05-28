package main

import (
	"errors"
	"fmt"
)

func validateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}

	if age < 18 {
		return fmt.Errorf("age %d is below the required minimum", age)
	}

	return nil
}

func main() {
	fmt.Println("Lesson 11: Error Handling")

	for _, age := range []int{12, 21, -1} {
		if err := validateAge(age); err != nil {
			fmt.Printf("Rejected: %v\n", err)
			continue
		}

		fmt.Printf("Accepted age: %d\n", age)
	}
}
