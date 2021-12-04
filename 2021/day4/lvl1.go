package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 1 of day 4 and returns the result as printable message.
func solveLvl1() string {
	inputs, err := common.GetFileInput(4)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	// Retrieve the ``Board``s from our `inputs`.
	// Note that the data for our first ``Board`` starts at line 3 (= index 2).
	allBoards := make([]Board, 0)
	for i := 2; i < len(inputs)-5; i += 6 {
		newBoard, err := retrieveBoard(inputs[i : i+5])
		if err != nil {
			return "Error when initializing the Boards. Aborting."
		}

		allBoards = append(allBoards, newBoard)
	}

	// The first line of `inputs` holds the called out numbers
	var lastCalledOutNumber int
	for _, calledOutNumber := range strings.Split(inputs[0], ",") {
		markNumber(allBoards, calledOutNumber)

		if didWin(allBoards) {
			lastCalledOutNumber, err = strconv.Atoi(calledOutNumber)
			if err != nil {
				return "Could not parse last called out number. Aborting."
			}

			break
		}
	}

	winningBoard, err := getWinningBoard(allBoards)
	if err != nil {
		return "Assumed there was a winning board, but there was none. Aborting."
	}

	sum, err := getUnmarkedSum(winningBoard)
	if err != nil {
		return "Could not parse number of board. Aborting."
	}

	return fmt.Sprintf("The score of the first winning board is %d.", sum*lastCalledOutNumber)
}
