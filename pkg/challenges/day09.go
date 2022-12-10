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

	// Yeah, I initialized TWO 700x700 matrices for this challenge. Sue me.
	visited := make([][]bool, 700)
	for i := range visited {
		visited[i] = make([]bool, 700)
	}

	visited2 := make([][]bool, 700)
	for i := range visited2 {
		visited2[i] = make([]bool, 700)
	}

	rope := make([]position, 10)
	for i := range rope {
		rope[i] = position{350, 350}
	}

	totalVisited := 0
	totalVisited2 := 0

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
				dx := rope[i].x - rope[i-1].x
				dy := rope[i].y - rope[i-1].y
				if dx <= 1 && dx >= -1 && dy <= 1 && dy >= -1 {
					break
				}
				if dx > 0 {
					rope[i].x--
				} else if dx < 0 {
					rope[i].x++
				}
				if dy > 0 {
					rope[i].y--
				} else if dy < 0 {
					rope[i].y++
				}
			}
			if !visited[rope[1].x][rope[1].y] {
				visited[rope[1].x][rope[1].y] = true
				totalVisited++
			}
			if !visited2[rope[9].x][rope[9].y] {
				visited2[rope[9].x][rope[9].y] = true
				totalVisited2++
			}
		}
	}

	println(totalVisited)
	println(totalVisited2)
}
