package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Solves level 2 of day 1 and returns the result as printable message.
func solveLvl2() string {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		return "Could not open input file. Aborting."
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// To make this puzzle a bit more simple, we will load all the measurements into this slice.
	// This gives us the ability to access the measurements via their indices.
	var allMeasurements []int
	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return "Could not convert string to integer. Aborting."
		}

		allMeasurements = append(allMeasurements, measurement)
	}

	// In this part of the puzzle `prevMeasurement` will refer to the previous sum of our sliding window.
	var prevMeasurement, totalIncreases int
	for i := 2; i < len(allMeasurements); i++ {
		// In this part of the puzzle `measurement` will refer to the current sum of our sliding window.
		measurement := allMeasurements[i] + allMeasurements[i-1] + allMeasurements[i-2]
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
