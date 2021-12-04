package day1

import (
	"fmt"
	"strconv"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 2 of day 1 and returns the result as printable message.
func solveLvl2() string {
	inputs, err := common.GetFileInput(1)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	// Convert `inputs` to an integer slice.
	// Converting it "on the fly" would also be possible (and probably the more performant way),
	// but it'd also make the code ugly.
	// I guess in this case we want the beautiful code.
	allMeasurements := make([]int, len(inputs))
	for i, input := range inputs {
		measurement, err := strconv.Atoi(input)
		if err != nil {
			return "Could not convert string to integer. Aborting."
		}

		allMeasurements[i] = measurement
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
