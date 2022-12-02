package main

import (
	"github.com/skaramicke/advent/day20221201"
	"github.com/skaramicke/advent/day20221202"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{
			"December 2",
			"December 1",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		println(err.Error())
		return
	}

	switch result {
	case "December 1":
		day20221201.Prompt()
	case "December 2":
		day20221202.Prompt()
	}
}
