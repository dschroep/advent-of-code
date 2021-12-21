package day12

// Special caves.
const (
	START Cave = "start"
	END   Cave = "end"
)

// Name of a cave.
type Cave string

// Returns true if both ``Cave``s have the same value.
func (cave Cave) equals(other Cave) bool {
	return string(cave) == string(other)
}

// Returns true if `cave` is small.
func (cave Cave) small() bool {
	return cave[0] >= 'a' && cave[0] <= 'z'
}

// Representation of a connection of two caves (like one line of input).
// Note that this cannot have an order.
// "start" could also be `Connection.b` and "end" could be `Connection.a`.
type Connection struct {
	a Cave
	b Cave
}

// Returns the opposite of `cave` in `connection`.
func (connection Connection) getOpposite(cave Cave) Cave {
	if connection.a.equals(cave) {
		return connection.b
	}

	return connection.a
}

// Representation of a path like "start,A,c,A,end"
type Path []Cave

// Returns a copy of `path`, which got `cave` appended.
func (path Path) createBranch(cave Cave) Path {
	newPath := make(Path, len(path))
	copy(newPath, path)

	return append(newPath, cave)
}

// Returns `path`'s last ``Cave``.
func (path Path) lastCave() Cave {
	return path[len(path)-1]
}

// Returns true if `path`'s last cave is `END`.
func (path Path) ended() bool {
	return path.lastCave() == END
}
