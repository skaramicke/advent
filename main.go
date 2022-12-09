package main

import (
	"github.com/skaramicke/advent/day20221201"
	"github.com/skaramicke/advent/day20221202"
	"github.com/skaramicke/advent/day20221203"
	"github.com/skaramicke/advent/day20221204"
	"github.com/skaramicke/advent/day20221205"
	"github.com/skaramicke/advent/day20221206"
	"github.com/skaramicke/advent/day20221207"
	"github.com/skaramicke/advent/day20221208"
	"github.com/skaramicke/advent/day20221209"
	// new imports go here

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{
			// new options go here
			"December 09, 2022",
			"December 08, 2022",
			"December 07, 2022",
			"December 6, 2022",
			"December 5, 2022",
			"December 4, 2022",
			"December 3, 2022",
			"December 2, 2022",
			"December 1, 2022",
		},
	}

	_, result, err := prompt.Run()

	if err != nil {
		println(err.Error())
		return
	}

	switch result {
	case "December 1, 2022":
		day20221201.Run()
	case "December 2, 2022":
		day20221202.Run()
	case "December 3, 2022":
		day20221203.Run()
	case "December 4, 2022":
		day20221204.Run()
	case "December 5, 2022":
		day20221205.Run()
	case "December 6, 2022":
		day20221206.Run()
	case "December 07, 2022":
		day20221207.Run()
	case "December 08, 2022":
		day20221208.Run()
	case "December 09, 2022":
		day20221209.Run()
		// new calls go here
	}
}
