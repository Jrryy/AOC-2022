package challenges

import (
	"AOC-2022/pkg/utils"
)

func Day02() {
	input := utils.ImportInputLines(2)

	points := 0
	points2 := 0

	roundsMap := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	roundsMap2 := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	for _, round := range input {
		points += roundsMap[round]
		points2 += roundsMap2[round]
	}

	println(points)
	println(points2)
}
