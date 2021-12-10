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
		openingChars := make([]rune, 0)

	CharacterLoop:
		for _, character := range input {
			switch character {
			// Keep track of all the opening characters.
			case '(':
				fallthrough
			case '[':
				fallthrough
			case '{':
				fallthrough
			case '<':
				openingChars = append(openingChars, character)
			// If we encounter a closing character, it has to match with the last opening character.
			// If it does match, we remove the last member of `openingChars`.
			// If it does not match, we encountered an error -> add to `errorScore` and move on to next line.
			case ')':
				if openingChars[len(openingChars)-1] != '(' {
					errorScore += getErrorScoreForChar(character)
					break CharacterLoop
				} else {
					openingChars = openingChars[:len(openingChars)-1]
				}
			case ']':
				if openingChars[len(openingChars)-1] != '[' {
					errorScore += getErrorScoreForChar(character)
					break CharacterLoop
				} else {
					openingChars = openingChars[:len(openingChars)-1]
				}
			case '}':
				if openingChars[len(openingChars)-1] != '{' {
					errorScore += getErrorScoreForChar(character)
					break CharacterLoop
				} else {
					openingChars = openingChars[:len(openingChars)-1]
				}
			case '>':
				if openingChars[len(openingChars)-1] != '<' {
					errorScore += getErrorScoreForChar(character)
					break CharacterLoop
				} else {
					openingChars = openingChars[:len(openingChars)-1]
				}
			}
		}
	}

	return fmt.Sprintf("The total syntax error score is %d.", errorScore)
}
