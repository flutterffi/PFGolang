package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/flutterffi/PFGolang/pkg/todo"
)

func printItems(items []todo.Item) {
	if len(items) == 0 {
		fmt.Println("No items yet.")
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
	add := flag.String("add", "", "add a task title")
	done := flag.Int("done", 0, "mark an item as completed by id")
	list := flag.Bool("list", false, "list the current tasks")
	flag.Parse()

	items := []todo.Item{
		{ID: 1, Title: "Practice functions", Completed: true},
		{ID: 2, Title: "Review structs", Completed: false},
		{ID: 3, Title: "Study interfaces", Completed: false},
	}

	fmt.Println("Lesson 24: CLI Todo App")

	switch {
	case strings.TrimSpace(*add) != "":
		nextID := len(items) + 1
		items = append(items, todo.Item{ID: nextID, Title: strings.TrimSpace(*add)})
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
		fmt.Printf("Marked item %d as done.\n", *done)
		printItems(items)
	case *list || flag.NFlag() == 0:
		printItems(items)
	default:
		flag.Usage()
	}
}
