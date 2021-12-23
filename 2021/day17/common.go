package day17

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
