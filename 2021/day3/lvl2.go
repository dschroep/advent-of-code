package day3

import (
	"fmt"
	"strconv"

	"github.com/dschroep/advent-of-code/common"
)

const INPUT_LENGTH = 12

// Returns the most common bit ('0' or '1') in `candidates` for the given `position`.
// Returns '1' if both are equally common.
func getCommonBit(candidates []string, position int) rune {
	var common int
	for _, candidate := range candidates {
		if candidate[position] == '1' {
			common++
		} else {
			common--
		}
	}

	if common >= 0 {
		return '1'
	}

	return '0'
}

// Returns a rating by either filtering the most or least common bits out.
func retrieveRating(candidates []string, filterCommon bool) (int64, error) {
	for position := 0; len(candidates) > 1; position++ {
		commonBit := getCommonBit(candidates, position)
		candidates = common.FilterSlice(candidates, func(candidate string) bool {
			if filterCommon {
				return rune(candidate[position]) == commonBit
			}

			return rune(candidate[position]) != commonBit
		})
	}

	return strconv.ParseInt(candidates[0], 2, 0)
}

// Solves level 2 of day 3 and returns the result as printable message.
func solveLvl2() string {
	inputs, err := common.GetFileInput(3)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	decGenerator, err := retrieveRating(inputs, true)
	if err != nil {
		return "Could not retrieve oxygen generator rating. Aborting."
	}

	decScrubber, err := retrieveRating(inputs, false)
	if err != nil {
		return "Could not retrieve CO2 scrubber rating. Aborting."
	}

	return fmt.Sprintf("The life support rating is %d.", decGenerator*decScrubber)
}
