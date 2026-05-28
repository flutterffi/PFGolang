package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("middleware log: method=%s path=%s duration=%s\n", r.Method, r.URL.Path, time.Since(started).Round(time.Microsecond))
	})
}

func newMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello from routed handler")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "this route explains the practice app")
	})
	return loggingMiddleware(mux)
}

func main() {
	fmt.Println("Lesson 31: Middleware and Routing")

	handler := newMux()
	for _, path := range []string{"/hello", "/about"} {
		request := httptest.NewRequest(http.MethodGet, path, nil)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)
		fmt.Printf("route=%s status=%d body=%q\n", path, recorder.Code, recorder.Body.String())
	}
}
