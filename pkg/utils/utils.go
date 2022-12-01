package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ImportInput(day int) string {
	filename := fmt.Sprintf("input/%02d.txt", day)
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func ImportInputLines(day int) []string {
	rawInput := ImportInput(day)
	return strings.Split(rawInput, "\n")
}

func ImportInputLinesInt(day int) []int {
	rawInput := ImportInput(day)
	var input []int
	for _, rawLine := range strings.Split(rawInput, "\n") {
		line, err := strconv.Atoi(rawLine)
		if err != nil {
			panic(err)
		}
		input = append(input, line)
	}
	return input
}
