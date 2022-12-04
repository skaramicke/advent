package day20221204

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/skaramicke/advent/utils"
)

func Run() {

	input := utils.ReadOrCreateInputFile("2022-12-04")

	part1, part2 := numOverlappingPairs(string(input))
	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

func numOverlappingPairs(input string) (int, int) {
	// Split the input string into lines
	lines := strings.Split(input, "\n")

	// Initialize a counter for the number of overlapping pairs
	var numFullyOverlapping int
	var numPartlyOverlapping int

	// Loop through each line in the input
	for _, line := range lines {
		// Split the line into pairs of assignments
		pairs := strings.Split(line, ",")

		// Loop through each pair of assignments
		for i := 0; i < len(pairs); i++ {
			for j := i + 1; j < len(pairs); j++ {
				// Split each assignment into its start and end sections
				start1, end1 := splitAssignment(pairs[i])
				start2, end2 := splitAssignment(pairs[j])

				// Check if one range completely contains the other
				if (start1 <= start2 && end1 >= end2) || (start2 <= start1 && end2 >= end1) {
					// If so, increment the counter
					numFullyOverlapping++
					numPartlyOverlapping++
				} else if (start1 <= start2 && end1 >= start2) || (start2 <= start1 && end2 >= start1) {
					// If the ranges overlap, increment the counter
					numPartlyOverlapping++
				}
			}
		}
	}

	// Return the total number of overlapping pairs
	return numFullyOverlapping, numPartlyOverlapping
}

func splitAssignment(assignment string) (int, int) {
	// Split the assignment into its start and end sections
	sections := strings.Split(assignment, "-")

	// Convert the start and end sections to integers
	start, _ := strconv.Atoi(sections[0])
	end, _ := strconv.Atoi(sections[1])

	// Return the start and end sections as integers
	return start, end
}
