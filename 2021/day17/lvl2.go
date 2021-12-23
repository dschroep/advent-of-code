package day17

import (
	"fmt"
)

// Solves level 2 of day 17 and returns the result as printable message.
// TODO: This works, but at the moment it's only trial-and-error. Is there a more elegant way?
func solveLvl2() string {
	workingVelocities := []Coordinate{}
	for tryX := -300; tryX <= 300; tryX++ {
		for tryY := -300; tryY <= 300; tryY++ {
			hit, _ := shoot(Coordinate{tryX, tryY})
			if hit {
				workingVelocities = append(workingVelocities, Coordinate{tryX, tryY})
			}
		}
	}

	return fmt.Sprintf("There are %d working initial velocities.", len(workingVelocities))
}
