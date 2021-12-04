package common

import "fmt"

// Formats and prints the results of a `day`.
func OutputResults(day byte, firstResult string, secondResult string) {
	fmt.Printf("[Day %d]\n", day)

	fmt.Printf("\t[Level 1]\n")
	if firstResult != "" {
		fmt.Printf("\t\t%s\n", firstResult)
	} else {
		fmt.Printf("\t\t-- No Result --\n")
	}

	fmt.Printf("\t[Level 2]\n")
	if secondResult != "" {
		fmt.Printf("\t\t%s\n", secondResult)
	} else {
		fmt.Printf("\t\t-- No Result --\n")
	}
}
