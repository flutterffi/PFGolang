package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/flutterffi/PFGolang/pkg/bookapi"
)

func main() {
	fmt.Println("Lesson 47: Integration Test Shape")

	mux := http.NewServeMux()
	store := bookapi.NewStore()
	mux.HandleFunc("/health", bookapi.HealthHandler)
	mux.HandleFunc("/books", store.ListHandler)

	for _, path := range []string{"/health", "/books"} {
		request := httptest.NewRequest(http.MethodGet, path, nil)
		recorder := httptest.NewRecorder()
		mux.ServeHTTP(recorder, request)
		fmt.Printf("GET %s -> status=%d body=%q\n", path, recorder.Code, recorder.Body.String())
	}

	fmt.Println("This keeps the integration-test flow shape without depending on local sockets.")
}
