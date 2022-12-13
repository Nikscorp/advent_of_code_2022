package days

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type elemType int

const (
	typeSlice = iota
	typeInt
)

type elements []sliceOrInt

type sliceOrInt struct {
	elemType elemType
	slice    elements
	value    int
}

func Day13() int {
	reader := bufio.NewReader(os.Stdin)
	res := 0
	i := 1
	allElements := make([]elements, 0)
	for {
		var left, right string
		_, err := fmt.Fscanf(reader, "%s\n%s\n\n", &left, &right)
		if err != nil {
			break
		}
		leftEl := parseLine(left)
		rightEl := parseLine(right)
		allElements = append(allElements, leftEl)
		allElements = append(allElements, rightEl)
		if compareElements(leftEl, rightEl) {
			fmt.Println(i)
			res += i
		}
		i++
	}

	allElements = append(allElements, elements{
		{
			elemType: typeSlice,
			slice: elements{
				{
					elemType: typeInt,
					value:    2,
				},
			},
		},
	})
	allElements = append(allElements, elements{
		{
			elemType: typeSlice,
			slice: elements{
				{
					elemType: typeInt,
					value:    6,
				},
			},
		},
	})
	sort.Slice(allElements, func(i, j int) bool {
		return compareElements(allElements[i], allElements[j])
	})

	foundTwo := false
	twoInd := -1
	foundSix := false
	sixInd := -1
	for i, elem := range allElements {
		if len(elem) != 1 || elem[0].elemType != typeSlice || len(elem[0].slice) != 1 || elem[0].slice[0].elemType != typeInt {
			continue
		}

		if elem[0].slice[0].value == 2 {
			foundTwo = true
			twoInd = i
		} else if elem[0].slice[0].value == 6 {
			foundSix = true
			sixInd = i
		}

		if foundTwo && foundSix {
			break
		}
	}

	// return res
	return (twoInd + 1) * (sixInd + 1)
}

func parseLine(line string) elements {
	stack := make([]elements, 0)
	elems := make(elements, 0)
	for i := 0; i < len(line); i++ {
		if line[i] == ',' {
			continue
		}
		if line[i] == '[' {
			stack = append(stack, elems)
			elems = make(elements, 0)
			continue
		}
		if line[i] == ']' {
			parent := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			parent = append(parent, sliceOrInt{
				elemType: typeSlice,
				slice:    elems,
			})
			elems = parent
			continue
		}
		num := int(line[i] - '0')
		for j := i + 1; j < len(line) && line[j] >= '0' && line[j] <= '9'; j++ {
			num = num*10 + int(line[j]-'0')
			// ugly, but lazy
			i = j - 1
		}
		elems = append(elems, sliceOrInt{
			elemType: typeInt,
			value:    num,
		})
	}

	return elems
}

func compareElements(left, right elements) bool {
	for i := range left {
		if i >= len(right) {
			return false
		}
		le, re := left[i], right[i]
		ok, stop := compare(le, re)
		if !ok || stop {
			return ok
		}
	}

	return true
}

func compare(left, right sliceOrInt) (res, stop bool) {
	if left.elemType == typeInt && right.elemType == typeInt {
		if left.value < right.value {
			return true, true
		}
		if left.value > right.value {
			return false, true
		}
		return true, false
	}
	if left.elemType == typeInt {
		return compare(
			sliceOrInt{
				elemType: typeSlice,
				slice: []sliceOrInt{
					{
						elemType: typeInt,
						value:    left.value,
					},
				},
			},
			right)
	}
	if right.elemType == typeInt {
		return compare(
			left,
			sliceOrInt{
				elemType: typeSlice,
				slice: []sliceOrInt{
					{
						elemType: typeInt,
						value:    right.value,
					},
				},
			})
	}

	for i := range left.slice {
		if i >= len(right.slice) {
			return false, true
		}
		ok, stop := compare(left.slice[i], right.slice[i])
		if !ok || stop {
			return ok, stop
		}
	}

	if len(left.slice) < len(right.slice) {
		return true, true
	}

	return true, false
}
