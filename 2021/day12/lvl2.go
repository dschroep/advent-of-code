package day12

import (
	"fmt"

	"github.com/dschroep/advent-of-code/common"
)

// Solves level 2 of day 12 and returns the result as printable message.
func solveLvl2() string {
	inputs, err := common.GetFileInput(12)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	connections := getConnections(inputs)

	// Returns true if `path`:
	// - visits all small caves only once
	// - visits a single small cave twice
	// - visits "start" exactly once
	//
	// It does not validate how often "end" gets visited as `getAllValidPaths` will stop there by itself.
	validator := func(path Path) bool {
		visitedSmallCaveTwice := false

		for i, cave := range path {
			if cave.small() {
				for j := i + 1; j < len(path); j++ {
					if cave.equals(path[j]) {
						if visitedSmallCaveTwice {
							return false
						} else {
							visitedSmallCaveTwice = true
						}
					}
				}
			}

			if i > 0 && cave.equals(START) {
				return false
			}
		}

		return true
	}

	paths := getAllValidPaths(connections, validator)

	return fmt.Sprintf("There are %d paths, that pass the extended validation.", len(paths))
}
