package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Name    string   `json:"name"`
	Focus   string   `json:"focus"`
	Topics  []string `json:"topics"`
	Minutes int      `json:"minutes"`
}

func main() {
	fmt.Println("Lesson 16: JSON Basics")

	profile := Profile{
		Name:    "Plato",
		Focus:   "Go practice",
		Topics:  []string{"functions", "structs", "testing"},
		Minutes: 45,
	}

	encoded, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		fmt.Printf("encode error: %v\n", err)
		return
	}

	fmt.Println("Encoded JSON:")
	fmt.Println(string(encoded))

	var decoded Profile
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		fmt.Printf("decode error: %v\n", err)
		return
	}

	fmt.Printf("Decoded struct: %+v\n", decoded)
}
