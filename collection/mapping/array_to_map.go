package collection

func ToMap[T any, K comparable, V any](collection []*T, getKey func(*T) K, getValue func(*T) V) map[K]V {
	result := make(map[K]V, len(collection))
	for _, element := range collection {
		result[getKey(element)] = getValue(element)
	}
	return result
}

func GroupBy[T any, K comparable](collection []*T, getKey func(*T) K) map[K][]*T {
	result := make(map[K][]*T, len(collection))
	for _, element := range collection {
		groupedElements := append(result[getKey(element)], element)
		result[getKey(element)] = groupedElements
	}
	return result
}
