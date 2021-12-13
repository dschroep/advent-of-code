package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 1 of day 13 and returns the result as printable message.
func solveLvl1() string {
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

	// Since there's only one instruction for now, we won't care about parsing...
	paper.fold('x', 655)

	return fmt.Sprintf("There are %d dots after just one fold.", paper.countUnique())
}
