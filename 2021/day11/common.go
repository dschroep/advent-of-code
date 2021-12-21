package day11

// Constructs a new ``OctopusField`` from `input`.
func getOctopusField(input []string) OctopusField {
	octopusField := OctopusField{}
	for rowIndex, row := range input {
		octopusField = append(octopusField, make([]Octopus, len(row)))

		for columnIndex, energyLevel := range row {
			octopusField[rowIndex][columnIndex] = Octopus{int(energyLevel - '0'), false}
		}
	}

	return octopusField
}

// Increases the energy level of all ``Octopus`` in `octopusField` by 1 and returns the result.
func increaseAll(octopusField OctopusField) OctopusField {
	for rowIndex, row := range octopusField {
		for colIndex := range row {
			octopusField[rowIndex][colIndex].energyLevel++
		}
	}

	return octopusField
}

// Let's all ``Octopus`` with an energy level greater than 9 flash.
// Also starts a chain reaction if necessary.
// Returns the modified `octopusField` and the total amount of flashes.
func flashFlashable(octopusField OctopusField) (OctopusField, int) {
	totalFlashes := 0

	for rowIndex, row := range octopusField {
		for colIndex := range row {
			currentOctopus := &octopusField[rowIndex][colIndex]

			if currentOctopus.energyLevel > 9 && !currentOctopus.flashed {
				currentOctopus.flashed = true
				totalFlashes++

				// Increase energy level of all neighbors.
				for neighborRow := rowIndex - 1; neighborRow <= rowIndex+1; neighborRow++ {
					for neighborCol := colIndex - 1; neighborCol <= colIndex+1; neighborCol++ {
						if neighborCol == colIndex && neighborRow == rowIndex {
							// This is not a neighbor, this is `currentOctopus`
							continue
						}

						if neighborRow >= 0 && neighborRow <= len(octopusField)-1 && neighborCol >= 0 && neighborCol <= len(octopusField[neighborRow])-1 {
							octopusField[neighborRow][neighborCol].energyLevel++
						}
					}
				}

				newFlashes := 0
				octopusField, newFlashes = flashFlashable(octopusField)
				totalFlashes += newFlashes
			}
		}
	}

	return octopusField, totalFlashes
}

// Resets all ``Octopus`` in `octopusField` that flashed and returns the result.
func resetFlashed(octopusField OctopusField) OctopusField {
	for rowIndex, row := range octopusField {
		for colIndex := range row {
			currentOctopus := &octopusField[rowIndex][colIndex]

			if currentOctopus.flashed {
				*currentOctopus = Octopus{0, false}
			}
		}
	}

	return octopusField
}

// Simulates a step on `octopusField` and returns the modified `octopusField` as well as the total amount of flashes.
func simulateStep(octopusField OctopusField) (OctopusField, int) {
	totalFlashes := 0

	octopusField = increaseAll(octopusField)
	octopusField, totalFlashes = flashFlashable(octopusField)
	octopusField = resetFlashed(octopusField)

	return octopusField, totalFlashes
}
