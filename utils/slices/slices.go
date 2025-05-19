package slices

// Comment
func Map[T any, R any](items []T, callback func(item T) R) []R {
	mapped := []R{}

	for _, item := range items {
		mapped = append(mapped, callback(item))
	}

	return mapped
}
