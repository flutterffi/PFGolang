package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type response struct {
	Message string `json:"message"`
	Path    string `json:"path"`
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	payload := response{
		Message: "Hello from the practice server.",
		Path:    r.URL.Path,
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	fmt.Println("Lesson 18: HTTP Server Basics")

	request := httptest.NewRequest(http.MethodGet, "/hello", nil)
	recorder := httptest.NewRecorder()

	greetingHandler(recorder, request)

	fmt.Printf("Status: %d\n", recorder.Code)
	fmt.Printf("Body: %s", recorder.Body.String())
	fmt.Println("Replace httptest with http.ListenAndServe later when you want a real local server.")
}
