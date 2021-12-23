package day17

import (
	"fmt"
)

// Solves level 1 of day 17 and returns the result as printable message.
// TODO: This works, but at the moment it's only trial-and-error. Is there a more elegant way?
func solveLvl1() string {
	maxY := 0
	for tryX := 1; tryX <= 100; tryX++ {
		for tryY := 1; tryY <= 100; tryY++ {
			hit, currentMaxY := shoot(Coordinate{tryX, tryY})
			if hit && currentMaxY > maxY {
				maxY = currentMaxY
			}
		}
	}

	return fmt.Sprintf("The highest possible y-position is: %d.", maxY)
}
