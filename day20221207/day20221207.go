package day20221207

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/skaramicke/advent/utils"
)

func Run() {
	input := utils.ReadOrCreateInputFile("2022-12-07")
	result1, result2 := processDirectories(input)
	fmt.Printf("Task 1 result: %d\n", result1)
	fmt.Printf("Task 2 result: %d\n", result2)
}

func processDirectories(input string) (int, int) {
	lines := strings.Split(input, "\n")
	directoryStack := []string{}
	directorySizes := map[string]int{}
	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd") {
			if strings.Contains(line, "..") {
				directoryStack = directoryStack[:len(directoryStack)-1]
			} else {
				directoryStack = append(directoryStack, strings.Split(line, " ")[2])
				directorySizes[strings.Join(directoryStack, "/")] = 0
			}
		}

		size := strings.Split(line, " ")[0]
		if sizeInt, err := strconv.Atoi(size); err == nil {
			path := strings.Join(directoryStack, "/")
			directorySizes[path] += sizeInt
		}
	}

	// Find the total size of a directory and its subdirectories
	recursiveDirSizes := map[string]int{}
	for dir := range directorySizes {
		sum := 0
		for directory, size := range directorySizes {
			if strings.HasPrefix(directory, dir) {
				sum += size
			}
		}
		recursiveDirSizes[dir] = sum
	}

	// Find directories with at most 100000 bytes
	smallDirectorySum := 0
	for _, size := range recursiveDirSizes {
		if size <= 100000 {
			smallDirectorySum += size
		}
	}

	requiredRemoval := 30000000 - (70000000 - recursiveDirSizes["/"])
	smallestDirSizeOverLimit := recursiveDirSizes["/"]
	for _, size := range recursiveDirSizes {
		if size > requiredRemoval && size < smallestDirSizeOverLimit {
			smallestDirSizeOverLimit = size
		}
	}

	return smallDirectorySum, smallestDirSizeOverLimit
}
