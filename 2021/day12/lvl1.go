package day12

import (
	"fmt"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Names of important caves.
const (
	NAME_START = "start"
	NAME_END   = "end"
)

// Returns `true` if the last element of all `paths` are the "end" `Cave`.
func allPathsEnded(paths []Path) bool {
	for _, path := range paths {
		if !path.ended() {
			return false
		}
	}

	return true
}

// Returns a slice of all ``Path``s starting with the "start" ``Cave``.
func retrieveStartPaths(connections []Connection) []Path {
	paths := make([]Path, 0)

	startConnections := filterForCave(connections, NAME_START)
	for _, startConnection := range startConnections {
		var newPath Path
		if startConnection.a == NAME_START {
			newPath = Path{startConnection.a, startConnection.b}
		} else {
			newPath = Path{startConnection.b, startConnection.a}
		}

		paths = append(paths, newPath)
	}

	return paths
}

// Returns a new slice of ``Path``s, that continue where `paths` ended.
func continuePaths(paths []Path, allConnections []Connection) []Path {
	newPaths := paths
	for newPathIndex, newPath := range newPaths {
		// Do nothing for this path if it already ended.
		if newPath.ended() {
			continue
		}

		lastCave := newPath[len(newPath)-1]
		possibleConnections := filterForCave(allConnections, lastCave)

		for possibleConnectionIndex, possibleConnection := range possibleConnections {
			var newCave Cave
			if possibleConnection.a == lastCave {
				newCave = possibleConnection.b
			} else {
				newCave = possibleConnection.a
			}

			// Do not visit a small cave multiple times.
			if newPath.alreadyWentThere(newCave) && !newCave.isBig() {
				continue
			}

			// The first `possibleConnection` can just overwrite the existing one.
			if possibleConnectionIndex == 0 {
				newPaths[newPathIndex] = append(newPaths[newPathIndex], newCave)
				continue
			}

			// FIXME: This seems to create paths that already exist.
			newPaths = append(newPaths, newPath.branchPath(newCave))
		}
	}

	return newPaths
}

// Solves level 1 of day 12 and returns the result as printable message.
func solveLvl1() string {
	inputs, err := common.GetFileInput(12)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	// Init connections.
	connections := make([]Connection, 0)
	for _, input := range inputs {
		splitInput := strings.Split(input, "-")
		connections = append(connections, Connection{
			Cave(splitInput[0]),
			Cave(splitInput[1]),
		})
	}

	// Now, let's figure out all paths.
	paths := retrieveStartPaths(connections)
	// FIXME:
	// for !allPathsEnded(paths) {
	// 	paths = continuePaths(paths, connections)
	// }

	return fmt.Sprintf("There are %d paths, that visit small caves only once.", len(paths))
}
