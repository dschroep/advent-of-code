package day10

import (
	"fmt"

	"github.com/dschroep/advent-of-code/common"
)

// Returns the error score for `char`.
func getErrorScoreForChar(char rune) int {
	switch char {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}

	return 0
}

// Solves level 1 of day 10 and returns the result as printable message.
func solveLvl1() string {
	inputs, err := common.GetFileInput(10)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	var errorScore int
	for _, input := range inputs {
		var openingString string

		for _, character := range input {
			if isOpeningChar(character) {
				// Keep track of all the opening characters.
				openingString += string(character)
			} else {
				// If we encounter a closing character, it has to match with the last opening character.
				if isCorrectClosingChar(character, rune(openingString[len(openingString)-1])) {
					// If it does match, we remove the last member of `openingChars`.
					openingString = openingString[:len(openingString)-1]
				} else {
					// If it does not match, we encountered an error -> add to score and move to next line.
					errorScore += getErrorScoreForChar(character)
					break
				}
			}
		}
	}

	return fmt.Sprintf("The total syntax error score is %d.", errorScore)
}
