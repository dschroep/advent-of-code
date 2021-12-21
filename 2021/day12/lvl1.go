package day12

import (
	"fmt"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 1 of day 12 and returns the result as printable message.
func solveLvl1() string {
	inputs, err := common.GetFileInput(12)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	connections := getConnections(inputs)

	// Returns true if `path` does not contain a small cave multiple times.
	validator := func(path Path) bool {
		for i, cave := range path {
			if cave.small() {
				for j := i + 1; j < len(path); j++ {
					if cave.equals(path[j]) {
						return false
					}
				}
			}
		}

		return true
	}

	paths := getAllValidPaths(connections, validator)

	return fmt.Sprintf("There are %d paths, that visit small caves only once.", len(paths))
}
