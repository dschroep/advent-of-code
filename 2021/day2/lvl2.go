package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 2 of day 2 and returns the result as printable message.
func solveLvl2() string {
	inputs, err := common.GetFileInput(2)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	var xPos, yPos, aim int
	for _, input := range inputs {
		splitInput := strings.Split(input, " ")
		command := splitInput[0]                   // either "up", "down", or "forward"
		amount, err := strconv.Atoi(splitInput[1]) // how far the submarine shall move
		if err != nil {
			return "Could not convert string to integer. Aborting."
		}

		switch command {
		case "up":
			aim -= amount
		case "down":
			aim += amount
		case "forward":
			xPos += amount
			yPos += aim * amount
		default:
			return "Encountered unknown command. Aborting."
		}
	}

	return fmt.Sprintf("The positions product is %d.", xPos*yPos)
}
