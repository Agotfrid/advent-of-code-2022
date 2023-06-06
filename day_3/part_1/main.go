package main

import (
	"fmt"
	"os"
	"strings"
)

type Backpack struct {
	Compartment1 string
	Compartment2 string
}

func processFile(filename string) ([]Backpack, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("THe file could not be read: %w", err)
	}

	var backpackList []Backpack
	backpacks := strings.Split(string(file), "\n")
	for _, backpack := range backpacks {
		if backpack == "" {
			continue
		}
		half := len(backpack) / 2
		currentPack := Backpack{
			Compartment1: backpack[:half],
			Compartment2: backpack[half:],
		}
		backpackList = append(backpackList, currentPack)
	}

	return backpackList, nil

}

func findCommonItemPriority(comp1, comp2 string) int {
	itemRanks := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	items := make(map[rune]bool)
	for _, item := range comp1 {
		items[item] = true
	}

	for _, item := range comp2 {
		if items[item] {
			index := strings.IndexRune(itemRanks, item)
			if index >= 0 {
				return index + 1
			}
		}
	}

	return 0

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide filename as argument")
		os.Exit(1)
	}
	filename := os.Args[1]

	list, err := processFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var totalPriority int
	for _, backpack := range list {
		commonItemPriority := findCommonItemPriority(backpack.Compartment1, backpack.Compartment2)
		totalPriority += commonItemPriority
	}

	fmt.Printf("Total Priority: %d", totalPriority)
}
