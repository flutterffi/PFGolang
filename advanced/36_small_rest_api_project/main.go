package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/flutterffi/PFGolang/pkg/bookapi"
)

func newMux() *http.ServeMux {
	store := bookapi.NewStore()
	mux := http.NewServeMux()
	mux.HandleFunc("/health", bookapi.HealthHandler)
	mux.HandleFunc("/books", store.ListHandler)
	return mux
}

func main() {
	fmt.Println("Lesson 36: Small REST API Project")

	mux := newMux()

	for _, path := range []string{"/health", "/books"} {
		request := httptest.NewRequest(http.MethodGet, path, nil)
		recorder := httptest.NewRecorder()
		mux.ServeHTTP(recorder, request)

		fmt.Printf("GET %s -> status=%d\n", path, recorder.Code)
		fmt.Printf("body=%s\n", recorder.Body.String())
	}

	fmt.Println("Try adding POST, path params, or auth checks as your next upgrade.")
}
