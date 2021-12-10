package day10

// Returns true if `char` is an opening character.
func isOpeningChar(char rune) bool {
	return char == '(' || char == '[' || char == '{' || char == '<'
}

// Returns true if `char` is the right closing character for `openingChar`.
func isCorrectClosingChar(char rune, openingChar rune) bool {
	switch {
	case openingChar == '(' && char == ')':
		fallthrough
	case openingChar == '[' && char == ']':
		fallthrough
	case openingChar == '{' && char == '}':
		fallthrough
	case openingChar == '<' && char == '>':
		return true
	}

	return false
}
