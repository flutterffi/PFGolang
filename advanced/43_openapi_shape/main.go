package main

import (
	"encoding/json"
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/openapix"
)

func main() {
	fmt.Println("Lesson 43: OpenAPI Shape")

	spec := openapix.TodoSpec()
	data, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		fmt.Printf("marshal error: %v\n", err)
		return
	}

	fmt.Println(string(data))
}
