package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struct describing elf assignment pairs
type Pairs struct {
	Elf1 []int
	Elf2 []int
}

// builds an array of Pairs structs containing int slices for each elf in the pair
func processFile(filename string) ([]Pairs, error) {

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %w", err)
	}

	// first seperate the string by newline representing each Pair
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	var pairs []Pairs
	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, ",")
		ranges := make([][]int, 2)
		for i, s := range split {
			nums := strings.Split(s, "-")
			start, err := strconv.Atoi(nums[0])
			if err != nil {
				return nil, fmt.Errorf("failed to parse start value: %w", err)
			}
			end, err := strconv.Atoi(nums[1])
			if err != nil {
				return nil, fmt.Errorf("failed to parse end value: %w", err)
			}
			for j := start; j <= end; j++ {
				ranges[i] = append(ranges[i], j)
			}
		}

		pair := Pairs{
			Elf1: ranges[0],
			Elf2: ranges[1],
		}

		// save the containing values to struct
		pairs = append(pairs, pair)
	}
	return pairs, nil

}

// checks if s2 is contained in s1, returns true if it does, or false if even one number deviates
func checkPairs(s1 []int, s2 []int) bool {
	// if slice 2 is bigger than slice 1 then slice 1 cant contain it, so return false
	if len(s2) > len(s1) {
		return false
	}

	m := make(map[int]bool)
	for _, e1 := range s1 {
		m[e1] = true
	}

	// checks if all values from s2 are contained in s1, it assumes that s2 will always be smaller than s1 to return true
	for _, e2 := range s2 {
		_, present := m[e2]
		if present == false {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Need to pass file argument")
		os.Exit(1)
	}

	filename := os.Args[1]

	//convert the input file to a slice of Pairs containing 2 int slices
	pairs, err := processFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var s int
	// iterate over all Pairs structs
	for _, pair := range pairs {
		// check if one slice contains the other and vice versa
		if checkPairs(pair.Elf1, pair.Elf2) || checkPairs(pair.Elf2, pair.Elf1) {
			s++
		}
	}

	fmt.Println(s)
}
