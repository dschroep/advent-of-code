package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Solves level 1 of day 2 and returns the result as printable message.
func solveLvl1() string {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		return "Could not open input file. Aborting."
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var xPos, yPos int
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		command := input[0]                   // either "up", "down", or "forward"
		amount, err := strconv.Atoi(input[1]) // how far the submarine shall move
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
