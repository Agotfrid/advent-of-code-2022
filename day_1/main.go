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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument.")
		os.Exit(1)
	}
	filename := os.Args[1]

	maxSum := 0
	maxUnit := 0
	units, err := readFile(filename)
	// fmt.Println(units)
	if err != nil {
		fmt.Println("Error reding file:", err)
	}

	for i, unit := range units {
		lines := strings.Split(strings.TrimSpace(unit), "\n")
		sum := 0

		for _, line := range lines {
			value, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error converting values to int:", err)
			}
			sum += value
		}

		if sum > maxSum {
			maxSum = sum
			maxUnit += i
		}
	}
	fmt.Printf("Elf #%d is carrying a total of %d calories\n", maxUnit, maxSum)
}
