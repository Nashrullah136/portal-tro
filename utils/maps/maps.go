package maps

func Values[K comparable, T any](instance map[K]T) []T {
	values := make([]T, len(instance))
	index := 0
	for _, val := range instance {
		values[index] = val
		index++
	}
	return values
}
