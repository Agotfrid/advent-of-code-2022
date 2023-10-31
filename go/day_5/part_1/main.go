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
	stacks, directions, err := processFile(filename)
	if err != nil {
		os.Exit(1)
	}

	topCrates := moveCrates(stacks, directions)
	fmt.Println(topCrates)
}

func moveCrates(stacks [][]string, directions []Directions) string {

	// follow steps from directions to affect the slices
	for _, direction := range directions {
		n := direction.Containers
		source := direction.SourceStack
		dest := direction.DestinationStack

		movingCrates := make([]string, n)
		copy(movingCrates, stacks[source][len(stacks[source])-n:])

		for i, j := 0, len(movingCrates)-1; i < j; i, j = i+1, j-1 {
			movingCrates[i], movingCrates[j] = movingCrates[j], movingCrates[i]
		}
		stacks[dest] = append(stacks[dest], movingCrates...)

		stacks[source] = stacks[source][:len(stacks[source])-n]
	}

	var topContainers string
	for _, container := range stacks {
		topContainers += string(container[len(container)-1])
	}
	return topContainers
}

func processFile(filename string) ([][]string, []Directions, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot read the file: %w", err)
	}

	// will split the file into 2 strings, one for the arrays, and the other for directions
	content := strings.Split(string(file), "\n\n")
	// take first string which contains stacks of boxes and convert those into slices

	stacksList := strings.Split(content[0], "\n")
	stacks, err := processStacks(stacksList)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not process stacks: %w", err)
	}

	directionsList := strings.Split(content[1], "\n")
	directions := processDirections(directionsList)

	return stacks, directions, nil
}

func reverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func processStacks(stacksList []string) ([][]string, error) {
	numLine := stacksList[len(stacksList)-1]
	stacksList = reverseSlice(stacksList[:len(stacksList)-1]) // remove line of numbers

	// will create map of stacks and relative character position of the boxes
	stacksMap := make(map[int]int)
	for i, char := range numLine {
		if char != ' ' {
			stackNum, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, fmt.Errorf("could not pupulate stack map: %w", err)
			}
			stacksMap[i] = stackNum - 1 // we assume num line will always have 1 as first number
		}
	}

	numOfStacks := len(stacksMap)
	stacks := make([][]string, numOfStacks)

	for _, stack := range stacksList {
		for pos, box := range stack {
			if box != ' ' && box != '[' && box != ']' {
				stackNum := stacksMap[pos]
				stacks[stackNum] = append(stacks[stackNum], string(box))
			}
		}
	}
	return stacks, nil
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
			SourceStack:      sourceStack - 1,
			DestinationStack: destStack - 1,
		}

		directions = append(directions, directionsStruct)
	}

	return directions
}
