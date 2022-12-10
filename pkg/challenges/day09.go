package challenges

import (
	"AOC-2022/pkg/utils"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

func Day09() {
	input := utils.ImportInputLines(9)

	// Yeah, I initialized TWO 1000x1000 matrices for this challenge. Sue me.
	visited := make(map[position]any)
	visited2 := make(map[position]any)

	rope := make([]position, 10)

	for _, line := range input {
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		distance, _ := strconv.Atoi(splitLine[1])
		for ; distance > 0; distance-- {
			switch direction {
			case "U":
				rope[0].y++
			case "D":
				rope[0].y--
			case "R":
				rope[0].x++
			case "L":
				rope[0].x--
			}
			for i := 1; i < 10; i++ {
				switch rope[i-1] {
				case position{rope[i].x + 2, rope[i].y}:
					rope[i].x++
				case position{rope[i].x - 2, rope[i].y}:
					rope[i].x--
				case position{rope[i].x, rope[i].y + 2}:
					rope[i].y++
				case position{rope[i].x, rope[i].y - 2}:
					rope[i].y--
				case position{rope[i].x + 2, rope[i].y + 2}:
					fallthrough
				case position{rope[i].x + 2, rope[i].y + 1}:
					fallthrough
				case position{rope[i].x + 1, rope[i].y + 2}:
					rope[i].x++
					rope[i].y++
				case position{rope[i].x + 2, rope[i].y - 2}:
					fallthrough
				case position{rope[i].x + 2, rope[i].y - 1}:
					fallthrough
				case position{rope[i].x + 1, rope[i].y - 2}:
					rope[i].x++
					rope[i].y--
				case position{rope[i].x - 2, rope[i].y + 2}:
					fallthrough
				case position{rope[i].x - 2, rope[i].y + 1}:
					fallthrough
				case position{rope[i].x - 1, rope[i].y + 2}:
					rope[i].x--
					rope[i].y++
				case position{rope[i].x - 2, rope[i].y - 2}:
					fallthrough
				case position{rope[i].x - 2, rope[i].y - 1}:
					fallthrough
				case position{rope[i].x - 1, rope[i].y - 2}:
					rope[i].x--
					rope[i].y--
				}
			}
			visited[rope[1]] = nil
			visited2[rope[9]] = nil
		}
	}

	println(len(visited))
	println(len(visited2))
}
