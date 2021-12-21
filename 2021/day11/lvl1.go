package day11

import (
	"fmt"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 1 of day 11 and returns the result as printable message.
func solveLvl1() string {
	inputs, err := common.GetFileInput(11)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	octopusField := getOctopusField(inputs)

	totalFlashes := 0
	for step := 0; step < 100; step++ {
		newFlashes := 0
		octopusField, newFlashes = simulateStep(octopusField)
		totalFlashes += newFlashes
	}

	return fmt.Sprintf("There were %d flashes after 100 steps.", totalFlashes)
}
