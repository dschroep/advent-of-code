package day14

import (
	"fmt"
	"math"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 1 of day 14 and returns the result as printable message.
func solveLvl1() string {
	inputs, err := common.GetFileInput(14)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	polymer := Polymer(inputs[0])

	rules := make([]InsertionRule, len(inputs[2:]))
	for ruleIndex, rule := range inputs[2:] {
		rules[ruleIndex] = InsertionRule{rule[:2], rule[6:]}
	}

	for step := 0; step < 10; step++ {
		polymer = polymer.applyRules(rules)
	}

	quantities := polymer.quantities()
	// Note: The values for these variables are the right way around!
	minQuantity, maxQuantity := math.MaxInt, 0
	for _, quantity := range quantities {
		if quantity < minQuantity {
			minQuantity = quantity
		} else if quantity > maxQuantity {
			maxQuantity = quantity
		}
	}

	return fmt.Sprintf("The searched difference is: %d.", maxQuantity-minQuantity)
}
