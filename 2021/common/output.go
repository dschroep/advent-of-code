package common

import "fmt"

// Formats and prints the results of a `day`.
func OutputResults(day byte, firstResult string, secondResult string) {
	fmt.Printf("[Day %d]\n", day)
	fmt.Printf("\t[Level 1]\n")
	fmt.Printf("\t\t%s\n", firstResult)
	fmt.Printf("\t[Level 2]\n")
	fmt.Printf("\t\t%s\n", secondResult)
}
