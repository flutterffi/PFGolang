package todoapi

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Item struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Store struct {
	items []Item
}

func NewStore() *Store {
	return &Store{
		items: []Item{
			{ID: 1, Title: "Read Go docs", Completed: true},
			{ID: 2, Title: "Build a handler", Completed: false},
		},
	}
}

func (s *Store) List() []Item {
	copyItems := make([]Item, len(s.items))
	copy(copyItems, s.items)
	return copyItems
}

func (s *Store) Add(title string) Item {
	item := Item{
		ID:        len(s.items) + 1,
		Title:     title,
		Completed: false,
	}
	s.items = append(s.items, item)
	return item
}

func (s *Store) Complete(id int) bool {
	for i := range s.items {
		if s.items[i].ID == id {
			s.items[i].Completed = true
			return true
		}
	}
	return false
}

func (s *Store) HandleList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(s.List())
}

func (s *Store) HandleAdd(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "missing title", http.StatusBadRequest)
		return
	}

	item := s.Add(title)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(item)
}

func (s *Store) HandleComplete(w http.ResponseWriter, r *http.Request) {
	idValue := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idValue)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if !s.Complete(id) {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
