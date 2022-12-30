package challenges

import (
	"AOC-2022/pkg/utils"
	"strconv"
	"strings"
	"sync"
)

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
				if !utils.Contains(visited, nextValve) {
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
	if !utils.Contains(opened, currentValve) {
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
		if utils.Contains(opened, valve) {
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
		result := make([]string, maxSize)
		copy(result, currentSubset)
		return [][]string{result}
	}
	var subsets [][]string
	for i := index; i <= len(valves)-maxSize+len(currentSubset); i++ {
		for _, newSubset := range findNewSubset(valves, maxSize, append(currentSubset, valves[i]), i+1) {
			subsets = append(subsets, newSubset)
		}
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
	bestPressurePerSubset := make([]int, len(subsets))

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(startJ int) {
			defer wg.Done()
			for j := startJ; j < len(subsets); j += 10 {
				opened := []string{"AA"}
				for valve := range distancesMap["AA"] {
					if !utils.Contains(subsets[j], valve) {
						opened = append(opened, valve)
					}
				}
				bestPressurePerSubset[j] = findBestPressure(
					opened,
					"AA",
					distancesMap,
					pressures,
					0,
					0,
					0,
					26,
				)
			}
		}(i)
	}
	wg.Wait()

	maxPressure := 0
	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(startJ int) {
			defer wg.Done()
			for j := startJ; j < len(subsets)-1; j += 10 {
			nextSubset:
				for k := j; k < len(subsets); k++ {
					for _, item := range subsets[j] {
						if utils.Contains(subsets[k], item) {
							continue nextSubset
						}
					}
					currentPressure := bestPressurePerSubset[j] + bestPressurePerSubset[k]
					mutex.Lock()
					if currentPressure > maxPressure {
						maxPressure = currentPressure
					}
					mutex.Unlock()
				}
			}
		}(i)
	}
	wg.Wait()

	println(maxPressure)
}
