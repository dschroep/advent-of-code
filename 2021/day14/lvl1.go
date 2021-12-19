package day14

import (
	"fmt"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 1 of day 14 and returns the result as printable message.
// FIXME: 2828 seems to be too high and 243 too low...
func solveLvl1() string {
	inputs, err := common.GetFileInput(14)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	polymer := inputs[0]

	rules := make([]InsertionRule, 0)
	for _, insertionRule := range inputs[2:] {
		rules = append(rules, ruleFromString(insertionRule))
	}

	for step := 0; step < 10; step++ {
		polymer = apply(polymer, rules)
	}

	elementQuantities := getElementQuantities(polymer)
	var minQuantity, maxQuantity int
	for _, quantity := range elementQuantities {
		if quantity < minQuantity {
			minQuantity = quantity
		} else if quantity > maxQuantity {
			maxQuantity = quantity
		}
	}

	return fmt.Sprintf("The searched difference is: %d.", maxQuantity-minQuantity)
}
