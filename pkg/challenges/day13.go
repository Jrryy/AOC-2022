package challenges

import (
	"AOC-2022/pkg/utils"
	"strconv"
	"strings"
)

func fillInnerList(data string) ([]any, int) {
	var innerPack []any
	i := 1
	numStr := ""
	for data[i] != ']' {
		if data[i] == '[' {
			innerList, skip := fillInnerList(data[i:])
			innerPack = append(innerPack, innerList)
			i += skip
		} else if data[i] == ',' {
			if numStr != "" {
				num, _ := strconv.Atoi(numStr)
				innerPack = append(innerPack, num)
				numStr = ""
			}
			i++
		} else {
			numStr += string(data[i])
			i++
		}
	}
	if numStr != "" {
		num, _ := strconv.Atoi(numStr)
		innerPack = append(innerPack, num)
	}
	return innerPack, i + 1
}

func compInnerList(list1, list2 []any) int {
	partialResult := 0
	i := 0
	for ; ; i++ {
		if len(list2) == i && len(list1) > i {
			return -1
		} else if len(list1) == i && len(list2) > i {
			return 1
		} else if len(list1) == i && len(list2) == i {
			return 0
		}
		switch list1[i].(type) {
		case int:
			num1 := list1[i].(int)
			switch list2[i].(type) {
			case int:
				num2 := list2[i].(int)
				if num1 > num2 {
					return -1
				} else if num1 < num2 {
					return 1
				}
			case []any:
				innerList2 := list2[i].([]any)
				partialResult = compInnerList([]any{num1}, innerList2)
			}
		case []any:
			innerList1 := list1[i].([]any)
			switch list2[i].(type) {
			case int:
				num2 := list2[i].(int)
				partialResult = compInnerList(innerList1, []any{num2})
			case []any:
				innerList2 := list2[i].([]any)
				partialResult = compInnerList(innerList1, innerList2)
			}
		}
		if partialResult != 0 {
			return partialResult
		}
	}
}

func Day13() {
	input := utils.ImportInput(13)
	pairs := strings.Split(input, "\n\n")

	total1 := 0
	extraPacket1 := 1
	extraPacket2 := 2

	for i, pair := range pairs {
		splitPair := strings.Split(pair, "\n")
		pair1, pair2 := splitPair[0], splitPair[1]

		pack1, _ := fillInnerList(pair1)
		pack2, _ := fillInnerList(pair2)

		if result := compInnerList(pack1, pack2); result > 0 {
			total1 += i + 1
		}

		if result := compInnerList(pack1, []any{[]any{2}}); result > 0 {
			extraPacket1++
		}
		if result := compInnerList(pack1, []any{[]any{6}}); result > 0 {
			extraPacket2++
		}
		if result := compInnerList(pack2, []any{[]any{2}}); result > 0 {
			extraPacket1++
		}
		if result := compInnerList(pack2, []any{[]any{6}}); result > 0 {
			extraPacket2++
		}
	}
	println(total1)
	println(extraPacket1 * extraPacket2)
}
