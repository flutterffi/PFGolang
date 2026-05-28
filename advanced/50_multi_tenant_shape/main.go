package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/flutterffi/PFGolang/pkg/tenantx"
)

func main() {
	fmt.Println("Lesson 50: Multi-Tenant Shape")

	request := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	request.Header.Set("X-Tenant-ID", "team-a")
	fmt.Printf("tenant from header: %s\n", tenantx.Extract(request))

	fallback := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	fmt.Printf("tenant fallback: %s\n", tenantx.Extract(fallback))
}
