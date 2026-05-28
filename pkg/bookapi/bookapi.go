package bookapi

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Store struct {
	Books []Book
}

func NewStore() Store {
	return Store{
		Books: []Book{
			{ID: 1, Title: "Learning Go", Author: "Team Practice"},
			{ID: 2, Title: "Concurrency Notes", Author: "Plato"},
		},
	}
}

func (s Store) ListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(s.Books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": "book-api",
	})
}
