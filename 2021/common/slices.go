package common

// Constructs a new slice that only contains values of `slice` which pass `condition`.
func FilterSlice(slice []string, condition func(string) bool) []string {
	result := make([]string, 0)
	for _, element := range slice {
		if condition(element) {
			result = append(result, element)
		}
	}

	return result
}
