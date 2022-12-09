package day20221209

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/skaramicke/advent/utils"
)

func Run() {
	input := utils.ReadOrCreateInputFile("2022-12-09")
	fmt.Printf("Task 1 result: %d\n", simulateRope(input, 2))
	fmt.Printf("Task 2 result: %d\n", simulateRope(input, 10))
}

func simulateRope(input string, knots int) int {
	tailPositions := map[string]bool{
		"0x0": true,
	}
	positions := map[int][2]int{}
	for i := 0; i < knots; i++ {
		positions[i] = [2]int{0, 0}
	}

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for _, line := range strings.Split(input, "\n") {
		direction := string(line[0])
		distance, err := strconv.Atoi(line[2:])
		if err != nil {
			panic(err)
		}

		targetPosition := [2]int{positions[0][0], positions[0][1]}
		switch direction {
		case "R":
			targetPosition[0] += distance
		case "U":
			targetPosition[1] -= distance
		case "L":
			targetPosition[0] -= distance
		case "D":
			targetPosition[1] += distance
		}

		for positions[0][0] != targetPosition[0] || positions[0][1] != targetPosition[1] {
			if positions[0][0] < targetPosition[0] {
				positions[0] = [2]int{positions[0][0] + 1, positions[0][1]}
			} else if positions[0][0] > targetPosition[0] {
				positions[0] = [2]int{positions[0][0] - 1, positions[0][1]}
			}

			if positions[0][1] < targetPosition[1] {
				positions[0] = [2]int{positions[0][0], positions[0][1] + 1}
			} else if positions[0][1] > targetPosition[1] {
				positions[0] = [2]int{positions[0][0], positions[0][1] - 1}
			}

			if positions[0][0] < minX {
				minX = positions[0][0]
			}
			if positions[0][0] > maxX {
				maxX = positions[0][0]
			}
			if positions[0][1] < minY {
				minY = positions[0][1]
			}
			if positions[0][1] > maxY {
				maxY = positions[0][1]
			}

			// Move the rest of the knots in order
			for i := 0; i < knots-1; i++ {
				head_x, head_y := positions[i][0], positions[i][1]
				tail_x, tail_y := positions[i+1][0], positions[i+1][1]
				dist_x, dist_y := head_x-tail_x, head_y-tail_y
				if utils.Abs(dist_x) >= 2 || utils.Abs(dist_y) >= 2 {
					dist_x = utils.Max(-1, utils.Min(1, dist_x))
					dist_y = utils.Max(-1, utils.Min(1, dist_y))
					positions[i+1] = [2]int{tail_x + dist_x, tail_y + dist_y}
				}
			}

			tailPositions[fmt.Sprintf("%dx%d", positions[knots-1][0], positions[knots-1][1])] = true
		}
	}
	// drawPositions(positions, tailPositions, minX, maxX, minY, maxY)
	return len(tailPositions)
}

func drawPositions(positions map[int][2]int, tailPositions map[string]bool, minX, maxX, minY, maxY int) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if x == 0 && y == 0 {
				fmt.Print("s")
			} else {
				found := false
				for i, position := range positions {
					if position[0] == x && position[1] == y {
						if i == 0 {
							fmt.Print("H")
						} else {
							fmt.Print(i)
						}
						found = true
						break
					}
				}
				if !found {
					for position := range tailPositions {
						if position == fmt.Sprintf("%dx%d", x, y) {
							fmt.Print("#")
							found = true
							break
						}
					}
				}
				if !found {
					fmt.Print(".")
				}
			}
		}
		fmt.Printf("  %d", y)

		if y >= 0 && y < len(positions) {
			fmt.Printf("  %dx%d", positions[y][0], positions[y][1])
		}
		fmt.Println()
	}
	fmt.Println()
}
