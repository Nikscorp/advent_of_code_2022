package days

import (
	"bufio"
	"fmt"
	"os"
)

var (
	intToChar = map[int]byte{-2: '=', -1: '-', 0: '0', 1: '1', 2: '2'}
	charToInt = map[byte]int{'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}
)

func Day25() string {
	file, _ := os.Open("input/input25.txt")
	reader := bufio.NewReader(file)
	res := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\n" {
			break
		}
		line = line[:len(line)-1]
		res += snafuToDesc(line)
	}

	return descToSnafu(res)
}

func snafuToDesc(s string) int {
	res := 0
	mul := 1
	for i := len(s) - 1; i >= 0; i-- {
		d := charToInt[s[i]]
		res += mul * d
		mul *= 5
	}
	return res
}

func descToSnafu(n int) string {
	fmt.Println(n)
	bkp := n
	res := []byte{}

	for n > 0 {
		d := n % 5
		if d >= 3 {
			n += d
			d -= 5
		}
		res = append(res, intToChar[d])
		n /= 5
	}

	reverseBytes(res)

	sRes := string(res)
	if bkp != snafuToDesc(sRes) {
		panic(sRes)
	}

	return sRes
}

func reverseBytes(b []byte) {
	l := 0
	r := len(b) - 1

	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}
