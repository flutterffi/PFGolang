package main

import (
	"fmt"

	"github.com/flutterffi/PFGolang/pkg/grpcx"
)

func main() {
	fmt.Println("Lesson 51: gRPC Basics Shape")

	response := grpcx.SimulateCreate(grpcx.TaskRequest{
		TenantID: "team-a",
		Title:    "practice rpc shapes",
	})

	fmt.Printf("response: %+v\n", response)
}
