package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Solves level 2 of day 2 and returns the result as printable message.
func solveLvl2() string {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		return "Could not open input file. Aborting."
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var xPos, yPos, aim int
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		command := input[0]                   // either "up", "down", or "forward"
		amount, err := strconv.Atoi(input[1]) // how far the submarine shall move
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
