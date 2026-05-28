package tasksvc

import "errors"

var ErrEmptyTitle = errors.New("title cannot be empty")

type Task struct {
	Title string
}

type Repository interface {
	Save(Task) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) AddTask(title string) error {
	if title == "" {
		return ErrEmptyTitle
	}

	return s.repo.Save(Task{Title: title})
}
