package day20221208

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/skaramicke/advent/utils"
)

func Run() {
	input := utils.ReadOrCreateInputFile("2022-12-08")
	fmt.Printf("Task 1 result: %d\n", task_1(input))
	fmt.Printf("Task 2 result: %d\n", task_2(input))
}

func task_1(input string) int {
	visiblePositions := map[string]bool{}

	// loop from left to right
	for i := 0; i < len(strings.Split(input, "\n")[0]); i++ {
		highestFoundTree := -1
		// loop from top to bottom
		for j := 0; j < len(strings.Split(input, "\n")); j++ {
			pos := fmt.Sprintf("%dx%d", i, j)
			tree, err := strconv.Atoi(string(strings.Split(input, "\n")[j][i]))
			if err != nil {
				panic(err)
			}
			if tree > highestFoundTree {
				highestFoundTree = tree
				visiblePositions[pos] = true
			}
		}

		highestFoundTree = -1
		// loop from bottom to top
		for j := len(strings.Split(input, "\n")) - 1; j >= 0; j-- {
			pos := fmt.Sprintf("%dx%d", i, j)
			tree, err := strconv.Atoi(string(strings.Split(input, "\n")[j][i]))
			if err != nil {
				panic(err)
			}
			if tree > highestFoundTree {
				highestFoundTree = tree
				visiblePositions[pos] = true
			}
		}
	}

	// loop from top to bottom
	for j := 0; j < len(strings.Split(input, "\n")); j++ {
		highestFoundTree := -1
		// loop from left to right
		for i := 0; i < len(strings.Split(input, "\n")[0]); i++ {
			tree, err := strconv.Atoi(string(strings.Split(input, "\n")[j][i]))
			if err != nil {
				panic(err)
			}
			if tree > highestFoundTree {
				highestFoundTree = tree
				visiblePositions[fmt.Sprintf("%dx%d", i, j)] = true
			}
		}

		highestFoundTree = -1
		// loop from right to left
		for i := len(strings.Split(input, "\n")[0]) - 1; i >= 0; i-- {
			tree, err := strconv.Atoi(string(strings.Split(input, "\n")[j][i]))
			if err != nil {
				panic(err)
			}
			if tree > highestFoundTree {
				highestFoundTree = tree
				visiblePositions[fmt.Sprintf("%dx%d", i, j)] = true
			}
		}
	}

	return len(visiblePositions)
}

func task_2(input string) int {
	max_x := len(strings.Split(input, "\n")[0])
	max_y := len(strings.Split(input, "\n"))
	scenicScores := map[string]int{}
	for x := 0; x < max_x; x++ {
		for y := 0; y < max_y; y++ {

			fmt.Printf("\nChecking %dx%d\n", x, y)

			tree, err := strconv.Atoi(string(strings.Split(input, "\n")[y][x]))
			if err != nil {
				panic(err)
			}
			right := 0
			left := 0
			up := 0
			down := 0
			// look to the right
			if x < max_x {
				for look_x := x + 1; look_x < max_x; look_x++ {
					ltree, err := strconv.Atoi(string(strings.Split(input, "\n")[y][look_x]))
					if err != nil {
						panic(err)
					}
					fmt.Printf("Looking right from %dx%d (%d) at %dx%d: %d", x, y, tree, look_x, y, ltree)
					right++
					if ltree >= tree {
						fmt.Printf(" - break\n")
						break
					}
					fmt.Printf(" - ok\n")
				}
			}

			// look to the left
			if x > 0 {
				for look_x := x - 1; look_x >= 0; look_x-- {
					ltree, err := strconv.Atoi(string(strings.Split(input, "\n")[y][look_x]))
					if err != nil {
						panic(err)
					}
					fmt.Printf("Looking left from %dx%d (%d) at %dx%d: %d", x, y, tree, look_x, y, ltree)
					left++
					if ltree >= tree {
						fmt.Printf(" - break\n")
						break
					}
					fmt.Printf(" - ok\n")
				}
			}

			// look up
			if y > 0 {
				for look_y := y - 1; look_y >= 0; look_y-- {
					ltree, err := strconv.Atoi(string(strings.Split(input, "\n")[look_y][x]))
					if err != nil {
						panic(err)
					}
					fmt.Printf("Looking up from %dx%d (%d) at %dx%d: %d", x, y, tree, x, look_y, ltree)
					up++
					if ltree >= tree {
						fmt.Printf(" - break\n")

						break
					}
					fmt.Printf(" - ok\n")

				}
			}

			// look down
			if y < max_y {
				for look_y := y + 1; look_y < max_y; look_y++ {
					ltree, err := strconv.Atoi(string(strings.Split(input, "\n")[look_y][x]))
					if err != nil {
						panic(err)
					}
					fmt.Printf("Looking down from %dx%d (%d) at %dx%d: %d", x, y, tree, x, look_y, ltree)
					down++
					if ltree >= tree {
						fmt.Printf(" - break\n")

						break
					}
					fmt.Printf(" - ok\n")

				}
			}
			fmt.Printf("Scenic score for %dx%d: %d * %d * %d * %d = %d\n", x, y, right, left, up, down, right*left*up*down)
			scenicScores[fmt.Sprintf("%dx%d", x, y)] = right * left * up * down

		}
	}

	// return highest scenicScores
	highest := 0
	for _, v := range scenicScores {
		if v > highest {
			highest = v
		}
	}

	return highest
}
