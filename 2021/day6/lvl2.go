package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 2 of day 6 and returns the result as printable message.
func solveLvl2() string {
	inputs, err := common.GetFileInput(6)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	// "raw fish", haha, no they're "just strings" and not numbers, yet.
	rawFishTimers := strings.Split(inputs[0], ",")
	// Just adapting the solution for level 1 would be too imperformant.
	// We need to come up with something different.
	// The following slice maps the timer's value (= index) to the amount of fish.
	fish := make([]int, 9)
	for _, rawTimer := range rawFishTimers {
		timer, err := strconv.Atoi(rawTimer)
		if err != nil {
			return "Could not parse timer. Aborting."
		}

		fish[timer]++
	}

	for day := 0; day < 256; day++ {
		newFish := make([]int, 9)

		for age := range fish {
			if age == 0 {
				newFish[8] = fish[age]
				newFish[6] = fish[age]
				continue
			}

			// newFish[6]: _Add_ the new generation the the "re-started" ones
			if age == 7 {
				newFish[age-1] += fish[age]
				continue
			}

			newFish[age-1] = fish[age]
		}

		fish = newFish
	}

	var fishAmount int
	for _, amount := range fish {
		fishAmount += amount
	}

	return fmt.Sprintf("There are %d lanternfish after 256 days.", fishAmount)
}
