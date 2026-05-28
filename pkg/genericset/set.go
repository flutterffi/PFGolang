package genericset

type Set[T comparable] map[T]struct{}

func New[T comparable](items ...T) Set[T] {
	set := make(Set[T])
	for _, item := range items {
		set.Add(item)
	}

	return set
}

func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) Has(item T) bool {
	_, exists := s[item]
	return exists
}

func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for item := range s {
		values = append(values, item)
	}

	return values
}
