package challenges

import (
	"AOC-2022/pkg/utils"
	"strings"
)

func Day06() {
	input := utils.ImportInput(6)

	var set4, set14 []byte
	set4Done := false
	set14Done := false
	for i := 0; i < len(input) && !set14Done; i++ {
		if len(set4) < 4 {
			set4 = append(set4, input[i])
		} else if !set4Done {
			found := false
			for _, char := range set4 {
				if strings.Count(string(set4), string(char)) > 1 {
					set4 = append(set4[1:], input[i])
					found = true
					break
				}
			}
			if !found {
				println(i)
				set4Done = true
			}
		}

		if len(set14) < 14 {
			set14 = append(set14, input[i])
		} else {
			found := false
			for _, char := range set14 {
				if strings.Count(string(set14), string(char)) > 1 {
					set14 = append(set14[1:], input[i])
					found = true
					break
				}
			}
			if !found {
				println(i)
				set14Done = true
			}
		}
	}
}
