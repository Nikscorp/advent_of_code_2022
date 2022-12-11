package days

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func Day1() int {
	reader := bufio.NewReader(os.Stdin)

	allCals := make([]int, 0)
	curCals := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			allCals = append(allCals, curCals)
			break
		}
		if line == "\n" {
			allCals = append(allCals, curCals)
			curCals = 0
			continue
		}
		n, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			panic(err)
		}
		curCals += n
	}

	sort.Ints(allCals)
	res := 0
	for i := 0; i < 3; i++ {
		res += allCals[len(allCals)-i-1]
	}

	return res
}
