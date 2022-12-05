package challenges

import (
	"AOC-2022/pkg/utils"
	"fmt"
)

func Day05() {
	input := utils.ImportInputLines(5)

	pilesNumber := (len(input[0]) + 1) / 4
	piles := make([][]byte, pilesNumber)
	piles2 := make([][]byte, pilesNumber)

	lineNumber := 0
	for ; ; lineNumber++ {
		line := input[lineNumber]
		if line[0] == ' ' {
			break
		}
		for i := 0; i < pilesNumber; i++ {
			char := line[i*4+1]
			if char != ' ' {
				piles[i] = append(piles[i], char)
				piles2[i] = append(piles2[i], char)
			}
		}
	}

	lineNumber += 2

	for ; lineNumber < len(input); lineNumber++ {
		var amount, source, destiny int
		_, err := fmt.Sscanf(input[lineNumber], "move %d from %d to %d", &amount, &source, &destiny)
		if err != nil {
			panic(err)
		}
		for i := 0; i < amount; i++ {
			item := piles[source-1][0]
			piles[source-1] = piles[source-1][1:]
			piles[destiny-1] = append([]byte{item}, piles[destiny-1]...)
		}
		piles2Slice := make([]byte, amount)
		copy(piles2Slice, piles2[source-1][:amount])
		piles2[destiny-1] = append(piles2Slice, piles2[destiny-1]...)
		piles2[source-1] = piles2[source-1][amount:]
	}

	var finalStringBytes, finalStringBytes2 []byte
	for i := 0; i < len(piles); i++ {
		finalStringBytes = append(finalStringBytes, piles[i][0])
		finalStringBytes2 = append(finalStringBytes2, piles2[i][0])
	}
	println(string(finalStringBytes))
	println(string(finalStringBytes2))
}
