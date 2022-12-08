package challenges

import (
	"AOC-2022/pkg/utils"
)

func Day08() {
	input := utils.ImportInputMatrixDigits(8)

	visible := 0
	length := len(input)

	visibleMatrix := make([][]bool, length)
	for i := range visibleMatrix {
		visibleMatrix[i] = make([]bool, length)
	}

	maxScenic := 0

	for i := 0; i < length; i++ {
		up := -1
		down := -1
		right := -1
		left := -1
		for j := 0; j < length; j++ {
			// First part
			tree := input[i][j]
			if tree > left {
				if !visibleMatrix[i][j] {
					visible++
					visibleMatrix[i][j] = true
				}
				left = tree
			}

			tree = input[i][length-1-j]
			if tree > right {
				if !visibleMatrix[i][length-1-j] {
					visible++
					visibleMatrix[i][length-1-j] = true
				}
				right = tree
			}

			tree = input[j][i]
			if tree > up {
				if !visibleMatrix[j][i] {
					visible++
					visibleMatrix[j][i] = true
				}
				up = tree
			}

			tree = input[length-1-j][i]
			if tree > down {
				if !visibleMatrix[length-1-j][i] {
					visible++
					visibleMatrix[length-1-j][i] = true
				}
				down = tree
			}

			// Second part
			if i != 0 && i != length-1 && j != 0 && j != length-1 {
				var scenicUp, scenicLeft, scenicDown, scenicRight int
				current := input[i][j]
				for scenicLeft = j - 1; scenicLeft > 0 && input[i][scenicLeft] < current; scenicLeft-- {
				}
				for scenicUp = i - 1; scenicUp > 0 && input[scenicUp][j] < current; scenicUp-- {
				}
				for scenicRight = j + 1; scenicRight < length-1 && input[i][scenicRight] < current; scenicRight++ {
				}
				for scenicDown = i + 1; scenicDown < length-1 && input[scenicDown][j] < current; scenicDown++ {
				}
				scenic := (j - scenicLeft) * (i - scenicUp) * (scenicRight - j) * (scenicDown - i)
				if scenic > maxScenic {
					maxScenic = scenic
				}
			}
		}
	}

	println(visible)
	println(maxScenic)
}
