package main

import (
	"AOC-2022/pkg/challenges"
	"fmt"
	"time"
)

var days = []func(){
	challenges.Day01,
}

func main() {
	for day, challenge := range days {
		fmt.Printf("Day %02d\n", day+1)
		start := time.Now()
		challenge()
		elapsed := time.Since(start)
		fmt.Printf("Elapsed: %v\n", elapsed)
	}
}
