package day20221212

import (
	"fmt"
	"strings"

	"github.com/skaramicke/advent/utils"
)

func Run() {
	input := utils.ReadOrCreateInputFile("2022-12-12")
	fmt.Printf("Task 1 result: %d\n", task_1(input))
	fmt.Printf("Task 2 result: %d\n", task_2(input))
}

type Location struct {
	X int
	Y int
}

func task_1(input string) int {

	start := Location{0, 0}
	end := Location{0, 0}

	area := map[int]map[int]int{}
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if char == 'S' {
				start = Location{x, y}
				continue
			} else if char == 'E' {
				end = Location{x, y}
				char = 'z'
			}
			if _, ok := area[x]; !ok {
				area[x] = map[int]int{}
			}
			area[x][y] = int(char) - 97
		}
	}

	distances := map[int]map[int]int{}
	findNextDistance(start.X, start.Y, 0, area, distances)

	return distances[end.X][end.Y]
}

func findNextDistance(x, y, distance int, area, distances map[int]map[int]int) {

	for _, dir := range []Location{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
		if _, ok := area[x+dir.X]; ok {
			if _, ok := area[x+dir.X][y+dir.Y]; ok {
				if area[x+dir.X][y+dir.Y]-area[x][y] > 1 {
					continue
				}
				if _, ok := distances[x+dir.X]; !ok {
					distances[x+dir.X] = map[int]int{}
				}
				if _, ok := distances[x+dir.X][y+dir.Y]; !ok || distances[x+dir.X][y+dir.Y] > distance+1 {
					distances[x+dir.X][y+dir.Y] = distance + 1
					findNextDistance(x+dir.X, y+dir.Y, distance+1, area, distances)
				}
			}
		}
	}

}

// Second failed attempt:

// func traverse(area map[int]map[int]int, end Location, path []Location) []Location {
// 	// For each neighbor of the current node
// 	current_node := path[len(path)-1]

// 	shortest_path := []Location{}

// 	for _, neighbour := range []Location{
// 		{current_node.X - 1, current_node.Y},
// 		{current_node.X + 1, current_node.Y},
// 		{current_node.X, current_node.Y - 1},
// 		{current_node.X, current_node.Y + 1},
// 	} {
// 		debug := current_node.X == 0 && current_node.Y == 0

// 		// If the neighbor has not been explored
// 		if _, ok := area[neighbour.X][neighbour.Y]; ok {

// 			found := false
// 			for _, path_node := range path {
// 				if path_node.X == neighbour.X && path_node.Y == neighbour.Y {
// 					found = true
// 					break
// 				}
// 			}
// 			if found {
// 				continue
// 			}

// 			// If the height difference is 1 or less
// 			if area[neighbour.X][neighbour.Y]-area[current_node.X][current_node.Y] > 1 {
// 				continue
// 			}

// 			// If the neighbor is the end node
// 			if neighbour.X == end.X && neighbour.Y == end.Y {
// 				// Found path
// 				return path
// 			}

// 			found_path := traverse(area, end, append(path, neighbour))
// 			if debug {
// 				fmt.Printf("%v\n", found_path)
// 			}
// 			if len(found_path) > 0 && (len(shortest_path) == 0 || len(found_path) < len(shortest_path)) {
// 				shortest_path = found_path
// 			}
// 		} else if debug {
// 			fmt.Printf("Not in area: %v\n", neighbour)
// 		}
// 	}
// 	return shortest_path
// }

// First failed attempt:

// // Recursive function to find the end location
// func traverse(width, height int, area map[int]int, target Location, path []Location) (int, []Location) {
// 	debug := false
// 	here := path[len(path)-1]
// 	here_height := area[width*here.Y+here.X]

// 	found_paths := [][]Location{}
// 	for _, neighbour := range []Location{
// 		{here.X - 1, here.Y},
// 		{here.X + 1, here.Y},
// 		{here.X, here.Y - 1},
// 		{here.X, here.Y + 1},
// 	} {
// 		if debug {
// 			fmt.Printf("%v. Checking %v: ", here, neighbour)
// 		}
// 		if neighbour == target {
// 			// Done! :D
// 			if debug {
// 				fmt.Printf("Found target %v\n", target)
// 			}
// 			path = append(path, neighbour)
// 			return len(path), path
// 		}
// 		if neighbour.X < 0 || neighbour.Y < 0 || neighbour.X >= width || neighbour.Y >= height {
// 			// Wall
// 			if debug {
// 				fmt.Printf("Out of bounds\n")
// 			}
// 			continue
// 		}
// 		found := false
// 		for _, p := range path {
// 			if p == neighbour {
// 				// Already visited
// 				if debug {
// 					fmt.Printf("Already visited\n")
// 				}
// 				found = true
// 				break
// 			}
// 		}
// 		if found {
// 			continue
// 		}
// 		neighbour_height := area[width*neighbour.Y+neighbour.X]
// 		if utils.Abs(neighbour_height-here_height) > 1 {
// 			// Too different in height
// 			if debug {
// 				fmt.Printf("Too different in height, here: %d, there: %d\n", here_height, neighbour_height)
// 			}
// 			continue
// 		}

// 		// Add neighbour to path
// 		if debug {
// 			fmt.Printf("Adding neighbour %v to path\n", neighbour)
// 		}
// 		path = append(path, neighbour)

// 		steps, found_path := traverse(width, height, area, target, path)
// 		if steps > 0 {
// 			found_paths = append(found_paths, found_path)
// 		}
// 	}

// 	if len(found_paths) > 0 {
// 		min := len(area)
// 		min_path := []Location{}
// 		for _, p := range found_paths {
// 			if len(p) < min {
// 				min = len(p)
// 				min_path = p
// 			}
// 		}
// 		return min, min_path
// 	}

// 	return -1, []Location{}
// }

func task_2(input string) int {
	return 0
}
