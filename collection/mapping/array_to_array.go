package collection

func Transform[T any, R any](collection []*T, transform func(*T) R) []R {
	result := make([]R, 0, len(collection))
	for _, element := range collection {
		result = append(result, transform(element))
	}
	return result
}

func Unique[T comparable](collection []T) []T {
	tmpMap := make(map[T]bool)
	var result []T
	for _, entry := range collection {
		if _, value := tmpMap[entry]; !value {
			tmpMap[entry] = true
			result = append(result, entry)
		}
	}
	return result
}
