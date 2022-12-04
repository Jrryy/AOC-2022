package challenges

import (
	"AOC-2022/pkg/utils"
	"strconv"
	"strings"
)

func Day04() {
	input := utils.ImportInputLines(4)

	fullyContained := 0
	justOverlap := 0
	for _, line := range input {
		intervals := strings.Split(line, ",")
		interval1, interval2 := strings.Split(intervals[0], "-"), strings.Split(intervals[1], "-")
		start1, _ := strconv.Atoi(interval1[0])
		end1, _ := strconv.Atoi(interval1[1])
		start2, _ := strconv.Atoi(interval2[0])
		end2, _ := strconv.Atoi(interval2[1])
		if !(start1 > end2 || start2 > end1) {
			justOverlap++
			if start1 >= start2 && end1 <= end2 || start1 <= start2 && end1 >= end2 {
				fullyContained++
			}
		}
	}

	println(fullyContained)
	println(justOverlap)
}
