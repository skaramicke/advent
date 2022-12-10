package day20221210

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/skaramicke/advent/utils"
)

func Run() {
	input := utils.ReadOrCreateInputFile("2022-12-10")
	strength_sum, output := read_cpu(input)
	fmt.Printf("Task 1 result: %d\n", strength_sum)
	fmt.Printf("Task 2 result:\n\n%s\n\n", output)
}

func draw(x, cycles int, output string) string {
	if x-1 <= (cycles)%40 && x+1 >= (cycles)%40 {
		output += "#"
	} else {
		output += " "
	}

	return output
}

func next_cycle(checks, strengths []int, output string, cycles, x int) (int, []int, string) {
	cycles++
	for _, check := range checks {
		if cycles == check {
			strengths = append(strengths, x*cycles)
		}
	}

	return cycles, strengths, output
}

func crt(buffer string, width int) string {
	numChunks := len(buffer) / width
	if len(buffer)%width > 0 {
		numChunks++
	}
	chunks := make([]string, 0, numChunks)

	for i := 0; i < len(buffer); i += width {
		chunk := ""
		if i+width > len(buffer) {
			chunk = buffer[i:]
		} else {
			chunk = buffer[i : i+width]
		}
		chunks = append(chunks, chunk)
	}

	output := "╔"
	// top frame
	for i := 0; i < width+2; i++ {
		output += "═"
	}
	output += "╗\n"
	for _, chunk := range chunks {
		output += fmt.Sprintf("║ %s ║\n", chunk)
	}
	output += "╚"
	// bottom frame
	for i := 0; i < width+2; i++ {
		output += "═"
	}
	output += "╝\n"

	return output
}

func read_cpu(input string) (int, string) {
	lines := strings.Split(input, "\n")
	output := ""
	x := 1
	cycles := 0
	strengths := []int{}

	checks := []int{
		20, 60, 100, 140, 180, 220,
	}

	for _, line := range lines {
		cycles, strengths, output = next_cycle(checks, strengths, output, cycles, x)
		output = draw(x, cycles, output)
		if strings.HasPrefix(line, "addx") {
			cycles, strengths, output = next_cycle(checks, strengths, output, cycles, x)
			value, _ := strconv.Atoi(strings.Split(line, " ")[1])
			x += value
			output = draw(x, cycles, output)
		}
	}
	strength_sum := 0
	for _, strength := range strengths {
		strength_sum += strength
	}

	crt := crt(output, 40)
	return strength_sum, crt
}
