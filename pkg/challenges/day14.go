package challenges

import (
	"AOC-2022/pkg/utils"
	"strconv"
	"strings"
)

func printGrid(grid [][]bool) {
	for _, row := range grid {
		for _, col := range row {
			if col {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
	println()
}

func Day14() {
	input := utils.ImportInputLines(14)

	lowest := 0

	grid := make([][]bool, 200)
	for i := range grid {
		grid[i] = make([]bool, 350)
	}

	for _, line := range input {
		prevX, prevY := -1, -1
		for _, rawCoord := range strings.Split(line, " -> ") {
			coord := strings.Split(rawCoord, ",")
			x, _ := strconv.Atoi(coord[0])
			x -= 325
			y, _ := strconv.Atoi(coord[1])
			if y > lowest {
				lowest = y
			}
			if prevX == -1 {
				prevX, prevY = x, y
				continue
			}
			diffX, diffY := prevX-x, prevY-y

			for prevY != y || prevX != x {
				grid[prevY][prevX] = true
				if diffY > 0 {
					prevY--
				} else if diffY < 0 {
					prevY++
				} else if diffX > 0 {
					prevX--
				} else {
					prevX++
				}
			}
			grid[prevY][prevX] = true
		}
	}

	i := 0
out:
	for ; ; i++ {
		x := 175
		y := 0
		for {
			if !grid[y+1][x] {
				y++
			} else if !grid[y+1][x-1] {
				y++
				x--
			} else if !grid[y+1][x+1] {
				y++
				x++
			} else {
				grid[y][x] = true
				break
			}

			if y > lowest {
				break out
			}
		}
	}

	println(i)

	for x := range grid[lowest+2] {
		grid[lowest+2][x] = true
	}

	for ; ; i++ {
		x := 175
		y := 0
		if grid[y][x] {
			break
		}
		for {
			if !grid[y+1][x] {
				y++
			} else if !grid[y+1][x-1] {
				y++
				x--
			} else if !grid[y+1][x+1] {
				y++
				x++
			} else {
				grid[y][x] = true
				break
			}
		}
	}
	//printGrid(grid)
	println(i)
}
