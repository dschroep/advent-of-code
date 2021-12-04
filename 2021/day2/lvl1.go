package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 1 of day 2 and returns the result as printable message.
func solveLvl1() string {
	inputs, err := common.GetFileInput(2)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	var xPos, yPos int
	for _, input := range inputs {
		splitInput := strings.Split(input, " ")
		command := splitInput[0]                   // either "up", "down", or "forward"
		amount, err := strconv.Atoi(splitInput[1]) // how far the submarine shall move
		if err != nil {
			return "Could not convert string to integer. Aborting."
		}

		switch command {
		case "up":
			yPos -= amount
		case "down":
			yPos += amount
		case "forward":
			xPos += amount
		default:
			return "Encountered unknown command. Aborting."
		}
	}

	return fmt.Sprintf("The positions product is %d.", xPos*yPos)
}
