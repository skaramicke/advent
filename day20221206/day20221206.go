package day20221206

import (
	"fmt"

	"github.com/skaramicke/advent/utils"
)

func Run() {
	input := utils.ReadOrCreateInputFile("2022-12-06")
	fmt.Printf("Number of chars: %d\n", find_marker_position(input, 4))
	fmt.Printf("Number of chars: %d\n", find_marker_position(input, 14))
}

func find_marker_position(input string, marker_length int) int {
	most_recent_chars := []rune{}

	// Read characters
	for index, char := range input {
		// add to most_recent_chars
		most_recent_chars = append(most_recent_chars, char)

		// Make sure it's 4 chars
		if len(most_recent_chars) > marker_length {
			most_recent_chars = most_recent_chars[1:]
		}

		if len(most_recent_chars) == marker_length {
			// count how many of each char
			char_counts := map[rune]int{}
			for _, char := range most_recent_chars {
				char_counts[char]++
			}
			fmt.Printf("%d: %c: %d, %c: %d, %c: %d, %c: %d\n", index, most_recent_chars[0], char_counts[most_recent_chars[0]], most_recent_chars[1], char_counts[most_recent_chars[1]], most_recent_chars[2], char_counts[most_recent_chars[2]], most_recent_chars[3], char_counts[most_recent_chars[3]])
			// if all four counts are 1, return the index
			non_four := false
			for _, char := range most_recent_chars {
				if char_counts[char] != 1 {
					non_four = true
				}
			}
			if !non_four {
				return index + 1
			}
		}
	}

	return -1
}
