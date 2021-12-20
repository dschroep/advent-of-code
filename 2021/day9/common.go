package day9

// Returns true if `height` is the smallest in [`height`, `neighbors`].
func isLowPoint(height int, neighbors []int) bool {
	for _, neighbor := range neighbors {
		if height >= neighbor {
			return false
		}
	}

	return true
}

// Returns (height, neighbors)
// - height: integer value of `inputs[y][x]`
// - neighbors: the surrounding values of `inputs[y][x]`
func getHeights(x, y int, inputs []string) (int, []int) {
	// This function makes heavy use of `int(<rune> - '0')`, which is a quick hack to convert
	// the rune representation of a digit to an int. Potentially dangerous, but not important here.
	height := int(inputs[y][x] - '0')

	neighbors := make([]int, 0)
	if y > 0 {
		topNeighbor := int(inputs[y-1][x] - '0')
		neighbors = append(neighbors, topNeighbor)
	}

	if y < len(inputs)-1 {
		bottomNeighbor := int(inputs[y+1][x] - '0')
		neighbors = append(neighbors, bottomNeighbor)
	}

	if x > 0 {
		leftNeighbor := int(inputs[y][x-1] - '0')
		neighbors = append(neighbors, leftNeighbor)
	}

	if x < len(inputs[y])-1 {
		rightNeighbor := int(inputs[y][x+1] - '0')
		neighbors = append(neighbors, rightNeighbor)
	}

	return height, neighbors
}
