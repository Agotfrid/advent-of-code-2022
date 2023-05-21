package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	units := strings.Split(string(file), "\n\n")

	return units, nil
}

func findTotalCalories(unit string) (int, error) {
	lines := strings.Split(strings.TrimSpace(unit), "\n")
	total := 0
	for _, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			return 0, fmt.Errorf("Error converting values to int: %w", err)
		}
		total += value
	}

	return total, nil
}

// func findHighestSum(sums []int) (int, int) {
// 	highestSum := 0
// 	unit := 0
// 	for i, sum := range sums {
// 		if sum > highestSum {
// 			highestSum = sum
// 			unit += i
// 		}
// 	}
// 	return unit, highestSum
// }

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument.")
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Println("Please provide number of top elves to assess")
		os.Exit(1)
	}
	filename := os.Args[1]
	topElves, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error converting values to int", err)
		os.Exit(1)
	}

	units, err := readFile(filename)
	if err != nil {
		panic("error reading file")
	}

	var sums []int
	for _, unit := range units {
		sum, err := findTotalCalories(unit)
		if err != nil {
			fmt.Println("Error caclulating total number of calories", err)
			os.Exit(1)
		}
		sums = append(sums, sum)
	}
	sort.Ints(sums)

	topTotal := sums[len(sums)-topElves:]

	var total int
	for _, number := range topTotal {
		total += number
	}
	fmt.Printf("Top %d Elves are carrying %d calories\n", topElves, total)

	// unit, highestSum := findHighestSum(sums)
	// fmt.Printf("Elf #%d is carrying a total of %d calories\n", unit, highestSum)
}
