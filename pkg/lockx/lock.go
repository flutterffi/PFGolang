package lockx

type Lock struct {
	held map[string]bool
}

func New() *Lock {
	return &Lock{held: map[string]bool{}}
}

func (l *Lock) Acquire(name string) bool {
	if l.held[name] {
		return false
	}
	l.held[name] = true
	return true
}

func (l *Lock) Release(name string) {
	delete(l.held, name)
}
