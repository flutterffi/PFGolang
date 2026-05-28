package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/flutterffi/PFGolang/pkg/todo"
)

func main() {
	fmt.Println("Lesson 23: JSON File Project")

	tempDir, err := os.MkdirTemp("", "pfgolang-json-project-*")
	if err != nil {
		fmt.Printf("temp dir error: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir)

	items := []todo.Item{
		{ID: 1, Title: "Read about generics", Completed: true},
		{ID: 2, Title: "Practice worker pools", Completed: false},
	}

	filePath := filepath.Join(tempDir, "tasks.json")
	if err := todo.Save(filePath, items); err != nil {
		fmt.Printf("save error: %v\n", err)
		return
	}

	loaded, err := todo.Load(filePath)
	if err != nil {
		fmt.Printf("load error: %v\n", err)
		return
	}

	fmt.Printf("JSON file: %s\n", filePath)
	for _, item := range loaded {
		fmt.Printf("%d | %s | completed=%v\n", item.ID, item.Title, item.Completed)
	}
}
