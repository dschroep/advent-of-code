package day12

import (
	"fmt"
	"strings"

	"github.com/dschroep/advent-of-code/common"
)

// Parses all ``Connection``s in `inputs` and returns the result.
func getConnections(inputs []string) []Connection {
	connections := []Connection{}
	for _, input := range inputs {
		splitInput := strings.Split(input, "-")
		connections = append(connections, Connection{Cave(splitInput[0]), Cave(splitInput[1])})
	}

	return connections
}

// Returns all members of `connections` that include `cave`.
func filterConnections(connections []Connection, cave Cave) []Connection {
	remaining := []Connection{}
	for _, connection := range connections {
		if connection.a.equals(cave) || connection.b.equals(cave) {
			remaining = append(remaining, connection)
		}
	}

	return remaining
}

// Returns true if not every member of `paths` ended.
func notAllPathsEnded(paths []Path) bool {
	for _, path := range paths {
		if !path.ended() {
			return true
		}
	}

	return false
}

// Returns all valid ``Path``s that are constructable of `connections`.
// Assumes there are `connections` with the start and end cave.
// ``Path``s are considered valid if they visit small caves only once.
func getAllValidPaths(connections []Connection) []Path {
	startingConnections := filterConnections(connections, START)

	// Init `paths`.
	paths := []Path{}
	for _, startingConnection := range startingConnections {
		paths = append(paths, Path{START, startingConnection.getOpposite(START)})
	}

	// Continue `paths` as long as necessary.
	for notAllPathsEnded(paths) {
		// These will help us to keep track of all modifications that need to be done,
		// but also prevent modification of `paths` during the following loop.
		continuedPaths := []int{}
		newPaths := []Path{}

		for pathIndex, path := range paths {
			if path.ended() {
				continue
			}

			// `path` did not end -> get possible new caves for `path`.
			connectionCanidates := filterConnections(connections, path.lastCave())
			for _, connectionCandidate := range connectionCanidates {
				branch := path.createBranch(connectionCandidate.getOpposite(path.lastCave()))
				if branch.valid() {
					newPaths = append(newPaths, branch)
				}
			}

			continuedPaths = append(continuedPaths, pathIndex)
		}

		paths = append(paths, newPaths...)

		// Remove all stems from `paths` that got new branches.
		// Note that it is important to go backwards through `continuedPaths` to prevent messing up the indices.
		for i := len(continuedPaths) - 1; i >= 0; i-- {
			paths[continuedPaths[i]] = paths[len(paths)-1]
			paths = paths[:len(paths)-1]
		}
	}

	return paths
}

// Solves level 1 of day 12 and returns the result as printable message.
func solveLvl1() string {
	inputs, err := common.GetFileInput(12)
	if err != nil {
		return "Could not open input file. Aborting."
	}

	connections := getConnections(inputs)
	paths := getAllValidPaths(connections)

	return fmt.Sprintf("There are %d paths, that visit small caves only once.", len(paths))
}
