package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 1 of day 13 and returns the result as printable message.
func solveLvl2() string {
	inputs, err := common.GetFileInput(13)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	// Index of the first line that contains a folding instruction.
	var indexFirstInstruction int

	// Init `paper`.
	var paper Paper
	for lineIndex, input := range inputs {
		if input == "" {
			// This is the part where the folding instructions begin.
			indexFirstInstruction = lineIndex + 1
			break
		}

		splitInput := strings.Split(input, ",")
		x, err := strconv.Atoi(splitInput[0])
		if err != nil {
			return "Could not parse X coordinate. Aborting."
		}
		y, err := strconv.Atoi(splitInput[1])
		if err != nil {
			return "Could not parse Y coordinate. Aborting."
		}

		paper = append(paper, Dot{x, y})
	}

	// Fold `paper`.
	for _, foldingInstruction := range inputs[indexFirstInstruction:] {
		equalsPosition := strings.Index(foldingInstruction, "=")

		along := rune(foldingInstruction[equalsPosition-1])
		position, err := strconv.Atoi(foldingInstruction[equalsPosition+1:])
		if err != nil {
			return "Could not parse folding position. Aborting."
		}

		paper.fold(along, position)
	}

	return fmt.Sprintf("This is the paper after folding it:\n%s", paper.toString())
}
