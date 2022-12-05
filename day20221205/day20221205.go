package day20221205

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/skaramicke/advent/utils"
)

func Run() {
	input := utils.ReadOrCreateInputFile("2022-12-05")

	topCrate := findTopCrate9000(input)

	fmt.Printf("Result by 9000: %s\n", topCrate)

	topCrate = findTopCrate9001(input)
	fmt.Printf("Result by 9001: %s\n", topCrate)
}

func readCrateStacks(input string) ([][]string, int) {
	// Find bottom index of crate lines
	lines := strings.Split(input, "\n")
	stackNumbersIndex := 0
	for i, line := range lines {
		if !strings.Contains(line, "[") {
			stackNumbersIndex = i
			break
		}
	}

	// Read crate characters from bottom up
	stacks := [][]string{}

	for i := stackNumbersIndex - 1; i >= 0; i-- {
		cratesOnRow := []string{}
		// Every third character is a crate
		for j := 0; j < len(lines[i]); j += 4 {
			crate := strings.Trim(string(lines[i][j+1:j+2]), " ")
			cratesOnRow = append(cratesOnRow, crate)
		}
		for k, crate := range cratesOnRow {
			if len(stacks) <= k {
				stacks = append(stacks, []string{})
			}
			if crate != "" {
				stacks[k] = append(stacks[k], crate)
			}
		}
	}
	return stacks, stackNumbersIndex + 2
}

func findTopCrate9000(input string) string {
	// Read crates
	stacks, movesIndex := readCrateStacks(input)

	lines := strings.Split(input, "\n")

	// Read moves
	for _, move := range lines[movesIndex:] {
		// move string: "move [number of crates] form [stack] to [stack]"
		moveParts := strings.Fields(move)
		numCrates, err := strconv.Atoi(moveParts[1])
		if err != nil {
			panic(err)
		}
		fromStack, err := strconv.Atoi(moveParts[3])
		fromStack -= 1
		if err != nil {
			panic(err)
		}
		toStack, err := strconv.Atoi(moveParts[5])
		toStack -= 1
		if err != nil {
			panic(err)
		}

		// Move crates
		for i := 0; i < numCrates; i++ {
			stacks[toStack] = append(stacks[toStack], stacks[fromStack][len(stacks[fromStack])-1])
			stacks[fromStack] = stacks[fromStack][:len(stacks[fromStack])-1]
		}
	}

	// Find highest stack
	highestStack := 0
	for i, stack := range stacks {
		if len(stack) > highestStack {
			highestStack = i
		}
	}

	// Find top crate in each stack and return the string of characters
	topStackString := ""
	for _, stack := range stacks {
		topStackString += stack[len(stack)-1]
	}
	return topStackString
}

func findTopCrate9001(input string) string {
	// Read crates
	stacks, movesIndex := readCrateStacks(input)

	lines := strings.Split(input, "\n")

	// Read moves
	for _, move := range lines[movesIndex:] {
		// move string: "move [number of crates] form [stack] to [stack]"
		moveParts := strings.Fields(move)
		numCrates, err := strconv.Atoi(moveParts[1])
		if err != nil {
			panic(err)
		}
		fromStack, err := strconv.Atoi(moveParts[3])
		fromStack -= 1
		if err != nil {
			panic(err)
		}
		toStack, err := strconv.Atoi(moveParts[5])
		toStack -= 1
		if err != nil {
			panic(err)
		}

		// Move crates
		// Find numCrates in fromStack
		cratesToMove := stacks[fromStack][len(stacks[fromStack])-numCrates:]
		stacks[fromStack] = stacks[fromStack][:len(stacks[fromStack])-numCrates]
		stacks[toStack] = append(stacks[toStack], cratesToMove...)
	}

	// Find highest stack
	highestStack := 0
	for i, stack := range stacks {
		if len(stack) > highestStack {
			highestStack = i
		}
	}

	// Find top crate in each stack and return the string of characters
	topStackString := ""
	for _, stack := range stacks {
		topStackString += stack[len(stack)-1]
	}
	return topStackString
}
