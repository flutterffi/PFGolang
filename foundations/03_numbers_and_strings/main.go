package main

import "fmt"

func main() {
	price := 19.95
	quantity := 3
	total := price * float64(quantity)

	name := "Go"
	version := 1.26

	fmt.Println("Lesson 03: Numbers and Strings")
	fmt.Printf("Buying %d books at %.2f each costs %.2f.\n", quantity, price, total)
	fmt.Printf("%s %.2f is strongly typed and easy to format.\n", name, version)
}
