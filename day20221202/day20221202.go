package day20221202

import (
	"fmt"
	"os"
	"strings"
)

func Prompt() {
	// If inputs/2022-12-02.txt exists, use that as input
	// Otherwise, create it and ask user to try again

	// Try reading file
	if input, err := os.ReadFile("inputs/2022-12-02.txt"); err == nil {
		part1 := totalRockPaperScissorScore(string(input))
		part2 := totalWithCorrectInstructions(string(input))
		fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
	} else {
		// If error, create file and ask user to try again
		os.WriteFile("inputs/2022-12-02.txt", []byte{}, 0644)
		println("Please try again after adding input to inputs/2022-12-02.txt")
	}

}

func totalRockPaperScissorScore(input string) int {

	// Split the input into lines.
	lines := strings.Split(input, "\n")

	// Initialize the total score to 0.
	totalScore := 0

	// A, X: Rock
	// B, Y: Paper
	// C, Z: Scissors

	// Loop over the lines.
	for _, line := range lines {

		moves := strings.Split(line, " ")

		// Split the line into opponent's shape and your shape.
		oppoMove, myMove := moves[0], moves[1]

		// Initialize the score for this round to 0
		roundScore := 0

		switch myMove {
		case "X":
			roundScore = 1
			switch oppoMove {
			case "C":
				roundScore += 6
			case "A":
				roundScore += 3
			}

		case "Y":
			roundScore = 2
			switch oppoMove {
			case "A":
				roundScore += 6
			case "B":
				roundScore += 3
			}
		case "Z":
			roundScore = 3
			switch oppoMove {
			case "B":
				roundScore += 6
			case "C":
				roundScore += 3
			}
		}
		// Add the score for this round to the total score
		totalScore += roundScore
	}

	// Return the total score.
	return totalScore
}

func totalWithCorrectInstructions(input string) int {

	// Split the input into lines.
	lines := strings.Split(input, "\n")

	// Initialize the total score to 0.
	totalScore := 0

	// A, X: Rock
	// B, Y: Paper
	// C, Z: Scissors

	// Loop over the lines.
	for _, line := range lines {

		moves := strings.Split(line, " ")

		// Split the line into opponent's shape and your shape.
		oppoMove, neededOutcome := moves[0], moves[1]

		var myMove string

		switch neededOutcome {
		case "Z":
			// I need to win
			switch oppoMove {
			case "A":
				myMove = "Y"
			case "B":
				myMove = "Z"
			case "C":
				myMove = "X"
			}
		case "Y":
			// I need to tie
			switch oppoMove {
			case "A":
				myMove = "X"
			case "B":
				myMove = "Y"
			case "C":
				myMove = "Z"
			}
		case "X":
			// I need to lose
			switch oppoMove {
			case "A":
				myMove = "Z"
			case "B":
				myMove = "X"
			case "C":
				myMove = "Y"
			}
		}

		// Initialize the score for this round to 0
		roundScore := 0

		switch myMove {
		case "X":
			roundScore = 1
			switch oppoMove {
			case "C":
				roundScore += 6
			case "A":
				roundScore += 3
			}

		case "Y":
			roundScore = 2
			switch oppoMove {
			case "A":
				roundScore += 6
			case "B":
				roundScore += 3
			}
		case "Z":
			roundScore = 3
			switch oppoMove {
			case "B":
				roundScore += 6
			case "C":
				roundScore += 3
			}
		}
		// Add the score for this round to the total score
		totalScore += roundScore
	}

	// Return the total score.
	return totalScore
}
