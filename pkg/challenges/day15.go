package challenges

import (
	"AOC-2022/pkg/utils"
	"regexp"
	"sort"
	"strconv"
)

type IntervalList [][2]int

func (x IntervalList) Len() int           { return len(x) }
func (x IntervalList) Less(i, j int) bool { return x[i][0] < x[j][0] }
func (x IntervalList) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func Day15() {
	input := utils.ImportInputLines(15)

	const checkRow = 2000000
	var intervals IntervalList
	var sensorsData [][3]int
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		sensorX, _ := strconv.Atoi(matches[1])
		sensorY, _ := strconv.Atoi(matches[2])
		beaconX, _ := strconv.Atoi(matches[3])
		beaconY, _ := strconv.Atoi(matches[4])

		distance := utils.Abs(sensorX-beaconX) + utils.Abs(sensorY-beaconY)
		sensorsData = append(sensorsData, [3]int{sensorX, sensorY, distance})
		if sensorY+distance < checkRow || sensorY-distance > checkRow || (sensorX == beaconX && beaconY == checkRow) {
			continue
		}
		possibleBeaconMinX := sensorX - (distance - utils.Abs(checkRow-sensorY))
		possibleBeaconMaxX := sensorX + (distance - utils.Abs(checkRow-sensorY))
		if beaconY == checkRow && beaconX == possibleBeaconMinX {
			possibleBeaconMinX++
		}
		if beaconY == checkRow && beaconX == possibleBeaconMaxX {
			possibleBeaconMaxX--
		}
		intervalModified := false
		for i := range intervals {
			currentInterval := &intervals[i]
			if !intervalModified {
				if currentInterval[0] > possibleBeaconMinX && currentInterval[0] <= possibleBeaconMaxX {
					currentInterval[0] = possibleBeaconMinX
					intervalModified = true
				}
				if currentInterval[1] > possibleBeaconMinX && currentInterval[1] <= possibleBeaconMaxX {
					currentInterval[1] = possibleBeaconMaxX
					intervalModified = true
				}
			} else {

			}
		}
		if !intervalModified {
			intervals = append(intervals, [2]int{possibleBeaconMinX, possibleBeaconMaxX})
		}
	}
	sort.Sort(intervals)

	newIntervals := IntervalList{intervals[0]}
	ni := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= newIntervals[ni][1] {
			if intervals[i][1] >= newIntervals[ni][1] {
				newIntervals[ni][1] = intervals[i][1]
			}
		} else {
			newIntervals = append(newIntervals, intervals[i])
			ni++
		}
	}
	total := 0
	for _, interval := range newIntervals {
		total += utils.Abs(interval[1]-interval[0]) + 1
	}
	println(total)

out:
	for _, sensorData := range sensorsData {
		x, y, distance := sensorData[0], sensorData[1], sensorData[2]+1
		for i := 0; i < distance; i++ {
			found := false
			coords := [4][2]int{
				{x + i, y + distance - i},
				{x - i, y + distance - i},
				{x + i, y - distance - i},
				{x - i, y - distance - i},
			}
			for _, coord := range coords {
				_x, _y := coord[0], coord[1]
				if _x < 0 || _x > 4000000 || _y < 0 || _y > 4000000 {
					continue
				}
				for _, sensor := range sensorsData {
					if utils.Abs(sensor[0]-coord[0])+utils.Abs(sensor[1]-coord[1]) <= sensor[2] {
						found = true
						break
					}
				}
				if !found {
					println(coord[0]*4000000 + coord[1])
					break out
				}
			}
		}
	}
}
