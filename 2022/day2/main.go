package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	finalScore := 0
	inputScanner := bufio.NewScanner(f)
	for inputScanner.Scan() {
		input := inputScanner.Text()
		// fmt.Println(input)
		val := strings.Split(input, " ")
		finalScore += score4Part2(val[0], val[1])
	}

	fmt.Println(finalScore)
}

// score4Part1 determines the score of the game
// given oChoice and yChoice.
// oChoice = Opponent's Choice
// yChoice = Your choice
func score4Part1(oChoice, yChoice string) int {
	score := 0

	//score for your choice
	switch yChoice {
	case "X": // rock
		score = 1
	case "Y": // paper
		score = 2
	case "Z": // scissors
		score = 3
	}

	// draw
	if (oChoice == "A" && yChoice == "X") || (oChoice == "B" && yChoice == "Y") || (oChoice == "C" && yChoice == "Z") {
		score += 3
	}

	// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
	// A/X - Rock
	// B/Y - Paper
	// C/Z - Scissors

	//win
	if (yChoice == "X" && oChoice == "C") || (yChoice == "Z" && oChoice == "B") || (yChoice == "Y" && oChoice == "A") {
		score += 6
	} else {
		score += 0
	}
	fmt.Printf("Opponent choice is %s and my choice is %s. Score is %d\n", oChoice, yChoice, score)
	return score
}

// score determines the score of the game
// given oChoice and endResult.
// oChoice = Opponent's Choice
// endResult = endResult of the game
func score4Part2(oChoice, endResult string) int {
	score := 0

	//score for result
	switch endResult {
	case "X": // lose
		score = 0
	case "Y": // draw
		score = 3
	case "Z": // win
		score = 6
	}

	//identify your choice from endResult and oChoice
	// A - Rock, B - Paper, C - Sciccors
	// X - Lose, Y - Draw, Z - Win
	yChoice := ""

	if endResult == "Y" { // draw
		yChoice = oChoice
	} else if endResult == "X" { // lose
		if oChoice == "A" {
			yChoice = "C"
		} else if oChoice == "B" {
			yChoice = "A"
		} else {
			yChoice = "B"
		}
	} else { //win
		if oChoice == "A" {
			yChoice = "B"
		} else if oChoice == "B" {
			yChoice = "C"
		} else {
			yChoice = "A"
		}
	}

	//score for your choice
	switch yChoice {
	case "A": // rock
		score += 1
	case "B": // paper
		score += 2
	case "C": // scissors
		score += 3
	}

	fmt.Printf("Opponent choice is %s and endResult is %s. So my choice is %s. Score is %d\n", oChoice, endResult, yChoice, score)
	return score
}
