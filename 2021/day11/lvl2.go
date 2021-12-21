package day11

import (
	"fmt"

	"github.com/dschroep/advent-of-code/common"
)

const FIELD_SIZE = 100 // = Field of 10 x 10 ``Octopus``

// Solves level 2 of day 11 and returns the result as printable message.
func solveLvl2() string {
	inputs, err := common.GetFileInput(11)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	octopusField := getOctopusField(inputs)

	allFlashStep := 0
	for step := 1; ; step++ {
		flashes := 0
		octopusField, flashes = simulateStep(octopusField)

		if flashes == FIELD_SIZE {
			allFlashStep = step
			break
		}
	}

	return fmt.Sprintf("The first step during which all octopuses flash is %d.", allFlashStep)
}
