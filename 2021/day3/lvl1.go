package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Solves level 1 of day 3 and returns the result as printable message.
func solveLvl1() string {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		return "Could not open input file. Aborting."
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Idea of this solution:
	// - Each input row is a binary number of 12 bits -> create slice `values` with size 12
	// - For each '1' we add 1 to the position's value - for '0' we subtract 1 from the position's value
	// - After iterating over every line we either have a positive or negative number for a position
	// -> If a position's value is positive, '1' was more common
	// -> If a position's value is negative, '0' was more common
	values := make([]int, 12)
	for scanner.Scan() {
		inputs := strings.Split(scanner.Text(), "")
		for index, bit := range inputs {
			switch bit {
			case "1":
				values[index]++
			case "0":
				values[index]--
			}
		}
	}

	// Now we have our values, so we can assemble gamma and epsilon rate
	binGammaRate := make([]string, 12)
	binEpsilonRate := make([]string, 12)
	for index, value := range values {
		if value > 0 {
			binGammaRate[index] = "1"
			binEpsilonRate[index] = "0"
		} else {
			binGammaRate[index] = "0"
			binEpsilonRate[index] = "1"
		}
	}

	decGammaRate, err := strconv.ParseInt(strings.Join(binGammaRate, ""), 2, 0)
	if err != nil {
		return "Could not parse binary gamma rate. Aborting."
	}

	decEpsilonRate, err := strconv.ParseInt(strings.Join(binEpsilonRate, ""), 2, 0)
	if err != nil {
		return "Could not parse binary epsilon rate. Aborting."
	}

	return fmt.Sprintf("The power consumption is %d.", decGammaRate*decEpsilonRate)
}
