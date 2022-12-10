package challenges

import (
	"AOC-2022/pkg/utils"
	"strconv"
	"strings"
)

func addPixelCRT(pixelPosition, spritePosition int) string {
	if pixelPosition == 0 {
		return ".\n"
	}
	if pixelPosition >= spritePosition && pixelPosition <= spritePosition+2 {
		return "#"
	}
	return "."
}

func Day10() {
	input := utils.ImportInputLines(10)

	x := 1
	cycles := 1
	signalStrength := 0
	CRT := ""

	for _, instruction := range input {
		instructionParts := strings.Split(instruction, " ")
		op := instructionParts[0]
		if op == "noop" {
			CRT += addPixelCRT(cycles%40, x%40)
			cycles++
			if cycles == 20 || cycles == 60 || cycles == 100 || cycles == 140 || cycles == 180 || cycles == 220 {
				signalStrength += cycles * x
			}
		} else {
			amount, err := strconv.Atoi(instructionParts[1])
			if err != nil {
				panic(err)
			}
			CRT += addPixelCRT(cycles%40, x%40)
			cycles++
			if cycles == 20 || cycles == 60 || cycles == 100 || cycles == 140 || cycles == 180 || cycles == 220 {
				signalStrength += cycles * x
			}
			CRT += addPixelCRT(cycles%40, x%40)
			cycles++
			x += amount
			if cycles == 20 || cycles == 60 || cycles == 100 || cycles == 140 || cycles == 180 || cycles == 220 {
				signalStrength += cycles * x
			}
		}
	}

	println(signalStrength)
	println(CRT)
}
