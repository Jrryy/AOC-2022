package challenges

import "AOC-2022/pkg/utils"

func Day02() {
	input := utils.ImportInputLines(2)
	input = input[:len(input)-1]

	points := 0
	points2 := 0

	for _, round := range input {
		opponent := round[0]
		player := round[2]
		if opponent == 'A' {
			if player == 'X' {
				points += 4
				points2 += 3
			}
			if player == 'Y' {
				points += 8
				points2 += 4
			}
			if player == 'Z' {
				points += 3
				points2 += 8
			}
		} else if opponent == 'B' {
			if player == 'X' {
				points += 1
				points2 += 1
			}
			if player == 'Y' {
				points += 5
				points2 += 5
			}
			if player == 'Z' {
				points += 9
				points2 += 9
			}
		} else if opponent == 'C' {
			if player == 'X' {
				points += 7
				points2 += 2
			}
			if player == 'Y' {
				points += 2
				points2 += 6
			}
			if player == 'Z' {
				points += 6
				points2 += 7
			}
		}
	}

	println(points)
	println(points2)
}

/*
// Approach with switches. Faster than maps, but unsurprisingly slower than ifs. This makes me mad.
func Day02() {
	input := utils.ImportInputLines(2)
	input = input[:len(input)-1]

	points := 0
	points2 := 0

	for _, round := range input {
		switch round[0] {
		case 'A':
			switch round[2] {
			case 'X':
				points += 4
				points2 += 3
			case 'Y':
				points += 8
				points2 += 4
			case 'Z':
				points += 3
				points2 += 8
			}
		case 'B':
			switch round[2] {
			case 'X':
				points += 1
				points2 += 1
			case 'Y':
				points += 5
				points2 += 5
			case 'Z':
				points += 9
				points2 += 9
			}
		case 'C':
			switch round[2] {
			case 'X':
				points += 7
				points2 += 2
			case 'Y':
				points += 2
				points2 += 6
			case 'Z':
				points += 6
				points2 += 7
			}
		}

	}

	println(points)
	println(points2)
}
*/
/*
// Approach using maps. It is, of course, much slower than the above ones,
// but much more clean and simple.
func Day02() {
	input := utils.ImportInputLines(2)

	points := 0
	points2 := 0

	roundsMap := map[string]struct {
		first  int
		second int
	}{
		"A X": {4, 3},
		"A Y": {8, 4},
		"A Z": {3, 8},
		"B X": {1, 1},
		"B Y": {5, 5},
		"B Z": {9, 9},
		"C X": {7, 2},
		"C Y": {2, 6},
		"C Z": {6, 7},
	}

	for _, round := range input {
		results := roundsMap[round]
		points += results.first
		points2 += results.second
	}

	println(points)
	println(points2)
}
*/
