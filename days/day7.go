package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dir struct {
	name    string
	sumSize int
}

func Day7() int {
	reader := bufio.NewReader(os.Stdin)
	lines := make([][]string, 0)

	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "" {
			break
		}
		line = line[:len(line)-1]
		lines = append(lines, strings.Split(line, " "))
	}

	stack := make([]dir, 0)
	state := "default"
	sizeMap := make(map[string]int)
	i := 0
	for i < len(lines) {
		line := lines[i]
		// fmt.Println(state)
		switch state {
		case "default":
			if !isCommand(line) {
				fmt.Println(line)
				panic("wrong state")
			}

			if line[1] == "ls" {
				state = "ls"
				i++
				continue
			}

			if line[1] == "cd" {
				// NOTE: without i++
				state = "cd"
				continue
			}
			panic("wrong command")

		case "cd":
			if line[2] == ".." {
				last := stack[len(stack)-1]
				path := stackToPath(stack)
				sizeMap[path] = last.sumSize

				stack = stack[:len(stack)-1]
				stack[len(stack)-1].sumSize += last.sumSize

				state = "default"
				i++
				continue
			}
			stack = append(stack, dir{
				name: line[2],
			})
			state = "default"
			i++

		case "ls":
			if isCommand(line) {
				state = "default"
				// NOTE: without i++
				continue
			}
			if line[0] == "dir" {
				i++
				continue
			}
			fsize, err := strconv.Atoi(line[0])
			if err != nil {
				panic(err)
			}
			stack[len(stack)-1].sumSize += fsize
			i++
		}
	}

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		path := stackToPath(stack)
		sizeMap[path] = last.sumSize

		stack = stack[:len(stack)-1]
		if len(stack) > 0 {
			stack[len(stack)-1].sumSize += last.sumSize
		}
	}

	fmt.Println(sizeMap)

	// NOTE: task 1
	// const atMost int = 100000
	// res := 0

	// for _, v := range sizeMap {
	// 	if v <= atMost {
	// 		res += v
	// 	}
	// }

	// return res

	const totalSpace int = 70000000
	const needSpace int = 30000000
	freeSpace := totalSpace - sizeMap["/"]
	if freeSpace > needSpace {
		return 0
	}
	needToDeleteSpace := needSpace - freeSpace
	minSizeToDel := sizeMap["/"]

	for _, v := range sizeMap {
		if v >= needToDeleteSpace {
			minSizeToDel = min(minSizeToDel, v)
		}
	}

	return minSizeToDel

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func stackToPath(stack []dir) string {
	names := make([]string, 0, len(stack)-1)

	for i := 1; i < len(stack); i++ {
		names = append(names, stack[i].name)
	}

	return "/" + strings.Join(names, "/")
}

func isCommand(line []string) bool {
	return line[0] == "$"
}
