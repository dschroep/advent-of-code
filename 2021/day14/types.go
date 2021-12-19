package day14

import "strings"

// Representation for "AB -> C", where "A" and "B" are the input, while "C" is the output.
type InsertionRule struct {
	input  string
	output string
}

// Takes an insertion `rule` in the format "AB -> C" and returns an instance of ``InsertionRule`` for it.
func ruleFromString(rule string) InsertionRule {
	return InsertionRule{
		string(rule[0]) + string(rule[1]),
		string(rule[6]),
	}
}

// Applies all `rules` to `polymer` and returns the result.
func apply(polymer string, rules []InsertionRule) string {
	newPolymer := polymer

	for i := 0; i < len(polymer)-1; i++ {
		for _, rule := range rules {
			if strings.Compare(rule.input, polymer[i:i+2]) == 0 {
				// Increment `i` to prevent going over the insertion again and again and again...
				i++
				// Insert `rule.output` at this position.
				newPolymer = newPolymer[:i] + rule.output + newPolymer[i:]
				// No need to check the other rules if this one was applicable.
				break
			}
		}
	}

	return newPolymer
}

// Returns a map[element name]quantity for `polymer`.
func getElementQuantities(polymer string) map[rune]int {
	result := make(map[rune]int)
	for _, element := range polymer {
		_, exists := result[element]
		if exists {
			result[element]++
			continue
		}

		result[element] = 1
	}

	return result
}
