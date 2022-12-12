package challenges

import (
	"AOC-2022/pkg/utils"
	"container/heap"
	"math"
)

type cell struct {
	x        int
	y        int
	height   rune
	steps    int
	priority int
	index    int
}

type PriorityQueue []*cell

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*cell)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func findPath(input [][]*cell, openSet *PriorityQueue, endX, endY int) int {
	maxY, maxX := len(input), len(input[0])
	for {
		current := heap.Pop(openSet).(*cell)
		nextStep := current.steps + 1
		x, y := current.x, current.y
		if x == endX && y == endY {
			return current.steps
		}

		if x < maxX-1 && input[y][x+1].height <= input[y][x].height+1 && input[y][x+1].steps > nextStep {
			input[y][x+1].steps = nextStep
			input[y][x+1].priority = int(math.Abs(float64(y-endY))+math.Abs(float64(x+1-endX))) + nextStep
			heap.Push(openSet, input[y][x+1])
		}
		if y < maxY-1 && input[y+1][x].height <= input[y][x].height+1 && input[y+1][x].steps > nextStep {
			input[y+1][x].steps = nextStep
			input[y+1][x].priority = int(math.Abs(float64(y+1-endY))+math.Abs(float64(x-endX))) + nextStep
			heap.Push(openSet, input[y+1][x])
		}
		if x > 0 && input[y][x-1].height <= input[y][x].height+1 && input[y][x-1].steps > nextStep {
			input[y][x-1].steps = nextStep
			input[y][x-1].priority = int(math.Abs(float64(y-endY))+math.Abs(float64(x-1-endX))) + nextStep
			heap.Push(openSet, input[y][x-1])
		}
		if y > 0 && input[y-1][x].height <= input[y][x].height+1 && input[y-1][x].steps > nextStep {
			input[y-1][x].steps = nextStep
			input[y-1][x].priority = int(math.Abs(float64(y-1-endY))+math.Abs(float64(x-endX))) + nextStep
			heap.Push(openSet, input[y-1][x])
		}
	}
}

func Day12() {
	input := utils.ImportInputLines(12)

	const maxSteps = 10000000
	var startX, startY, endX, endY int
	matrix := make([][]*cell, len(input))
	for y := range input {
		matrix[y] = make([]*cell, len(input[y]))
		for x, value := range input[y] {
			height := value
			if value == 'S' {
				startX, startY = x, y
				height = 'a'
			}
			if value == 'E' {
				endX, endY = x, y
				height = 'z'
			}

			matrix[y][x] = &cell{
				x:      x,
				y:      y,
				height: height,
				steps:  maxSteps,
			}
		}
	}

	matrix[startY][startX].steps = 0
	openSet := make(PriorityQueue, 0)
	heap.Push(&openSet, matrix[startY][startX])

	finalSteps := findPath(matrix, &openSet, endX, endY)
	println(finalSteps)

	for y := range matrix {
		for yy := range matrix {
			for _, resetCell := range matrix[yy] {
				resetCell.steps = maxSteps
			}
		}
		matrix[y][0].steps = 0
		openSet = make(PriorityQueue, 0)
		heap.Push(&openSet, matrix[y][0])
		newFinalSteps := findPath(matrix, &openSet, endX, endY)
		if newFinalSteps < finalSteps {
			finalSteps = newFinalSteps
		}
	}
	println(finalSteps)

}
