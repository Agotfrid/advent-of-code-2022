package main

import (
	"fmt"
	"os"
	"strings"
)

func processFile(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Could not read file: %w", err)
	}

	games := strings.Split(string(file), "\n")

	return games, nil
}

func findScore(game string) int {

	if game == "" {
		return 0
	}

	// define rules, A and X = rock, B and Y = paper, C and Z = scissors
	wins := map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}

	ties := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	shapes := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	//split the string into 2 values
	split := strings.Split(game, " ")

	player1, player2 := split[0], split[1]
	var score int = shapes[player2]

	if wins[player2] == player1 {
		score += 6
	} else if ties[player2] == player1 {
		score += 3
	}

	return score

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument.")
		os.Exit(1)
	}

	filename := os.Args[1]

	games, err := processFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var totalScore int
	for _, game := range games {
		score := findScore(strings.TrimSpace(game))
		totalScore += score
	}
	fmt.Printf("Total score is: %d", totalScore)
}
