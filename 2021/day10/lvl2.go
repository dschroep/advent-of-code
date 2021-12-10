package day10

import (
	"fmt"
	"sort"

	"github.com/dschroep/advent-of-code/common"
)

// Returns the completion score for `char`.
func getCompletionScoreForChar(char rune) int {
	switch char {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	}

	return 0
}

// Returns the corresponding completion string for `openingString`.
func getCompletionString(openingString string) string {
	var completionString string
	for i := len(openingString) - 1; i >= 0; i-- {
		char := openingString[i]

		switch char {
		case '(':
			completionString += ")"
		case '[':
			completionString += "]"
		case '{':
			completionString += "}"
		case '<':
			completionString += ">"
		}
	}

	return completionString
}

// Solves level 2 of day 10 and returns the result as printable message.
func solveLvl2() string {
	inputs, err := common.GetFileInput(10)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	completionScores := make([]int, 0)
	for _, input := range inputs {
		var openingString string

		for charIndex, character := range input {
			if isOpeningChar(character) {
				openingString += string(character)
			} else {
				// If we encounter a closing character, it has to match with the last opening character.
				if isCorrectClosingChar(character, rune(openingString[len(openingString)-1])) {
					// If it does match, we remove the last member of `openingChars`.
					openingString = openingString[:len(openingString)-1]
				} else {
					// If it does not match, we encountered an error -> move on to next line.
					break
				}
			}

			// Reached the last character -> calculate the completion score.
			if charIndex == len(input)-1 {
				completionString := getCompletionString(openingString)

				var completionScore int
				for _, completionChar := range completionString {
					completionScore *= 5
					completionScore += getCompletionScoreForChar(completionChar)
				}
				completionScores = append(completionScores, completionScore)
			}
		}
	}

	sort.Ints(completionScores)
	// We know that `len(completionScores)` will be odd
	// -> subtract 1 and divide by 2 to get the middle score
	searchedScoreIndex := (len(completionScores) - 1) / 2
	searchedScore := completionScores[searchedScoreIndex]
	return fmt.Sprintf("The searched completion score is %d.", searchedScore)
}
