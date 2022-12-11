package days

import (
	"bufio"
	"fmt"
	"os"
)

func Day3() int {
	reader := bufio.NewReader(os.Stdin)
	res := 0

	for {
		var s1, s2, s3 string

		_, err := fmt.Fscanf(reader, "%s\n%s\n%s\n", &s1, &s2, &s3)
		if err != nil {
			break
		}

		counter1 := make(map[byte]struct{})
		counter2 := make(map[byte]struct{})
		counter3 := make(map[byte]struct{})

		for i := 0; i < len(s1); i++ {
			counter1[s1[i]] = struct{}{}
		}

		for i := 0; i < len(s2); i++ {
			counter2[s2[i]] = struct{}{}
		}

		for i := 0; i < len(s3); i++ {
			counter3[s3[i]] = struct{}{}
		}

		for k := range counter1 {
			if _, ok := counter2[k]; ok {
				if _, ok := counter3[k]; ok {
					res += score(k)
					break
				}
			}
		}

	}

	return res
}

func score(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}

	return int(c - 'A' + 27)
}
