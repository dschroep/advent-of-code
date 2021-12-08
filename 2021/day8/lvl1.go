package day8

import (
	"fmt"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 1 of day 8 and returns the result as printable message.
func solveLvl1() string {
	inputs, err := common.GetFileInput(8)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	outputDigits := make([]string, 0)
	for _, input := range inputs {
		outputNumber := strings.Split(input, " | ")[1]
		outputDigits = append(outputDigits, strings.Split(outputNumber, " ")...)
	}

	var amount int
	for _, digit := range outputDigits {
		digitLength := len(digit)

		// display '1'
		if digitLength == 2 ||
			// display '4'
			digitLength == 4 ||
			// display '7'
			digitLength == 3 ||
			// display '8'
			digitLength == 7 {
			amount++
		}
	}

	return fmt.Sprintf("The outputs contained '1', '4', '7', and '8' %d times.", amount)
}
