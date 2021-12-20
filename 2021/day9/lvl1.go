package day9

import (
	"fmt"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 1 of day 9 and returns the result as printable message.
func solveLvl1() string {
	inputs, err := common.GetFileInput(9)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	var riskLevelSum int
	for rowIndex, row := range inputs {
		for columnIndex := range row {
			height, neighbors := getHeights(columnIndex, rowIndex, inputs)

			if isLowPoint(height, neighbors) {
				riskLevelSum += height + 1
			}
		}
	}

	return fmt.Sprintf("The riks level sum is %d.", riskLevelSum)
}
