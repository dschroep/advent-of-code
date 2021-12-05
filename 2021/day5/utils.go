package day5

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// Parses a `line` of input (format: "x1,y1 -> x2,y2") and returns the represented ``Line``.
// Throws an error if the input was faulty.
func parseInputLine(line string) (Line, error) {
	splitInput := strings.Split(line, " -> ")
	startString, endString := splitInput[0], splitInput[1]

	splitStart := strings.Split(startString, ",")
	splitEnd := strings.Split(endString, ",")

	startX, err := strconv.Atoi(splitStart[0])
	if err != nil {
		return Line{}, errors.New("line.start.x could not be parsed")
	}

	startY, err := strconv.Atoi(splitStart[1])
	if err != nil {
		return Line{}, errors.New("line.start.y could not be parsed")
	}

	endX, err := strconv.Atoi(splitEnd[0])
	if err != nil {
		return Line{}, errors.New("line.end.x could not be parsed")
	}

	endY, err := strconv.Atoi(splitEnd[1])
	if err != nil {
		return Line{}, errors.New("line.end.y could not be parsed")
	}

	return Line{
		start: Coordinate{x: startX, y: startY},
		end:   Coordinate{x: endX, y: endY},
	}, nil
}

// Returns true only if both y coordinates of `line` are equal (= "horizontal").
func isHorizontal(line Line) bool {
	return line.start.y == line.end.y
}

// Returns true only if both x coordinates of `line` are equal (= "vertical").
func isVertical(line Line) bool {
	return line.start.x == line.end.x
}

// Returns true if line is neither horizontal nor vertical.
func isDiagonal(line Line) bool {
	return !isHorizontal(line) && !isVertical(line)
}

// Marks all the coordinates, that `line` passes on `corSys`.
// `rowSize` is needed because `corSys` is a flat representation of two dimensions.
// Returns the changed version of `corSys`.
// Note that this currently only works with horizontal or vertical lines.
func markCoordinates(line Line, corSys []byte, rowSize int) []byte {
	if isHorizontal(line) {
		minX := math.Min(float64(line.start.x), float64(line.end.x))
		maxX := math.Max(float64(line.start.x), float64(line.end.x))
		for x := minX; x <= maxX; x++ {
			corSys[line.start.y*rowSize+int(x)]++
		}
	} else if isVertical(line) {
		minY := math.Min(float64(line.start.y), float64(line.end.y))
		maxY := math.Max(float64(line.start.y), float64(line.end.y))
		for y := minY; y <= maxY; y++ {
			corSys[int(y)*rowSize+line.start.x]++
		}
	}

	return corSys
}
