package main

import (
	"fmt"
	"os"
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

func findTotalCalories(lines []string) int {
	total := 0
	for _, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error converting values to int", err)
		}
		total += value
	}

	return total
}

func findHighestSum(sums []int) (int, int) {
	highestSum := 0
	unit := 0
	for i, sum := range sums {
		if sum > highestSum {
			highestSum = sum
			unit += i
		}
	}
	return unit, highestSum
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument.")
		os.Exit(1)
	}
	filename := os.Args[1]

	units, err := readFile(filename)
	if err != nil {
		panic("error reading file")
	}

	sums := findTotalCalories(units)
	unit, highestSum := findHighestSum(sums)
	fmt.Printf("Elf #%d is carrying a total of %d calories\n", unit, highestSum)
}
