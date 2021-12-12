package day12

// Abstraction of a cave to its name.
type Cave string

// Checks whether a cave is big or small based on its name.
func (cave Cave) isBig() bool {
	return cave[0] >= 'A' && cave[0] <= 'Z'
}

// A connection from one ``Cave`` to another.
type Connection struct {
	a Cave
	b Cave
}

// Returns the members of `connections` that contain `cave` either as `Connection.a` or `Connection.b`.
func filterForCave(connections []Connection, cave Cave) []Connection {
	result := make([]Connection, 0)
	for _, connection := range connections {
		if connection.a == cave || connection.b == cave {
			result = append(result, connection)
		}
	}

	return result
}

// A path from the "start" ``Cave`` to the "end" ``Cave``.
type Path []Cave

// Returns `true` if the last element of `path` is the "end" ``Cave``
func (path Path) ended() bool {
	return path[len(path)-1] == NAME_END
}

// Returns `true` if `path` already contains `to`.
func (path Path) alreadyWentThere(to Cave) bool {
	for _, cave := range path {
		if cave == to {
			return true
		}
	}

	return false
}

// Returns a copy of `path`, where `cave` was appended.
func (path Path) branchPath(cave Cave) Path {
	return append(path, cave)
}
