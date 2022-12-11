package days

import (
	"bufio"
	"fmt"
	"os"
)

func Day4() int {
	reader := bufio.NewReader(os.Stdin)
	res := 0
	for {
		var l1, r1, l2, r2 int
		_, err := fmt.Fscanf(reader, "%d-%d,%d-%d\n", &l1, &r1, &l2, &r2)
		if err != nil {
			break
		}
		if (l1 >= l2 && l1 <= r2) || (l2 >= l1 && l2 <= r1) {
			res++
		}
	}
	return res
}
