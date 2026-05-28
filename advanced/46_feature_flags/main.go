package main

import (
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/flagx"
)

func main() {
	fmt.Println("Lesson 46: Feature Flags")

	flags := flagx.New(map[string]bool{
		"publish_task_events": true,
		"beta_search":         false,
	})

	fmt.Printf("publish_task_events enabled: %v\n", flags.Enabled("publish_task_events"))
	fmt.Printf("beta_search enabled: %v\n", flags.Enabled("beta_search"))
}
