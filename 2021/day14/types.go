package day14

// Representation for "AB -> C", where "AB" is `input`, while "C" is `output`.
type InsertionRule struct {
	input  string
	output string
}

// Defines a single insertion that needs to be applied to a ``Polymer``.
type Insertion struct {
	// Index _after_ which this shall be made.
	at int
	// What shall be inserted.
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
	insertions := []Insertion{}

	// Gather all the `insertions`.
	for elementIndex := 0; elementIndex < len(polymer)-1; elementIndex++ {
		elements := string(polymer[elementIndex : elementIndex+2])

		for _, rule := range rules {
			if elements == rule.input {
				insertions = append(insertions, Insertion{elementIndex, rule.output})

				// Checking all the other rules would be a waste of time at this point...
				break
			}
		}
	}

	// Apply all `insertions` "simultaneously".
	for insertionIndex := len(insertions) - 1; insertionIndex >= 0; insertionIndex-- {
		insertion := insertions[insertionIndex]
		polymer = polymer[:insertion.at+1] + Polymer(insertion.output) + polymer[insertion.at+1:]
	}

	return polymer
}
