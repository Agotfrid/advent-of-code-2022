package main

import (
	"fmt"
	"os"
	"strings"
)

type Game struct {
	Opponent string
	Outcome  string
}

func processFile(filename string) ([]Game, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Could not read file: %w", err)
	}

	games := strings.Split(string(file), "\n")

	var gamesList []Game
	for _, game := range games {
		if game == "" {
			continue
		}
		split := strings.Split(game, " ")
		currentGame := Game{
			Opponent: split[0],
			Outcome:  split[1],
		}
		gamesList = append(gamesList, currentGame)

	}

	return gamesList, nil
}

func findScore(game Game) int {

	var totalPoints int

	if game.Outcome == "X" { //lose the game
		totalPoints = 0
		switch game.Opponent {
		case "A":
			// play scizzors = 3 points
			totalPoints += 3
		case "B":
			// play A (rock) = 1 point
			totalPoints++
		case "C":
			// play B (paper) = 2 points
			totalPoints += 2
		}
	} else if game.Outcome == "Y" { // tie the game by using same opponent hand
		totalPoints = 3
		switch game.Opponent {
		case "A":
			// play rock = 1 point
			totalPoints++
		case "B":
			// play paper = 2 points
			totalPoints += 2
		case "C":
			// play scissors = 3 points
			totalPoints += 3
		}
	} else if game.Outcome == "Z" { // win the game
		totalPoints = 6
		switch game.Opponent {
		case "A":
			// play paper = 2 point
			totalPoints += 2
		case "B":
			// play scissors = 3 points
			totalPoints += 3
		case "C":
			// play rock = 1 point
			totalPoints++
		}
	} else {
		panic("Passed incorrect outcome value")
	}

	return totalPoints

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
		score := findScore(game)
		totalScore += score
	}
	fmt.Printf("Total score is: %d", totalScore)
}
