package days

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type opType int

const (
	opTypeAdd opType = iota
	opTypeMul
	opTypeSqr
)

type opWithArg struct {
	opType opType
	arg    int
}

type monkey struct {
	id             int
	items          []int
	testDiv        int
	monkeyIDOnPass int
	monkeyIDOnFail int
	opWithArg      opWithArg
}

func (m *monkey) doOp(worryLevel int) int {
	switch m.opWithArg.opType {
	case opTypeAdd:
		return worryLevel + m.opWithArg.arg
	case opTypeMul:
		return worryLevel * m.opWithArg.arg
	case opTypeSqr:
		return worryLevel * worryLevel
	}

	panic("invalid op")
}

func (m *monkey) testDivOp(worryLevel int) bool {
	return worryLevel%m.testDiv == 0
}

func Day11() int {
	reader := bufio.NewReader(os.Stdin)
	monkeys := make([]*monkey, 0)

	for {
		// line is Monkey i: or EOF
		_, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		// line is Starting items: 79, 98
		startingItemsStr, _ := reader.ReadString('\n')
		startingItemsStr = startingItemsStr[:len(startingItemsStr)-1]
		startingItemsStr = startingItemsStr[len("  Starting items: "):]
		startingItemsSplitted := strings.Split(startingItemsStr, ", ")
		startingItems := make([]int, 0, len(startingItemsSplitted))
		for _, v := range startingItemsSplitted {
			item, _ := strconv.Atoi(v)
			startingItems = append(startingItems, item)
		}
		// line is Operation: new = old * 19
		operationStr, _ := reader.ReadString('\n')
		operationStr = operationStr[:len(operationStr)-1]
		operationStr = operationStr[len("  Operation: new = "):]
		operationSplitted := strings.Split(operationStr, " ")
		var (
			opType opType
			arg    int
		)

		if operationSplitted[2] == "old" {
			opType = opTypeSqr
		} else if operationSplitted[1] == "+" {
			opType = opTypeAdd
			arg, _ = strconv.Atoi(operationSplitted[2])
		} else {
			opType = opTypeMul
			arg, _ = strconv.Atoi(operationSplitted[2])
		}

		// line is Test: divisible by 23
		divStr, _ := reader.ReadString('\n')
		divStr = divStr[:len(divStr)-1]
		divStr = divStr[len("  Test: divisible by "):]
		testDiv, _ := strconv.Atoi(divStr)

		// line is If true: throw to monkey 2
		monkeyTrueStr, _ := reader.ReadString('\n')
		monkeyTrueStr = monkeyTrueStr[:len(monkeyTrueStr)-1]
		monkeyTrueStr = monkeyTrueStr[len("    If true: throw to monkey "):]
		monkeyIDOnPass, _ := strconv.Atoi(monkeyTrueStr)

		// line is If false: throw to monkey 3
		monkeyFalseStr, _ := reader.ReadString('\n')
		monkeyFalseStr = monkeyFalseStr[:len(monkeyFalseStr)-1]
		monkeyFalseStr = monkeyFalseStr[len("    If false: throw to monkey "):]
		monkeyIDOnFail, _ := strconv.Atoi(monkeyFalseStr)

		curMonkey := &monkey{
			id:             len(monkeys),
			items:          startingItems,
			testDiv:        testDiv,
			monkeyIDOnPass: monkeyIDOnPass,
			monkeyIDOnFail: monkeyIDOnFail,
			opWithArg: opWithArg{
				opType: opType,
				arg:    arg,
			},
		}

		monkeys = append(monkeys, curMonkey)
		reader.ReadString('\n')
	}

	monkeysActivity := make([]int, len(monkeys))
	mod := 1
	for _, monkey := range monkeys {
		mod *= monkey.testDiv
	}

	for i := 0; i < 10000; i++ { // i < 20
		for _, monkey := range monkeys {
			monkeysActivity[monkey.id] += len(monkey.items)
			for _, worryLevel := range monkey.items {
				oldWorryLevel := worryLevel
				worryLevel = monkey.doOp(worryLevel)
				if oldWorryLevel > worryLevel {
					panic(fmt.Sprintf("round %d: %d > %d", i, oldWorryLevel, worryLevel))
				}

				// worryLevel /= 3
				worryLevel %= mod

				var monkeyIdToThrow int
				if monkey.testDivOp(worryLevel) {
					monkeyIdToThrow = monkey.monkeyIDOnPass
				} else {
					monkeyIdToThrow = monkey.monkeyIDOnFail
				}
				monkeys[monkeyIdToThrow].items = append(monkeys[monkeyIdToThrow].items, worryLevel)
				monkey.items = monkey.items[:0]
			}
		}
	}

	fmt.Println(monkeysActivity)
	sort.Ints(monkeysActivity)

	return monkeysActivity[len(monkeysActivity)-1] * monkeysActivity[len(monkeysActivity)-2]
}
