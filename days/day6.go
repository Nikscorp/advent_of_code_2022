package days

import (
	"bufio"
	"fmt"
	"os"
)

func Day6() int {
	reader := bufio.NewReader(os.Stdin)

	var line string
	fmt.Fscanf(reader, "%s\n", &line)

	k := 14
	window := make(map[byte]int)
	for i := 0; i < k; i++ {
		window[line[i]]++
	}

	for i := k; i < len(line); i++ {
		if len(window) == k {
			return i
		}
		window[line[i-k]]--
		if window[line[i-k]] == 0 {
			delete(window, line[i-k])
		}
		window[line[i]]++
	}

	return -1
}
