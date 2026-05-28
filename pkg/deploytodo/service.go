package deploytodo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/flutterffi/PFGolang/pkg/jobx"
	"github.com/flutterffi/PFGolang/pkg/lockx"
	"github.com/flutterffi/PFGolang/pkg/pagex"
	"github.com/flutterffi/PFGolang/pkg/tenantx"
)

type Item struct {
	ID        int    `json:"id"`
	TenantID  string `json:"tenant_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Service struct {
	items []Item
	jobs  *jobx.Runner
	locks *lockx.Lock
}

func New(jobs *jobx.Runner, locks *lockx.Lock) *Service {
	return &Service{
		items: []Item{
			{ID: 1, TenantID: "team-a", Title: "Prepare rollout", Completed: false},
			{ID: 2, TenantID: "team-a", Title: "Review config", Completed: true},
			{ID: 3, TenantID: "team-b", Title: "Check logs", Completed: false},
		},
		jobs:  jobs,
		locks: locks,
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
	mux.HandleFunc("/maintenance/run", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		s.handleMaintenance(w, r)
	})
	return mux
}

func (s *Service) handleList(w http.ResponseWriter, r *http.Request) {
	tenantID := tenantx.Extract(r)
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))

	filtered := make([]Item, 0)
	for _, item := range s.items {
		if item.TenantID == tenantID {
			filtered = append(filtered, item)
		}
	}

	result := pagex.Slice(filtered, page, pageSize)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}

func (s *Service) handleCreate(w http.ResponseWriter, r *http.Request) {
	tenantID := tenantx.Extract(r)
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "missing title", http.StatusBadRequest)
		return
	}

	item := Item{
		ID:        len(s.items) + 1,
		TenantID:  tenantID,
		Title:     title,
		Completed: false,
	}
	s.items = append(s.items, item)
	s.jobs.Enqueue("index task " + strconv.Itoa(item.ID))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(item)
}

func (s *Service) handleMaintenance(w http.ResponseWriter, r *http.Request) {
	const lockName = "maintenance"
	if !s.locks.Acquire(lockName) {
		http.Error(w, "maintenance already running", http.StatusConflict)
		return
	}
	defer s.locks.Release(lockName)

	results := s.jobs.RunAll()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"results": results,
	})
}
