package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 2 of day 3 and returns the result as printable message.
func solveLvl2() string {
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

	// The first line of `inputs` holds the called out numbers.
	winEvents := []WinEvent{}
	for _, calledOutNumber := range strings.Split(inputs[0], ",") {
		markNumber(allBoards, calledOutNumber)

		winBoardsIndices := getWinningBoards(allBoards)
		if len(winBoardsIndices) > 0 {
			lastCalledOutNumber, err := strconv.Atoi(calledOutNumber)
			if err != nil {
				return "Could not parse last called out number. Aborting."
			}

			// It is important that we go backwards through `winBoardsIndices`, as it is ordered from smallest to greatest index.
			// This makes the board removal process way easier, because this way the other indices don't get messed up.
			for i := len(winBoardsIndices) - 1; i >= 0; i-- {
				winEvents = append(winEvents, WinEvent{allBoards[winBoardsIndices[i]], lastCalledOutNumber})

				// Remove the winning board from `allBoards` to prevent confusion from now on.
				allBoards[winBoardsIndices[i]] = allBoards[len(allBoards)-1]
				allBoards = allBoards[:len(allBoards)-1]
			}
		}
	}

	lastWinEvent := winEvents[len(winEvents)-1]

	sum, err := getUnmarkedSum(lastWinEvent.board)
	if err != nil {
		return "Could not parse number of board. Aborting."
	}

	return fmt.Sprintf("The score of the last winning board is %d.", sum*lastWinEvent.lastCalledOut)
}
