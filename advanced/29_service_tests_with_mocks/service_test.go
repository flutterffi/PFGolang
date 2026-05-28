package main

import (
	"errors"
	"testing"

	"github.com/flutterffi/PFGolang/pkg/tasksvc"
)

type fakeRepo struct {
	saved []tasksvc.Task
	err   error
}

func (f *fakeRepo) Save(task tasksvc.Task) error {
	if f.err != nil {
		return f.err
	}

	f.saved = append(f.saved, task)
	return nil
}

func TestAddTaskSavesTask(t *testing.T) {
	repo := &fakeRepo{}
	service := tasksvc.New(repo)

	if err := service.AddTask("practice mocks"); err != nil {
		t.Fatalf("AddTask returned error: %v", err)
	}

	if len(repo.saved) != 1 {
		t.Fatalf("expected 1 saved task, got %d", len(repo.saved))
	}

	if repo.saved[0].Title != "practice mocks" {
		t.Fatalf("saved title = %q, want %q", repo.saved[0].Title, "practice mocks")
	}
}

func TestAddTaskRejectsEmptyTitle(t *testing.T) {
	repo := &fakeRepo{}
	service := tasksvc.New(repo)

	err := service.AddTask("")
	if !errors.Is(err, tasksvc.ErrEmptyTitle) {
		t.Fatalf("expected ErrEmptyTitle, got %v", err)
	}
}

func TestAddTaskPropagatesRepositoryError(t *testing.T) {
	repo := &fakeRepo{err: errors.New("disk unavailable")}
	service := tasksvc.New(repo)

	err := service.AddTask("practice services")
	if err == nil {
		t.Fatal("expected repository error")
	}

	if err.Error() != "disk unavailable" {
		t.Fatalf("unexpected error: %v", err)
	}
}
