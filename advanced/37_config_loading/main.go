package main

import (
	"fmt"
	"os"

	"github.com/flutterffi/PFGolang/pkg/configx"
)

func main() {
	fmt.Println("Lesson 37: Config Loading")

	os.Setenv("PFGOLANG_APP_NAME", "PracticeServer")
	os.Setenv("PFGOLANG_PORT", "9090")
	os.Setenv("PFGOLANG_DEBUG", "true")

	config := configx.Load()
	fmt.Printf("config: %+v\n", config)
	fmt.Println("Try deleting one env var to see the default value return.")
}
