package day9

import (
	"fmt"
	"sort"

	"github.com/dschroep/advent-of-code/common"
)

// TODO: Use from previous day
type Coordinate struct{ x, y int }

// Returns true if `coordinate` is not yet present in `exploredCoordinates`.
func isNotExplored(coordinate Coordinate, exploredCoordinates []Coordinate) bool {
	for _, exploredCoordinate := range exploredCoordinates {
		if exploredCoordinate.x == coordinate.x && exploredCoordinate.y == coordinate.y {
			return false
		}
	}

	return true
}

// Returns the total basin size for basin at `coordinate` in `heights`.
func getBasinSize(coordinate Coordinate, heights []string) int {
	exploredLocations := []Coordinate{coordinate}
	for {
		newLocations := []Coordinate{}

		// TODO: This could probably be improved by only looking at the new locations
		for _, exploredLocation := range exploredLocations {
			top := Coordinate{exploredLocation.x, exploredLocation.y - 1}
			if top.y >= 0 && heights[top.y][top.x] != '9' && isNotExplored(top, append(exploredLocations, newLocations...)) {
				newLocations = append(newLocations, top)
			}

			bottom := Coordinate{exploredLocation.x, exploredLocation.y + 1}
			if bottom.y < len(heights) && heights[bottom.y][bottom.x] != '9' && isNotExplored(bottom, append(exploredLocations, newLocations...)) {
				newLocations = append(newLocations, bottom)
			}

			left := Coordinate{exploredLocation.x - 1, exploredLocation.y}
			if left.x >= 0 && heights[left.y][left.x] != '9' && isNotExplored(left, append(exploredLocations, newLocations...)) {
				newLocations = append(newLocations, left)
			}

			right := Coordinate{exploredLocation.x + 1, exploredLocation.y}
			if right.x < len(heights[right.y]) && heights[right.y][right.x] != '9' && isNotExplored(right, append(exploredLocations, newLocations...)) {
				newLocations = append(newLocations, right)
			}
		}

		if len(newLocations) == 0 {
			break
		}

		exploredLocations = append(exploredLocations, newLocations...)
	}

	return len(exploredLocations)
}

// Solves level 2 of day 9 and returns the result as printable message.
func solveLvl2() string {
	heights, err := common.GetFileInput(9)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	basinSizes := make([]int, 0)
	for rowIndex, row := range heights {
		for columnIndex := range row {
			if isLowPoint(getHeights(columnIndex, rowIndex, heights)) {
				basinSizes = append(basinSizes, getBasinSize(Coordinate{columnIndex, rowIndex}, heights))
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	result := basinSizes[0] * basinSizes[1] * basinSizes[2]
	return fmt.Sprintf("The product of the three largest basins' sizes is %d.", result)
}
