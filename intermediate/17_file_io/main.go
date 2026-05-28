package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Lesson 17: File I/O")

	tempDir, err := os.MkdirTemp("", "pfgolang-file-io-*")
	if err != nil {
		fmt.Printf("create temp dir error: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, "notes.txt")
	content := []byte("Go grows through repetition.\nPractice, inspect, refine.\n")

	if err := os.WriteFile(filePath, content, 0o644); err != nil {
		fmt.Printf("write error: %v\n", err)
		return
	}

	readBack, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read error: %v\n", err)
		return
	}

	fmt.Printf("Temporary file: %s\n", filePath)
	fmt.Println("Content:")
	fmt.Print(string(readBack))
}
