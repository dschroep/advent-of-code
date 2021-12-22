package day14

// Representation for "AB -> C", where "AB" is `input`, while "C" is `output`.
type InsertionRule struct {
	input  string
	output string
}

type Polymer string

// Counts the amount of each element in `polymer` and returns it as `map[element]quantity`.
func (polymer Polymer) quantities() map[rune]int {
	quantities := map[rune]int{}
	for _, element := range polymer {
		quantities[element]++
	}

	return quantities
}

// Applies all `rules` to `polymer` and returns the result.
func (polymer Polymer) applyRules(rules []InsertionRule) Polymer {
	for elementIndex := len(polymer) - 2; elementIndex >= 0; elementIndex-- {
		elements := string(polymer[elementIndex : elementIndex+2])

		for _, rule := range rules {
			if elements == rule.input {
				polymer = polymer[:elementIndex+1] + Polymer(rule.output) + polymer[elementIndex+1:]

				// Checking all the other rules would be a waste of time at this point...
				break
			}
		}
	}

	return polymer
}
