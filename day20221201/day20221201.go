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
		part1, part2 := count_calories(string(input))
		fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
	} else {
		// If error, create file and ask user to try again
		os.WriteFile("inputs/2022-12-01.txt", []byte{}, 0644)
		println("Please try again after adding input to inputs/2022-12-01.txt")
	}

}

func count_calories(input string) (int, int) {
	// For each line in the input
	scanner := bufio.NewScanner(strings.NewReader(input))
	sums := make([]int, 0)
	currentSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		// Try to convert line to integer
		lineInt := 0
		if _, err := fmt.Sscanf(line, "%d", &lineInt); err == nil {
			currentSum += lineInt
		} else {
			// Otherwise, add the current sum to the list of sums and reset the current sum
			sums = append(sums, currentSum)
			currentSum = 0
		}
	}

	// Find the three largest sums
	// Sort the sums
	for i := 0; i < len(sums); i++ {
		for j := i + 1; j < len(sums); j++ {
			if sums[i] < sums[j] {
				sums[i], sums[j] = sums[j], sums[i]
			}
		}
	}
	// Add the first three
	largestThreeSums := sums[0] + sums[1] + sums[2]

	return sums[0], largestThreeSums
}
