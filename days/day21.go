package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type monkeyParsed struct {
	name       string
	left       string
	right      string
	sign       byte
	yellAmount int
}

type monkeyNode struct {
	name       string
	left       *monkeyNode
	right      *monkeyNode
	sign       byte
	yellAmount int
}

func Day21() int {
	reader := bufio.NewReader(os.Stdin)
	monkeys := make(map[string]*monkeyParsed)

	for {
		var (
			name, left, right string
			yellAmount        int
			sign              byte
		)
		//"root: pppw + sjmn"
		line, err := reader.ReadString('\n')
		if err != nil || line == "\n" {
			break
		}

		line = line[:len(line)-1]
		lineS := strings.Split(line, " ")
		name = lineS[0][:len(lineS[0])-1]
		if len(lineS) == 2 {
			yellAmount, _ = strconv.Atoi(lineS[1])
		} else {
			left = lineS[1]
			right = lineS[3]
			sign = lineS[2][0]
		}
		parsed := &monkeyParsed{
			name:       name,
			left:       left,
			right:      right,
			sign:       sign,
			yellAmount: yellAmount,
		}
		monkeys[name] = parsed
	}

	tree := createTree("root", monkeys)
	res, isAccurate := dfs(tree)

	fmt.Println(res, isAccurate)
	// return res

	// task 2
	human := findHuman(tree)
	fmt.Println(human)
	tree.sign = '-'
	l := math.MaxInt16 * math.MinInt32
	r := math.MaxInt16 * math.MaxInt32

	var pos int
	for l < r {
		pos = l/2 + r/2
		human.yellAmount = pos
		res, isAccurate = dfs(tree)
		fmt.Println(l, r, pos, res)
		if res == 0 {
			break
		}
		if res > 0 {
			l = pos + 1
		}
		if res < 0 {
			r = pos - 1
		}
	}

	fmt.Println(res, pos, isAccurate)
	if isAccurate {
		return pos
	}

	for {
		newPos := pos - 1
		human.yellAmount = newPos
		res, isAccurate = dfs(tree)
		fmt.Println(res, newPos, isAccurate)
		if res != 0 {
			break
		}
		pos = newPos
	}

	return pos
}

func createTree(rootName string, monkeys map[string]*monkeyParsed) *monkeyNode {
	parsed := monkeys[rootName]
	if parsed == nil {
		panic(rootName)
	}
	root := &monkeyNode{
		name:       parsed.name,
		sign:       parsed.sign,
		yellAmount: parsed.yellAmount,
	}

	if parsed.sign == 0 {
		return root
	}
	if parsed.left == "" || parsed.right == "" {
		panic("invalid input")
	}
	root.left = createTree(parsed.left, monkeys)
	root.right = createTree(parsed.right, monkeys)

	return root
}

func dfs(tree *monkeyNode) (int, bool) {
	if tree.left == nil && tree.right == nil {
		return tree.yellAmount, true
	}
	a, isAccurateA := dfs(tree.left)
	b, isAccurateB := dfs(tree.right)

	var res int
	var isAccurate bool = true
	switch tree.sign {
	case '+':
		res = a + b
	case '*':
		res = a * b
	case '/':
		res = a / b
		isAccurate = a%b == 0
	case '-':
		res = a - b
	}

	return res, isAccurate && isAccurateA && isAccurateB
}

func findHuman(root *monkeyNode) *monkeyNode {
	if root.name == "humn" {
		return root
	}
	var res *monkeyNode
	if root.left != nil {
		res = findHuman(root.left)
	}
	if res != nil {
		return res
	}

	if root.right != nil {
		res = findHuman(root.right)
	}
	return res
}
