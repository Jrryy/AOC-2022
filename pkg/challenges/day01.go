package challenges

import (
	"AOC-2022/pkg/utils"
	"strconv"
)

func Day01() {
	input := utils.ImportInputLines(1)

	maxCalories := 0
	maxCalories2 := 0
	maxCalories3 := 0
	currentCalories := 0
	for _, rawCalories := range input {
		if rawCalories == "" {
			switch {
			case currentCalories > maxCalories:
				maxCalories, currentCalories = currentCalories, maxCalories
				fallthrough
			case currentCalories > maxCalories2:
				maxCalories2, currentCalories = currentCalories, maxCalories2
				fallthrough
			case currentCalories > maxCalories3:
				maxCalories3 = currentCalories
			}
			currentCalories = 0
			continue
		}
		calories, err := strconv.Atoi(rawCalories)
		if err != nil {
			panic(err)
		}
		currentCalories += calories
	}

	println(maxCalories)
	println(maxCalories + maxCalories2 + maxCalories3)
}

/*
// Sorted approach, which as I expected is slower than the above one (at least with my input)
func Day01() {
	input := utils.ImportInputLines(1)

	currentCalories := 0
	var calsPerElf []int
	for _, rawCalories := range input {
		if rawCalories == "" {
			calsPerElf = append(calsPerElf, currentCalories)
			currentCalories = 0
			continue
		}
		calories, err := strconv.Atoi(rawCalories)
		if err != nil {
			panic(err)
		}
		currentCalories += calories
	}

	sort.Ints(calsPerElf)
	n := len(calsPerElf)

	println(calsPerElf[n-1])
	println(calsPerElf[n-1] + calsPerElf[n-2] + calsPerElf[n-3])
}
*/
