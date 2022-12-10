package main

import (
	"AOC-2022/pkg/challenges"
	"fmt"
	"time"
)

var days = []func(){
	challenges.Day01,
	challenges.Day02,
	challenges.Day03,
	challenges.Day04,
	challenges.Day05,
	challenges.Day06,
	challenges.Day07,
	challenges.Day08,
	challenges.Day09,
}

func main() {
	totalElapsed := time.Duration(0)
	for day, challenge := range days {
		fmt.Printf("Day %02d\n", day+1)
		start := time.Now()
		challenge()
		elapsed := time.Since(start)
		fmt.Printf("Elapsed: %v\n", elapsed)
		totalElapsed += elapsed
	}
	fmt.Printf("Total elapsed time: %v\n", totalElapsed)
}
