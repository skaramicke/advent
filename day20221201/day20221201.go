package day20221201

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Prompt() {
	// If inputs/2022-12-01.txt exists, use that as input
	// Otherwise, create it and ask user to try again

	// Try reading file
	if input, err := os.ReadFile("inputs/2022-12-01.txt"); err == nil {
		fmt.Printf("Result: %d\n", count_calories(string(input)))
	} else {
		// If error, create file and ask user to try again
		os.WriteFile("inputs/2022-12-01.txt", []byte{}, 0644)
		println("Please try again after adding input to inputs/2022-12-01.txt")
	}

}

func count_calories(input string) int {

	// For each line in the input
	sums := make([]int, 0)
	currentSum := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		// Try to convert line to integer
		lineInt := 0
		if _, err := fmt.Sscanf(line, "%d", &lineInt); err == nil {
			currentSum += int(lineInt)
		} else {
			// Otherwise, add the current sum to the list of sums and reset the current sum
			sums = append(sums, currentSum)
			currentSum = 0
		}
	}

	// Find the largest sum in sums
	largestSum := 0
	for _, sum := range sums {
		if sum > largestSum {
			largestSum = sum
		}
	}

	return largestSum
}
