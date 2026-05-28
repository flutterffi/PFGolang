package main

import "fmt"

type DB interface {
	Query(query string, args ...any) ([]map[string]any, error)
}

type fakeDB struct{}

func (fakeDB) Query(query string, args ...any) ([]map[string]any, error) {
	fmt.Printf("query=%q args=%v\n", query, args)
	return []map[string]any{
		{"id": 1, "title": "Study SQL shape"},
		{"id": 2, "title": "Add real driver later"},
	}, nil
}

type TaskRepository struct {
	db DB
}

func NewTaskRepository(db DB) TaskRepository {
	return TaskRepository{db: db}
}

func (r TaskRepository) ListOpenTasks() ([]map[string]any, error) {
	const query = "SELECT id, title FROM tasks WHERE completed = ? ORDER BY id"
	return r.db.Query(query, false)
}

func main() {
	fmt.Println("Lesson 38: SQL Repository Shape")

	repository := NewTaskRepository(fakeDB{})
	rows, err := repository.ListOpenTasks()
	if err != nil {
		fmt.Printf("repository error: %v\n", err)
		return
	}

	fmt.Printf("rows: %+v\n", rows)
}
