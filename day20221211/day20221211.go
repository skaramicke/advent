package day20221211

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/skaramicke/advent/utils"
)

func Run() {
	input := utils.ReadOrCreateInputFile("2022-12-11")
	fmt.Printf("Easy monkey business: %d\n", findMonkeyBusiness(input, true))
	fmt.Printf("Extra worrying monkey business: %d\n", findMonkeyBusiness(input, false))
}

type Monkey struct {
	ID          int
	Items       []int
	Operation   string
	TestDivisor int
	TrueMonkey  int
	FalseMonkey int
}

func findMonkeyBusiness(input string, ease bool) int {

	monkeyStrings := strings.Split(input, "\n\n")
	monkeys := []Monkey{}
	supermodulus := 1

	re := regexp.MustCompile(`Monkey (\d+):\n.*Starting items: (.+)\n.*Operation: new = old (.+)\n.*divisible by (\d+)\n.*true: throw to monkey (\d+)\n.*false: throw to monkey (\d+)$`)
	for _, monkeyString := range monkeyStrings {
		data := re.FindStringSubmatch(monkeyString)

		id, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}

		items := []int{}
		for _, item := range strings.Split(data[2], ", ") {
			item, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			items = append(items, item)
		}

		testDivisor, err := strconv.Atoi(data[4])
		if err != nil {
			panic(err)
		}

		supermodulus *= testDivisor

		trueMonkey, err := strconv.Atoi(data[5])
		if err != nil {
			panic(err)
		}

		falseMonkey, err := strconv.Atoi(data[6])
		if err != nil {
			panic(err)
		}

		monkey := Monkey{
			ID:          id,
			Items:       items,
			Operation:   data[3],
			TestDivisor: testDivisor,
			TrueMonkey:  trueMonkey,
			FalseMonkey: falseMonkey,
		}

		monkeys = append(monkeys, monkey)
	}

	inspectionCounters := make([]int, len(monkeys))

	rounds := 20
	if !ease {
		rounds = 10000
	}

	// Uncomment to print inspection counters after certain rounds. I used it for debugging.
	// I noticed the int overflow at round 1000, so I added the supermodulo trick to get rid of it.
	// checks := []int{
	// 	1, 20, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000,
	// }

	for round := 1; round <= rounds; round++ {
		for i, monkey := range monkeys {
			operation := strings.Split(monkey.Operation, " ")
			for _, item := range monkey.Items {
				value := item
				if _, err := strconv.Atoi(operation[1]); err == nil {
					value, _ = strconv.Atoi(operation[1])
				}

				switch operation[0] {
				case "+":
					item += value
				case "-":
					item -= value
				case "*":
					item *= value
				case "/":
					item /= value
				}

				if ease {
					item /= 3
				}

				item %= supermodulus

				if item%monkey.TestDivisor == 0 {
					monkeys[monkey.TrueMonkey].Items = append(monkeys[monkey.TrueMonkey].Items, item)
				} else {
					monkeys[monkey.FalseMonkey].Items = append(monkeys[monkey.FalseMonkey].Items, item)
				}
			}

			inspectionCounters[i] += len(monkey.Items)
			monkey.Items = []int{}
			monkeys[i] = monkey
		}

		// Uncomment to print inspection counters after certain rounds. I used it for debugging.
		// for i := range checks {
		// 	if round == checks[i] {
		// 		fmt.Printf("\n== After round %d ==\n", round)
		// 		for monkey, counter := range inspectionCounters {
		// 			fmt.Printf("Monkey %d inspected items %d times.\n", monkey, counter)
		// 		}
		// 		break
		// 	}
		// }

	}

	// Sort the list of inspection counters
	sort.Slice(inspectionCounters, func(i, j int) bool {
		return inspectionCounters[i] > inspectionCounters[j]
	})
	result := inspectionCounters[0] * inspectionCounters[1]

	return result
}
