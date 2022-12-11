package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day10() int {
	reader := bufio.NewReader(os.Stdin)
	state := make([]int, 1)
	x := 1
	state[0] = 1
	for {
		cmd, err := reader.ReadString('\n')
		if err != nil || cmd == "" || cmd == "\n" {
			break
		}
		cmd = cmd[:len(cmd)-1]

		if cmd == "noop" {
			state = append(state, x)
			continue
		}

		cmdS := strings.Split(cmd, " ")
		arg, _ := strconv.Atoi(cmdS[1])
		state = append(state, x)
		state = append(state, x)
		x += arg
	}

	res := 0
	for i := 20; i < len(state); i += 40 {
		res += state[i] * i
	}

	screen := make([][]rune, 6)
	for i := 0; i < len(screen); i++ {
		row := make([]rune, 40)
		screen[i] = row
	}

	for i := 1; i < len(state); i++ {
		curRow := (i - 1) / 40
		curColumn := (i - 1) % 40
		x = state[i]
		if x == curColumn || x-1 == curColumn || x+1 == curColumn {
			screen[curRow][curColumn] = '█'
		} else {
			screen[curRow][curColumn] = '░'
		}
	}

	for i := range screen {
		fmt.Println(string(screen[i]))
	}

	return res
}
