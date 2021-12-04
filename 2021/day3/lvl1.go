package day3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 1 of day 3 and returns the result as printable message.
func solveLvl1() string {
	inputs, err := common.GetFileInput(3)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	// Idea of this solution:
	// - Each input row is a binary number of 12 bits -> create slice `commons` with size 12
	// - For each '1' we add 1 to the position's "common" - for '0' we subtract 1
	// - After iterating over every line we either have a positive or negative "common" for a position
	// -> If a position's "common" is positive, '1' was more common
	// -> If a position's "common" is negative, '0' was more common
	commons := make([]int, 12)
	for _, input := range inputs {
		splitInput := strings.Split(input, "")

		for index, bit := range splitInput {
			if bit == "1" {
				commons[index]++
			} else {
				commons[index]--
			}
		}
	}

	// Now we have our values, so we can assemble gamma and epsilon rate
	var binGammaRate, binEpsilonRate string
	for _, value := range commons {
		if value > 0 {
			binGammaRate += "1"
			binEpsilonRate += "0"
		} else {
			binGammaRate += "0"
			binEpsilonRate += "1"
		}
	}

	decGammaRate, err := strconv.ParseInt(binGammaRate, 2, 0)
	if err != nil {
		return "Could not parse binary gamma rate. Aborting."
	}

	decEpsilonRate, err := strconv.ParseInt(binEpsilonRate, 2, 0)
	if err != nil {
		return "Could not parse binary epsilon rate. Aborting."
	}

	return fmt.Sprintf("The power consumption is %d.", decGammaRate*decEpsilonRate)
}
