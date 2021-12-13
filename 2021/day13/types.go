package day13

import "strings"

type Dot struct {
	x int
	y int
}

func (dot Dot) equals(other Dot) bool {
	return dot.x == other.x && dot.y == other.y
}

type Paper []Dot

// Returns `paper` as a string representation.
// TODO: There has to be a more elegant way for this.
func (paper Paper) toString() string {
	// Find out how large the string has to be.
	var maxX, maxY int
	for _, dot := range paper {
		if dot.x > maxX {
			maxX = dot.x
		}
		if dot.y > maxY {
			maxY = dot.y
		}
	}

	var byteRepresentation [][]byte
	for y := 0; y <= maxY; y++ {
		byteRepresentation = append(byteRepresentation, []byte(strings.Repeat(".", maxX+1)))
	}

	for _, dot := range paper {
		byteRepresentation[dot.y][dot.x] = '#'
	}

	var stringRepresentation string
	for _, line := range byteRepresentation {
		// The three tabs at the beginning of a line aren't necessary, but they integrate the result well into the overall output.
		// I know, it's not beautiful to make this receiver function for only one purpose but ¯\_(ツ)_/¯
		stringRepresentation += "\t\t\t" + string(line) + "\n"
	}

	return stringRepresentation
}

// Returns the total amount of dots, but counts overlapping dots only once.
func (paper Paper) countUnique() int {
	var uniqueDots []Dot

DotLoop:
	for _, dot := range paper {
		// Check if `dot`'s value was already counted and go on if it was.
		for _, uniqueDot := range uniqueDots {
			if dot.equals(uniqueDot) {
				continue DotLoop
			}
		}

		uniqueDots = append(uniqueDots, dot)
	}

	return len(uniqueDots)
}

// Folds the paper `along` 'x' or 'y' at the given `position`.
func (paper *Paper) fold(along rune, position int) {
	if along == 'x' {
		for dotIndex, dot := range *paper {
			xDistance := dot.x - position
			if xDistance > 0 {
				(*paper)[dotIndex].x = position - xDistance
			}
		}
	} else if along == 'y' {
		for dotIndex, dot := range *paper {
			yDistance := dot.y - position
			if yDistance > 0 {
				(*paper)[dotIndex].y = position - yDistance
			}
		}
	} // TODO: Normally we'd need to check for other values than 'x' and 'y' now...
}
