package day8

import (
	"fmt"
	"math"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Converts `digit` to an integer by looking at `inputLine` (format "<input> | <output>").
// Returns -1 if parsing was not possible.
func toInt(digit string, inputLine string) int {
	digitLength := len(digit)

	// Sort out the obvious digits.
	switch {
	case digitLength == 2:
		return 1
	case digitLength == 4:
		return 4
	case digitLength == 3:
		return 7
	case digitLength == 7:
		return 8
	}

	// For the following approach it is necessary to know
	// which segments result in a 1 and which result in a 4.
	numbers := strings.Split(strings.Join(strings.Split(inputLine, " | "), " "), " ")
	oneSegments := common.FilterSlice(numbers, func(number string) bool {
		return len(number) == 2
	})[0]
	fourSegments := common.FilterSlice(numbers, func(number string) bool {
		return len(number) == 4
	})[0]

	// At this point, the digit can only be 0, 2, 3, 5, 6, or 9.
	// These digits have a length of either 6 or 5.
	if digitLength == 6 {
		// The possible values here are 0, 6, and 9.
		// 9 is the only one of these that matches with all four segments of 4.
		for i, segment := range fourSegments {
			if !strings.ContainsRune(digit, segment) {
				break
			}

			if i == len(fourSegments)-1 {
				return 9
			}
		}

		// 0 matches with both segments of 1 while 6 doesn't.
		for i, segment := range oneSegments {
			if !strings.ContainsRune(digit, segment) {
				break
			}

			if i == len(oneSegments)-1 {
				return 0
			}
		}

		return 6
	} else if digitLength == 5 {
		// The possible values here are 2, 3, and 5.
		// 3 is the only one of these that matches with both segments of 1.
		for i, segment := range oneSegments {
			if !strings.ContainsRune(digit, segment) {
				break
			}

			if i == len(oneSegments)-1 {
				return 3
			}
		}

		// 5 matches with three segments of 4 while 2 only matches with two.
		var matches int
		for _, segment := range fourSegments {
			if strings.ContainsRune(digit, segment) {
				matches++
			}
		}

		if matches == 3 {
			return 5
		}

		return 2
	}

	return -1
}

// Solves level 2 of day 8 and returns the result as printable message.
func solveLvl2() string {
	inputs, err := common.GetFileInput(8)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	var sum int
	for _, input := range inputs {
		outputNumber := strings.Split(input, " | ")[1]

		var parsedNumber int
		for index, outputDigit := range strings.Split(outputNumber, " ") {
			parsedDigit := toInt(outputDigit, input)
			if parsedDigit == -1 {
				return "Could not parse output digit. Aborting."
			}

			parsedNumber += parsedDigit * int(math.Pow10(3-index))
		}

		sum += parsedNumber
	}

	return fmt.Sprintf("The output's sum is %d.", sum)
}
