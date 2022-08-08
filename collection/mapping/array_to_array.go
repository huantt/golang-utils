package collection

func Transform[T any, R any](collection []*T, transform func(*T) R) []R {
	result := make([]R, 0, len(collection))
	for _, element := range collection {
		result = append(result, transform(element))
	}
	return result
}
