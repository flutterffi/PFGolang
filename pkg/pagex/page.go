package pagex

type Page[T any] struct {
	Items      []T `json:"items"`
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalItems int `json:"total_items"`
}

func Slice[T any](items []T, page, pageSize int) Page[T] {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 1
	}

	start := (page - 1) * pageSize
	if start > len(items) {
		start = len(items)
	}

	end := start + pageSize
	if end > len(items) {
		end = len(items)
	}

	return Page[T]{
		Items:      items[start:end],
		Page:       page,
		PageSize:   pageSize,
		TotalItems: len(items),
	}
}
