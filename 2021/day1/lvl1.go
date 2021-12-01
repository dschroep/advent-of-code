package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Solves level 1 of day 1 and returns the result as printable message.
func solveLvl1() string {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		return "Could not open input file. Aborting."
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var prevMeasurement, totalIncreases int
	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return "Could not convert string to integer. Aborting."
		}

		if measurement > prevMeasurement {
			totalIncreases++
		}

		prevMeasurement = measurement
	}

	// Quick hack: The first measurement is not to be interpreted as increase.
	// However, during the first iteration `prevMeasurement == 0` and therefore it will be
	// counted as increase.
	// In order to fix this we will just remove this increase here *smirk*
	totalIncreases--

	return fmt.Sprintf("The measurements contained %d increases.", totalIncreases)
}
