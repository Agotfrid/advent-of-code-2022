package main

import (
	"fmt"
	"os"
	"strings"
)

type Backpack struct {
	Backpack1 string
	Backpack2 string
	Backpack3 string
}

func processFile(filename string) ([]Backpack, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("THe file could not be read: %w", err)
	}

	backpacks := strings.Split(strings.TrimSpace(string(file)), "\n")
	if len(backpacks)%3 != 0 {
		return nil, fmt.Errorf("File should have lines divisible by 3")
	}

	var backpackList []Backpack
	for i := 0; i < len(backpacks); i += 3 {
		currentPack := Backpack{
			Backpack1: backpacks[i],
			Backpack2: backpacks[i+1],
			Backpack3: backpacks[i+2],
		}
		backpackList = append(backpackList, currentPack)
	}

	return backpackList, nil

}

func findCommonItemPriority(backpack1, backpack2, backpack3 string) int {
	itemRanks := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// populate first map with items from backpack1
	items1 := make(map[rune]bool)
	for _, item := range backpack1 {
		items1[item] = true
	}

	// populate second map with items present in both backpack1 and backpack2
	commonItems := make(map[rune]bool)
	for _, item := range backpack2 {
		if items1[item] {
			commonItems[item] = true
		}
	}

	// iterate through common items to see if backpack3 has it
	for _, item := range backpack3 {
		if commonItems[item] {
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
	for _, group := range list {
		commonItemPriority := findCommonItemPriority(group.Backpack1, group.Backpack2, group.Backpack3)
		totalPriority += commonItemPriority
	}

	fmt.Printf("Total Priority: %d", totalPriority)
}
