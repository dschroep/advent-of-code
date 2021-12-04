package common

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// Reads the input file for `day` and returns it's lines as string slice.
// Returns an error if the file could not be found.
func GetFileInput(day byte) ([]string, error) {
	filePath := fmt.Sprintf("day%d/input.txt", day)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("file not found")
	}
	defer file.Close()

	result := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}
