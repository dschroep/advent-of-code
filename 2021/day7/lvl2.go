package day7

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 2 of day 7 and returns the result as printable message.
func solveLvl2() string {
	inputs, err := common.GetFileInput(7)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	rawPositions := strings.Split(inputs[0], ",")
	positions := make([]int, 0)
	for _, rawPosition := range rawPositions {
		position, err := strconv.Atoi(rawPosition)
		if err != nil {
			return "Could not parse position. Aborting."
		}

		positions = append(positions, position)
	}

	// The ideal position has to be between the minimum and maximum position.
	minPosition, maxPosition := common.Min(positions), common.Max(positions)
	minConsumption := math.MaxInt
	for testPosition := minPosition; testPosition <= maxPosition; testPosition++ {
		var consumption int
		for _, position := range positions {
			for i := 1; i <= int(math.Abs(float64(position-testPosition))); i++ {
				consumption += i
			}
		}

		if consumption < minConsumption {
			minConsumption = consumption
		}
	}

	return fmt.Sprintf("The least fuel consumption is %d.", minConsumption)
}
