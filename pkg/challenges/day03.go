package challenges

import (
	"AOC-2022/pkg/utils"
	"strings"
)

func Day03() {
	priorities := " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	totalPriority := 0
	totalPriority2 := 0
	group := 1
	commonItems := ""
	input := utils.ImportInputLines(3)
	for _, rucksack := range input {

		// Part 1: Find common in both halves of the rucksack
		compartment1, compartment2 := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]
		for _, item := range compartment1 {
			if strings.ContainsRune(compartment2, item) {
				totalPriority += strings.IndexRune(priorities, item)
				break
			}
		}

		// Part 2: Find common in 3 consecutive rucksacks
		switch group {
		case 1:
			commonItems = rucksack
			group++
		case 2:
			for _, item := range commonItems {
				if !strings.ContainsRune(rucksack, item) {
					commonItems = strings.ReplaceAll(commonItems, string(item), "")
				}
			}
			group++
		case 3:
			for _, item := range commonItems {
				if strings.ContainsRune(rucksack, item) {
					totalPriority2 += strings.IndexRune(priorities, item)
					break
				}
			}
			commonItems = ""
			group = 1
		}
	}
	println(totalPriority)
	println(totalPriority2)
}
