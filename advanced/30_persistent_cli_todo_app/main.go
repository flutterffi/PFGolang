package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/flutterffi/PFGolang/pkg/todo"
)

func loadItems(path string) ([]todo.Item, error) {
	items, err := todo.Load(path)
	if err == nil {
		return items, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return []todo.Item{}, nil
	}

	return nil, err
}

func printItems(items []todo.Item) {
	if len(items) == 0 {
		fmt.Println("No stored items.")
		return
	}

	for _, item := range items {
		status := "pending"
		if item.Completed {
			status = "done"
		}
		fmt.Printf("%d. [%s] %s\n", item.ID, status, item.Title)
	}
}

func main() {
	filePath := flag.String("file", "todo-data.json", "path to the JSON storage file")
	add := flag.String("add", "", "add a task title")
	done := flag.Int("done", 0, "mark an item as completed by id")
	list := flag.Bool("list", false, "list stored tasks")
	flag.Parse()

	items, err := loadItems(*filePath)
	if err != nil {
		fmt.Printf("load error: %v\n", err)
		return
	}

	fmt.Println("Lesson 30: Persistent CLI Todo App")
	fmt.Printf("Storage file: %s\n", *filePath)

	switch {
	case strings.TrimSpace(*add) != "":
		nextID := len(items) + 1
		items = append(items, todo.Item{
			ID:        nextID,
			Title:     strings.TrimSpace(*add),
			Completed: false,
		})
		if err := todo.Save(*filePath, items); err != nil {
			fmt.Printf("save error: %v\n", err)
			return
		}
		fmt.Printf("Added item %d.\n", nextID)
		printItems(items)
	case *done > 0:
		updated := false
		for i := range items {
			if items[i].ID == *done {
				items[i].Completed = true
				updated = true
				break
			}
		}
		if !updated {
			fmt.Printf("Item %d was not found.\n", *done)
			return
		}
		if err := todo.Save(*filePath, items); err != nil {
			fmt.Printf("save error: %v\n", err)
			return
		}
		fmt.Printf("Marked item %d as done.\n", *done)
		printItems(items)
	case *list || flag.NFlag() == 1:
		printItems(items)
	default:
		flag.Usage()
	}
}
