package main

import (
	"AOC-2022/pkg/challenges"
	"flag"
	"fmt"
	"os"
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
	challenges.Day10,
	challenges.Day11,
	challenges.Day12,
	challenges.Day13,
	challenges.Day14,
	challenges.Day15,
	challenges.Day16,
}

func main() {
	dayFlag := flag.Int(
		"d",
		0,
		fmt.Sprintf("The day of the challenge you want to run, from 1 to %d", len(days)),
	)
	flag.Parse()
	if *dayFlag == 0 {
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
	} else if *dayFlag > len(days) {
		println("The challenge for this day has not yet been solved.")
		os.Exit(1)
	} else if *dayFlag < 0 {
		fmt.Printf("Please input a positive integer from 1 to %d\n", len(days))
		os.Exit(1)
	} else {
		start := time.Now()
		days[*dayFlag-1]()
		fmt.Printf("Elapsed: %v\n", time.Since(start))
	}
}
