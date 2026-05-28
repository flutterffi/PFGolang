package todoservice

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/flutterffi/PFGolang/pkg/cachex"
	"github.com/flutterffi/PFGolang/pkg/flagx"
	"github.com/flutterffi/PFGolang/pkg/queuex"
)

type Item struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Service struct {
	items     []Item
	cache     *cachex.Cache
	flags     flagx.Flags
	publisher queuex.Publisher
}

func New(cache *cachex.Cache, flags flagx.Flags, publisher queuex.Publisher) *Service {
	return &Service{
		items: []Item{
			{ID: 1, Title: "Review handlers", Completed: false},
			{ID: 2, Title: "Measure benchmark output", Completed: true},
		},
		cache:     cache,
		flags:     flags,
		publisher: publisher,
	}
}

func (s *Service) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			s.handleList(w, r)
		case http.MethodPost:
			s.handleCreate(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/tasks/complete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		s.handleComplete(w, r)
	})
	return mux
}

func (s *Service) handleList(w http.ResponseWriter, r *http.Request) {
	if cached, ok := s.cache.Get("tasks:list"); ok {
		w.Header().Set("X-Cache", "hit")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(cached))
		return
	}

	data, _ := json.Marshal(s.items)
	s.cache.Set("tasks:list", string(data))

	w.Header().Set("X-Cache", "miss")
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func (s *Service) handleCreate(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "missing title", http.StatusBadRequest)
		return
	}

	item := Item{
		ID:        len(s.items) + 1,
		Title:     title,
		Completed: false,
	}
	s.items = append(s.items, item)
	s.cache.Delete("tasks:list")

	if s.flags.Enabled("publish_task_events") {
		_ = s.publisher.Publish("tasks.created", title)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(item)
}

func (s *Service) handleComplete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	for i := range s.items {
		if s.items[i].ID == id {
			s.items[i].Completed = true
			s.cache.Delete("tasks:list")
			if s.flags.Enabled("publish_task_events") {
				_ = s.publisher.Publish("tasks.completed", strconv.Itoa(id))
			}
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "task not found", http.StatusNotFound)
}
