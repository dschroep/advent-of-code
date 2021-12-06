package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 1 of day 6 and returns the result as printable message.
func solveLvl1() string {
	inputs, err := common.GetFileInput(6)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	// "raw fish", haha, no they're "just strings" and not numbers, yet.
	rawFishTimers := strings.Split(inputs[0], ",")
	fishTimers := make([]int, 0)
	for _, rawTimer := range rawFishTimers {
		timer, err := strconv.Atoi(rawTimer)
		if err != nil {
			return "Could not parse timer. Aborting."
		}

		fishTimers = append(fishTimers, timer)
	}

	for i := 0; i < 80; i++ {
		for i := range fishTimers {
			if fishTimers[i] > 0 {
				fishTimers[i]--
			} else {
				fishTimers[i] = 6
				fishTimers = append(fishTimers, 8)
			}
		}
	}

	return fmt.Sprintf("There are %d lanternfish after 80 days.", len(fishTimers))
}
