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

// Returns the element with the smallest value of `slice`.
func Min(slice []int) int {
	minValue := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < minValue {
			minValue = slice[i]
		}
	}

	return minValue
}

// Returns the element with the greatest value of `slice`.
func Max(slice []int) int {
	maxValue := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > maxValue {
			maxValue = slice[i]
		}
	}

	return maxValue
}
