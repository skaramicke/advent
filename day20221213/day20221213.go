package day20221213

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/skaramicke/advent/utils"
)

func Run() {
	input := utils.ReadOrCreateInputFile("2022-12-13")
	fmt.Printf("N/o sorted packets: %d\n", countSortedPackets(input))
	fmt.Printf("Divisor index product: %d\n", sortAndDividePackets(input))
}

// When comparing two values, the first value is called left and the second value is called right. Then:
// If both values are integers, the lower integer should come first. If the left integer is lower than the right integer, the inputs are in the right order. If the left integer is higher than the right integer, the inputs are not in the right order. Otherwise, the inputs are the same integer; continue checking the next part of the input.
// If both values are lists, compare the first value of each list, then the second value, and so on. If the left list runs out of items first, the inputs are in the right order. If the right list runs out of items first, the inputs are not in the right order. If the lists are the same length and no comparison makes a decision about the order, continue checking the next part of the input.
// If exactly one value is an integer, convert the integer to a list which contains that integer as its only value, then retry the comparison. For example, if comparing [0,0,0] and 2, convert the right value to [2] (a list containing 2); the result is then found by instead comparing [0,0,0] and [2].

func comparePackets(left, right interface{}) int {

	// If both are integers, compare them
	if leftInt, ok := left.(float64); ok {
		if rightInt, ok := right.(float64); ok {
			if leftInt < rightInt {
				return -1
			} else if leftInt > rightInt {
				return 1
			} else {
				return 0
			}
		}
	}

	// If both are lists, compare their items
	if leftList, ok := left.([]interface{}); ok {
		if rightList, ok := right.([]interface{}); ok {
			for i := 0; i < len(leftList); i++ {
				if i >= len(rightList) {
					return 1
				}

				result := comparePackets(leftList[i], rightList[i])
				if result != 0 {
					return result
				}
			}

			if len(leftList) < len(rightList) {
				return -1
			} else {
				return 0
			}
		}
	}

	// If one is an integer and the other is a list, convert the integer to a list and retry
	if leftInt, ok := left.(float64); ok {
		return comparePackets([]interface{}{leftInt}, right)
	}

	if rightInt, ok := right.(float64); ok {
		return comparePackets(left, []interface{}{rightInt})
	}

	return 0
}

func countSortedPackets(input string) int {
	sortedPackets := 0
	pairs := strings.Split(input, "\n\n")
	for i, pair_str := range pairs {
		pair := strings.Split(pair_str, "\n")

		left := []interface{}{}
		err := json.Unmarshal([]byte(pair[0]), &left)
		if err != nil {
			fmt.Printf("Error parsing %s: %s\n", input, err)
		}

		right := []interface{}{}
		err = json.Unmarshal([]byte(pair[1]), &right)
		if err != nil {
			fmt.Printf("Error parsing %s: %s\n", input, err)
		}

		if comparePackets(left, right) == -1 {
			sortedPackets += i + 1
		}
	}

	return sortedPackets
}

func sortAndDividePackets(input string) int {
	// Remove blank lines from input
	input = strings.ReplaceAll(input, "\n\n", "\n")
	packetStrings := strings.Split(input+"\n[[2]]\n[[6]]", "\n")

	packets := make([][]interface{}, len(packetStrings))

	// Convert packets to lists
	for i, packet := range packetStrings {
		err := json.Unmarshal([]byte(packet), &packets[i])
		if err != nil {
			fmt.Printf("Error parsing %s: %s\n", input, err)
		}
	}

	// sort packets using comparePackets
	for i := 0; i < len(packets); i++ {
		for j := i + 1; j < len(packets); j++ {
			if comparePackets(packets[i], packets[j]) == 1 {
				packets[i], packets[j] = packets[j], packets[i]
			}
		}
	}

	divisors := []int{}
	// convert packets back to strings and find the two rows containing [[2]] and [[6]]
	for i, packet := range packets {
		packetBytes, err := json.Marshal(packet)
		if err != nil {
			fmt.Printf("Error parsing %s: %s\n", input, err)
		}
		packetString := string(packetBytes)
		packetStrings[i] = packetString

		if packetString == "[[2]]" {
			divisors = append(divisors, i+1)
		} else if packetString == "[[6]]" {
			divisors = append(divisors, i+1)
		}
	}

	return divisors[0] * divisors[1]
}
