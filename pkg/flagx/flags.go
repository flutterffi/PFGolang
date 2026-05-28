package flagx

type Flags struct {
	values map[string]bool
}

func New(pairs map[string]bool) Flags {
	copyMap := make(map[string]bool, len(pairs))
	for key, value := range pairs {
		copyMap[key] = value
	}
	return Flags{values: copyMap}
}

func (f Flags) Enabled(name string) bool {
	return f.values[name]
}
