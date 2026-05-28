package main

import (
	"encoding/json"
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/pagex"
)

func main() {
	fmt.Println("Lesson 49: Pagination Shape")

	items := []string{"a", "b", "c", "d", "e"}
	page := pagex.Slice(items, 2, 2)

	data, err := json.MarshalIndent(page, "", "  ")
	if err != nil {
		fmt.Printf("marshal error: %v\n", err)
		return
	}

	fmt.Println(string(data))
}
