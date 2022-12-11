package days

import (
	"bufio"
	"fmt"
	"os"
)

func Day8() int {
	reader := bufio.NewReader(os.Stdin)
	grid := make([][]byte, 0)
	for {
		var line string
		_, err := fmt.Fscanf(reader, "%s\n", &line)
		if err != nil {
			break
		}
		grid = append(grid, []byte(line))
	}

	res := 0
	bestScore := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// NOTE: commented part is for task 1
			isVisible := true
			var a, b, c, d int
			for jj := j + 1; jj < len(grid[0]); jj++ {
				a++
				if grid[i][jj] >= grid[i][j] {
					isVisible = false
					break
				}
			}
			if isVisible {
				res++
				// continue
			}

			isVisible = true
			for jj := j - 1; jj >= 0; jj-- {
				b++
				if grid[i][jj] >= grid[i][j] {
					isVisible = false
					break
				}
			}
			if isVisible {
				res++
				// continue
			}

			isVisible = true
			for ii := i + 1; ii < len(grid); ii++ {
				c++
				if grid[ii][j] >= grid[i][j] {
					isVisible = false
					break
				}
			}
			if isVisible {
				res++
				// continue
			}

			isVisible = true
			for ii := i - 1; ii >= 0; ii-- {
				d++
				if grid[ii][j] >= grid[i][j] {
					isVisible = false
					break
				}
			}
			if isVisible {
				res++
				// continue
			}

			score := a * b * c * d
			bestScore = max(bestScore, score)
		}
	}

	// return res
	return bestScore
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
