package day5

import (
	"fmt"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 1 of day 5 and returns the result as printable message.
func solveLvl1() string {
	inputs, err := common.GetFileInput(5)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	var maxX, maxY int
	allLines := make([]Line, 0)
	for _, input := range inputs {
		newLine, err := parseInputLine(input)
		if err != nil {
			return "Input line could not be parsed."
		}

		// Only consider `newLine` if it is horizontal or vertical
		if isDiagonal(newLine) {
			continue
		}

		allLines = append(allLines, newLine)

		// While we're parsing the input, let's figure out how large our coordinate system shall be
		switch {
		case newLine.start.x > maxX:
			maxX = newLine.start.x
		case newLine.end.x > maxX:
			maxX = newLine.end.x
		case newLine.start.y > maxY:
			maxY = newLine.start.y
		case newLine.end.y > maxY:
			maxY = newLine.end.y
		}
	}

	corSys := make([][]byte, maxY+1)
	for row := range corSys {
		corSys[row] = make([]byte, maxX+1)
	}

	for _, line := range allLines {
		corSys = markCoordinates(line, corSys)
	}

	// Now that our coordinate system is filled with values,
	// we only need to figure out how many elements of it are `>= 2`.
	var amountDangPoints int
	for _, row := range corSys {
		for _, dangerLevel := range row {
			if dangerLevel >= 2 {
				amountDangPoints++
			}
		}
	}

	return fmt.Sprintf("The amount of dangerous points is %d.", amountDangPoints)
}
