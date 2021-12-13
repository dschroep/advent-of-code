package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

const INDEX_FIRST_INSTRUCTION = 862

// Solves level 1 of day 13 and returns the result as printable message.
func solveLvl2() string {
	inputs, err := common.GetFileInput(13)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	// Init `paper`.
	var paper Paper
	for _, input := range inputs {
		if input == "" {
			// This is the part where the folding instructions begin.
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
	for _, foldingInstruction := range inputs[INDEX_FIRST_INSTRUCTION:] {
		along := rune(foldingInstruction[11])
		position, err := strconv.Atoi(strings.Split(foldingInstruction, "=")[1])
		if err != nil {
			return "Could not parse folding position. Aborting."
		}

		paper.fold(along, position)
	}

	return fmt.Sprintf("This is the paper after folding it:\n%s", paper.toString())
}
