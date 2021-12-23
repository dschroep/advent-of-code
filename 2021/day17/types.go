package day17

// Probe's vertical acceleration.
const GRAVITY = -1

// TODO: Take from previous days
type Coordinate struct{ x, y int }

type TargetArea struct {
	min Coordinate
	max Coordinate
}

// Returns todays input. Workaround for not being able to have constant structs.
func getInput() TargetArea {
	return TargetArea{
		Coordinate{257, -101},
		Coordinate{286, -57},
	}
}

// Returns the probe's horizontal acceleration, depending on its `xVelocity`.
func getXAcc(xVelocity int) int {
	if xVelocity < 0 {
		return 1
	} else if xVelocity > 0 {
		return -1
	}

	return 0
}
