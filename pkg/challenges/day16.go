package challenges

import (
	"AOC-2022/pkg/utils"
	"strconv"
	"strings"
	"sync"
)

func contains[T comparable](array []T, item T) bool {
	for _, _item := range array {
		if _item == item {
			return true
		}
	}
	return false
}

func bfs(toVisit [][]string, valves map[string][]string, pressures map[string]int) map[string]int {
	visited := []string{toVisit[0][0]}
	distances := make(map[string]int)
	for i := 0; i < len(toVisit); i++ {
		var nextToVisit []string
		for _, currentValve := range toVisit[i] {
			if i > 0 && pressures[currentValve] > 0 && distances[currentValve] == 0 {
				distances[currentValve] = i
			}
			nextValves := valves[currentValve]
			for _, nextValve := range nextValves {
				if !contains(visited, nextValve) {
					nextToVisit = append(nextToVisit, nextValve)
					visited = append(visited, nextValve)
				}
			}
		}
		if len(nextToVisit) > 0 {
			toVisit = append(toVisit, nextToVisit)
		}
	}
	return distances
}

func findBestPressure(
	opened []string,
	currentValve string,
	distances map[string]map[string]int,
	pressures map[string]int,
	minutes,
	accumulatedPressure,
	minutePressure,
	maxMinutes int,
) int {
	if minutes > maxMinutes {
		return accumulatedPressure - minutePressure*(minutes-maxMinutes)
	}
	var bestPressure int
	if !contains(opened, currentValve) {
		return findBestPressure(
			append(opened, currentValve),
			currentValve,
			distances,
			pressures,
			minutes+1,
			accumulatedPressure+minutePressure,
			minutePressure+pressures[currentValve],
			maxMinutes,
		)
	}
	nextValves := distances[currentValve]
	for valve, distance := range nextValves {
		if contains(opened, valve) {
			continue
		}
		nextBestPressure := findBestPressure(
			opened,
			valve,
			distances,
			pressures,
			minutes+distance,
			accumulatedPressure+minutePressure*distance,
			minutePressure,
			maxMinutes,
		)
		if nextBestPressure > bestPressure {
			bestPressure = nextBestPressure
		}
	}
	if bestPressure == 0 {
		return accumulatedPressure + minutePressure*(maxMinutes-minutes)
	} else {
		return bestPressure
	}
}

func findNewSubset(valves []string, maxSize int, currentSubset []string, index int) [][]string {
	if len(currentSubset) == maxSize {
		return [][]string{currentSubset}
	}
	var subsets [][]string
	for i := index; i <= len(valves)-maxSize+len(currentSubset); i++ {
		subsets = append(subsets, findNewSubset(valves, maxSize, append(currentSubset, valves[i]), i+1)...)
	}
	return subsets
}

func findAllSubsets(distances map[string]map[string]int) [][]string {
	var valves []string
	for valve := range distances["AA"] {
		valves = append(valves, valve)
	}
	var subsets [][]string
	for i := 1; i < 8; i++ {
		subsets = append(subsets, findNewSubset(valves, i, []string{}, 0)...)
	}
	return subsets
}

func findMaxPath(
	distances map[string]map[string]int,
	pressures map[string]int,
	current string,
	remaining []string,
	minutes int,
) int {
	maxPressure := 0
	if len(remaining) == 0 {
		return pressures[current] * (26 - minutes)
	}
	for _, valve := range remaining {
		if distances[current][valve] > 26-minutes {
			continue
		}
		var nextRemaining []string
		for _, item := range remaining {
			if item != valve {
				nextRemaining = append(nextRemaining, item)
			}
		}
		nextPressure := pressures[current]*(26-minutes) + findMaxPath(
			distances,
			pressures,
			valve,
			nextRemaining,
			minutes+distances[current][valve]+1,
		)
		if nextPressure > maxPressure {
			maxPressure = nextPressure
		}
	}
	return maxPressure
}

func Day16() {
	input := utils.ImportInputLines(16)

	pressures := make(map[string]int)
	valves := make(map[string][]string)

	for _, line := range input {
		splitInput := strings.Split(line, ";")
		splitHalf1 := strings.Split(splitInput[0], "=")
		name := splitHalf1[0][6:8]
		pressure, _ := strconv.Atoi(splitHalf1[1])
		_valves := strings.Split(splitInput[1], ", ")
		firstValve := _valves[0][len(_valves[0])-2 : len(_valves[0])]
		distances := make([]int, len(_valves))
		for i := range distances {
			distances[i] = 1
		}
		pressures[name] = pressure
		valves[name] = append([]string{firstValve}, _valves[1:]...)
	}

	distancesMap := make(map[string]map[string]int)

	for name := range valves {
		if name == "AA" || pressures[name] > 0 {
			distancesMap[name] = bfs([][]string{{name}}, valves, pressures)
		}
	}

	bestPressure := findBestPressure(
		[]string{"AA"},
		"AA",
		distancesMap,
		pressures,
		0,
		0,
		0,
		30,
	)

	println(bestPressure)
	subsets := findAllSubsets(distancesMap)

	var wg sync.WaitGroup
	maxPressure := 0
	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(startJ int) {
			defer wg.Done()
			for j := startJ - 1; j < len(subsets1); j += 10 {
				nextMaxPressure1 := findMaxPath(
					distancesMap,
					pressures,
					"AA",
					subsets1[j],
					0,
				)
				nextMaxPressure2 := findMaxPath(
					distancesMap,
					pressures,
					"AA",
					subsets2[j],
					0,
				)
				nextMaxPressure := nextMaxPressure1 + nextMaxPressure2
				// fmt.Println(nextMaxPressure)
				if nextMaxPressure > maxPressure {
					maxPressure = nextMaxPressure
				}
			}
		}(i)
	}
	wg.Wait()
	println(maxPressure)
}
