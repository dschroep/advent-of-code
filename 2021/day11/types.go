package day11

type Octopus struct {
	// Current energy level of this.
	energyLevel int
	// If this flashed during the current step. Should be reset after every step.
	flashed bool
}

type OctopusField [][]Octopus
