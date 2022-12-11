package challenges

import (
	"AOC-2022/pkg/utils"
	"sort"
	"strconv"
	"strings"
)

func add(amount int) func(int) int {
	return func(item int) int {
		return item + amount
	}
}

func mul(amount int) func(int) int {
	return func(item int) int {
		return item * amount
	}
}

func power2(item int) int {
	return item * item
}

type monkey struct {
	items       []int
	items2      []int
	operation   func(int) int
	test        int
	trueMonkey  byte
	falseMonkey byte
	inspected   int
	inspected2  int
}

func Day11() {
	rawInput := utils.ImportInput(11)

	monkeyDataList := strings.Split(rawInput, "\n\n")
	monkeys := make([]*monkey, len(monkeyDataList))
	GCD := 1
	for i, monkeyData := range monkeyDataList {
		monkeyLines := strings.Split(monkeyData, "\n")

		rawStartingItems := strings.Split(strings.Split(monkeyLines[1], ": ")[1], ", ")
		startingItems := make([]int, len(rawStartingItems))
		for j, rawStartingItem := range rawStartingItems {
			startingItem, _ := strconv.Atoi(rawStartingItem)
			startingItems[j] = startingItem
		}

		var operatorFunc func(int) int
		rawOperation := strings.Split(strings.Split(monkeyLines[2], "old ")[1], " ")
		operand, err := strconv.Atoi(rawOperation[1])
		if rawOperation[0] == "+" {
			operatorFunc = add(operand)
		} else {
			if err != nil {
				operatorFunc = power2
			} else {
				operatorFunc = mul(operand)
			}
		}

		testNumber, _ := strconv.Atoi(strings.Split(monkeyLines[3], "by ")[1])
		GCD *= testNumber

		testTrue := monkeyLines[4][len(monkeyLines[4])-1] - '0'

		testFalse := monkeyLines[5][len(monkeyLines[5])-1] - '0'

		monkeys[i] = &monkey{
			items:       startingItems,
			items2:      startingItems,
			operation:   operatorFunc,
			test:        testNumber,
			trueMonkey:  testTrue,
			falseMonkey: testFalse,
		}
	}

	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.inspected += len(m.items)
			for _, item := range m.items {
				newItem := m.operation(item)
				dividedItem := newItem / 3
				if dividedItem%m.test == 0 {
					monkeys[m.trueMonkey].items = append(monkeys[m.trueMonkey].items, dividedItem)
				} else {
					monkeys[m.falseMonkey].items = append(monkeys[m.falseMonkey].items, dividedItem)
				}
			}
			m.items = make([]int, 0)

			m.inspected2 += len(m.items2)
			for _, item := range m.items2 {
				newItem := m.operation(item) % GCD
				if newItem%m.test == 0 {
					monkeys[m.trueMonkey].items2 = append(monkeys[m.trueMonkey].items2, newItem)
				} else {
					monkeys[m.falseMonkey].items2 = append(monkeys[m.falseMonkey].items2, newItem)
				}
			}
			m.items2 = make([]int, 0)
		}
	}

	for i := 20; i < 10000; i++ {
		for _, m := range monkeys {
			m.inspected2 += len(m.items2)
			for _, item := range m.items2 {
				newItem := m.operation(item) % GCD
				if newItem%m.test == 0 {
					monkeys[m.trueMonkey].items2 = append(monkeys[m.trueMonkey].items2, newItem)
				} else {
					monkeys[m.falseMonkey].items2 = append(monkeys[m.falseMonkey].items2, newItem)
				}
			}
			m.items2 = make([]int, 0)
		}
	}

	var currentInspections, currentInspections2 []int
	for _, m := range monkeys {
		currentInspections = append(currentInspections, m.inspected)
		currentInspections2 = append(currentInspections2, m.inspected2)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(currentInspections)))
	sort.Sort(sort.Reverse(sort.IntSlice(currentInspections2)))
	println(currentInspections[0] * currentInspections[1])
	println(currentInspections2[0] * currentInspections2[1])
}
