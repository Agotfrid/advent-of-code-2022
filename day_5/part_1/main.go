package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Directions is as struct to contain directions for movement of boxes
type Directions struct {
	Containers       int
	SourceStack      int
	DestinationStack int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("too few arguments, please pass input file")
		os.Exit(1)
	}

	filename := os.Args[1]
	processFile(filename)
}

func moveCrates(directions string) string {
	return ""
}

func processFile(filename string) ([]string, []Directions, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot read the file: %w", err)
	}

	// will split the file into 2 strings, one for the arrays, and the other for directions
	content := strings.Split(string(file), "\n\n")
	// take first string which contains stacks of boxes and convert those into slices

	stacksList := strings.Split(content[0], "\n")
	stacks := processStacks(stacksList)

	directionsList := strings.Split(content[1], "\n")
	directions := processDirections(directionsList)

	return stacks, directions, nil
}

func processStacks(stacksList []string) []string {
	numLine := stacksList[len(stacksList)-1]
	stacksList = stacksList[:len(stacksList)-1] // remove line of numbers

	stacksMap := make(map[int]int)
	for i, char := range numLine {
		if char != ' ' {
			stackNum, _ := strconv.Atoi(string(char))
			stacksMap[i] = stackNum - 1 // we assume num line will always have 1 as first number
		}
	}

	return nil
}

func processDirections(directionsList []string) []Directions {
	var directions []Directions
	for _, direct := range directionsList {
		if direct == "" {
			continue
		}
		// split line into Containers to be moved, source Stack and destination stack
		directionParts := strings.Fields(direct) // Splits the string into a slice on whitespace

		// assuming this string format "move {number of boxes} from {source} to {destination}"
		containers, _ := strconv.Atoi(directionParts[1])  // The second item (index 1) is the number of containers to move
		sourceStack, _ := strconv.Atoi(directionParts[3]) // The fourth item (index 3) is the source stack
		destStack, _ := strconv.Atoi(directionParts[5])   // The sixth item (index 5) is the destination stack

		directionsStruct := Directions{
			Containers:       containers,
			SourceStack:      sourceStack,
			DestinationStack: destStack,
		}

		directions = append(directions, directionsStruct)
	}

	return directions
}
