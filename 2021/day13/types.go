package day13

type Dot struct {
	x int
	y int
}

func (dot Dot) equals(other Dot) bool {
	return dot.x == other.x && dot.y == other.y
}

type Paper []Dot

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
