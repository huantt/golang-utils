package collection

func FindOne[T any](collection []*T, condition func(element *T) bool) *T {
	for _, element := range collection {
		if condition(element) {
			return element
		}
	}
	return nil
}

func FindAll[T any](collection []*T, condition func(element *T) bool) []*T {
	var listResult []*T
	for _, element := range collection {
		if condition(element) {
			listResult = append(listResult, element)
		}
	}
	return listResult
}
