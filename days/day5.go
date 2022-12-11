package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day5() string {
	stacks := make([][]byte, 0)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\n" || !strings.Contains(line, "[") {
			break
		}
		line = line[:len(line)-1]

		split := make([]string, 0)
		for i := 0; i < len(line); i += 4 {
			fmt.Println(line[i:])
			if string(line[i:i+3]) == "   " {
				split = append(split, "#")
			} else {
				split = append(split, line[i:i+3])
			}
		}

		fmt.Println(split, len(split))

		if len(stacks) == 0 {
			for i := 0; i < len(split); i++ {
				stack := make([]byte, 0)
				stacks = append(stacks, stack)
			}
		}

		for i, v := range split {
			if v == "#" {
				continue
			}
			stacks[i] = append(stacks[i], v[1])
		}
	}

	for _, stack := range stacks {
		reverse(stack)
	}

	fmt.Println(stacks)
	fmt.Fscanln(reader)

	for {
		var from, to, cnt int
		_, err := fmt.Fscanf(reader, "move %d from %d to %d\n", &cnt, &from, &to)
		if err != nil {
			break
		}

		from -= 1
		to -= 1

		for i := 0; i < cnt; i++ {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}

		reverse(stacks[to][len(stacks[to])-cnt : len(stacks[to])])
		fmt.Println(stacks)

	}

	res := make([]byte, 0, len(stacks))
	for _, stack := range stacks {
		res = append(res, stack[len(stack)-1])
	}

	return string(res)
}

func reverse(a []byte) {
	l := 0
	r := len(a) - 1

	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}
