package jobx

import "fmt"

type Runner struct {
	jobs []string
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Enqueue(name string) {
	r.jobs = append(r.jobs, name)
}

func (r *Runner) RunAll() []string {
	results := make([]string, 0, len(r.jobs))
	for _, job := range r.jobs {
		results = append(results, fmt.Sprintf("ran job: %s", job))
	}
	r.jobs = nil
	return results
}
