package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/flutterffi/PFGolang/pkg/todoapi"
)

func newMux(store *todoapi.Store) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			store.HandleList(w, r)
		case http.MethodPost:
			store.HandleAdd(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/tasks/complete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		store.HandleComplete(w, r)
	})
	return mux
}

func perform(mux *http.ServeMux, method, path string) {
	request := httptest.NewRequest(method, path, nil)
	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, request)

	fmt.Printf("%s %s -> status=%d\n", method, path, recorder.Code)
	if recorder.Body.Len() > 0 {
		fmt.Printf("body=%s\n", recorder.Body.String())
	}
}

func main() {
	fmt.Println("Lesson 42: Full Todo API Project")

	store := todoapi.NewStore()
	mux := newMux(store)

	perform(mux, http.MethodGet, "/tasks")
	perform(mux, http.MethodPost, "/tasks?title=Practice%20API%20design")
	perform(mux, http.MethodPost, "/tasks/complete?id=2")
	perform(mux, http.MethodGet, "/tasks")

	fmt.Println("Try adding auth, middleware, persistence, or JSON request bodies next.")
}
