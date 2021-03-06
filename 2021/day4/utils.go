package day4

import (
	"errors"
	"strconv"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Sets up a ``Board`` for five lines of `input` data.
func retrieveBoard(input []string) (Board, error) {
	newBoard := *new(Board)

	if len(input) != 5 {
		return newBoard, errors.New("invalid input length")
	}

	for rowIndex, row := range input {
		splitColumn := strings.Split(row, " ")

		// If a number consists of only one digit, we will have some empty elements in `splitColumn`.
		// So let's just filter these out.
		splitColumn = common.FilterSlice(splitColumn, func(number string) bool {
			return number != ""
		})

		for columnIndex, number := range splitColumn {
			newBoard[rowIndex][columnIndex] = BoardElement{value: number, marked: false}
		}
	}

	return newBoard, nil
}

// Marks `calledOutNumber` on `allBoards` and returns the modified state.
func markNumber(allBoards []Board, calledOutNumber string) []Board {
	for boardIndex, board := range allBoards {
		for rowIndex, row := range board {
			for columnIndex, number := range row {
				if number.value == calledOutNumber {
					allBoards[boardIndex][rowIndex][columnIndex].marked = true
				}
			}
		}
	}

	return allBoards
}

// Returns a slice of all winning boards' indices in `allBoards` (empty if no board won).
// The slice will be ordered from smallest to greatest index.
func getWinningBoards(allBoards []Board) []int {
	winningBoards := []int{}

	for boardIndex, board := range allBoards {
		for i := 0; i < 5; i++ {
			// Check if a row won
			if board[i][0].marked &&
				board[i][1].marked &&
				board[i][2].marked &&
				board[i][3].marked &&
				board[i][4].marked {
				winningBoards = append(winningBoards, boardIndex)
				// We should continue here to prevent counting this board multiple times.
				continue
			}

			// Check if a column won
			if board[0][i].marked &&
				board[1][i].marked &&
				board[2][i].marked &&
				board[3][i].marked &&
				board[4][i].marked {
				winningBoards = append(winningBoards, boardIndex)
			}
		}
	}

	return winningBoards
}

// Returns the sum of all unmarked numbers in `board`.
// Throws an error if an unmarked number of `board` could not be parsed.
// Should not happen. I hope.
func getUnmarkedSum(board Board) (int, error) {
	var result int
	for _, row := range board {
		for _, number := range row {
			if !number.marked {
				parsedNumber, err := strconv.Atoi(number.value)
				if err != nil {
					return 0, err
				}

				result += parsedNumber
			}
		}
	}

	return result, nil
}
