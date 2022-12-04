package utils

import (
	"fmt"
	"os"
)

func ReadOrCreateInputFile(day string) string {
	filename := fmt.Sprintf("inputs/%s.txt", day)
	if input, err := os.ReadFile(filename); err == nil {
		return string(input)
	} else {
		// If error, create file and ask user to try again
		os.WriteFile(filename, []byte{}, 0644)
		fmt.Printf("please try again after adding input to %s", filename)
		panic("input file not found")
	}
}
