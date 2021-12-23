package day17

import (
	"fmt"
)

// Simulates a probe shot with the given `velocity`. Returns (hit, maxY).
//   hit: if the probe hit the target area
//   maxY: probe's highest y position during the shot
func shoot(velocity Coordinate) (bool, int) {
	targetArea := getInput()

	probe := Coordinate{0, 0}
	maxY := probe.y

	for {
		// Move `probe`.
		probe.x += velocity.x
		probe.y += velocity.y

		// Update `maxY` if necessary.
		if probe.y > maxY {
			maxY = probe.y
		}

		// Check if we reached our target.
		if probe.x >= targetArea.min.x && probe.x <= targetArea.max.x &&
			probe.y >= targetArea.min.y && probe.y <= targetArea.max.y {
			return true, maxY
		}

		// Check if we overshot.
		if probe.x > targetArea.max.x || probe.y < targetArea.min.y {
			return false, maxY
		}

		// Add the acceleration to `velocity`.
		velocity.x += getXAcc(velocity.x)
		velocity.y += GRAVITY
	}
}

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
