package main

import (
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/cachex"
)

func main() {
	fmt.Println("Lesson 44: Cache Layer")

	cache := cachex.New()
	cache.Set("user:1", "Plato")

	if value, ok := cache.Get("user:1"); ok {
		fmt.Printf("cache hit: %s\n", value)
	}

	if _, ok := cache.Get("user:2"); !ok {
		fmt.Println("cache miss: user:2")
	}
}
