package day11

import (
	"fmt"

	"github.com/dschroep/advent-of-code/common"
)

type Octopus struct {
	// Current energy level
	energyLevel int
	// If this ``Octopus`` flashed during a step. Should be set to `false` after a step.
	flashed bool
}
type OctopusField [10][10]Octopus

// Increases the energy level of all ``Octopus`` in `field` by 1.
func (field *OctopusField) increaseAll() {
	for rowIndex, row := range field {
		for columnIndex := range row {
			field[rowIndex][columnIndex].energyLevel++
		}
	}
}

// Let's the specified ``Octopus`` flash, if possible.
// Will start a chain reaction.
func (field *OctopusField) flash(row, column int) {
	octopus := &field[row][column]

	// Check if `octopus` qualifies for a flash.
	// Does it have the needed energy level and didn't it flash already?
	if octopus.energyLevel < 10 || octopus.flashed {
		return
	}

	// `octopus` qualifies for a flash.
	octopus.flashed = true

	// Increase the energy levels of the surrounding ``Octopus``.
	for neighborRow := -1; neighborRow <= 1; neighborRow++ {
		for neighborCol := -1; neighborCol <= 1; neighborCol++ {
			if neighborRow == 0 && neighborCol == 0 {
				// This is not a neighbor, this is `octopus`.
				continue
			}

			// Check if the selected neighbor even exists.
			if row > 0 && row < 9 && column > 0 && column < 9 {
				neighbor := &field[row+neighborRow][column+neighborCol]
				neighbor.energyLevel++

				field.flash(row+neighborRow, column+neighborCol)
			}
		}
	}
}

// Let's all ``Octopus`` flash, if possible.
// Returns the total amount of flashes.
func (field *OctopusField) flashFlashable() int {
	// Flash all
	for rowIndex, row := range field {
		for columnIndex := range row {
			field.flash(rowIndex, columnIndex)
		}
	}

	// Now count which ``Octopus`` are flashed
	var flashAmount int
	for _, row := range field {
		for _, octopus := range row {
			if octopus.flashed {
				flashAmount++
			}
		}
	}

	return flashAmount
}

// Resets all ``Octopus`` in `field`, that flashed.
func (field *OctopusField) resetFlashed() {
	for rowIndex, row := range field {
		for columnIndex := range row {
			if field[rowIndex][columnIndex].flashed {
				field[rowIndex][columnIndex] = Octopus{0, false}
			}
		}
	}
}

// Simulates one entire step and returns the amount of flashes that occured.
func (field *OctopusField) simulateStep() int {
	field.increaseAll()
	flashAmount := field.flashFlashable()
	field.resetFlashed()

	return flashAmount
}

// Solves level 1 of day 11 and returns the result as printable message.
// FIXME: This does not work at the moment, but I have absolutely no clue why...
func solveLvl1() string {
	inputs, err := common.GetFileInput(11)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	// Init OctopusField.
	octopusField := new(OctopusField)
	for rowIndex, row := range inputs {
		for columnIndex, energyLevel := range row {
			octopusField[rowIndex][columnIndex] = Octopus{int(energyLevel - '0'), false}
		}
	}

	// Simulate the 100 steps.
	var flashAmount int
	for step := 0; step < 100; step++ {
		flashAmount += octopusField.simulateStep()
	}

	return fmt.Sprintf("There were %d flashes after 100 steps.", flashAmount)
}
