package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func Day14() int {
	reader := bufio.NewReader(os.Stdin)
	minX := math.MaxInt
	maxX := math.MinInt
	maxY := math.MinInt
	paths := make([][]coord, 0)
	for {
		var line string
		line, err := reader.ReadString('\n')
		if err != nil || line == "\n" {
			break
		}
		line = line[:len(line)-1]
		lineS := strings.Split(line, " -> ")
		path := make([]coord, 0, len(lineS))
		for _, coordRaw := range lineS {
			xRaw, yRaw, _ := strings.Cut(coordRaw, ",")
			x, _ := strconv.Atoi(xRaw)
			y, _ := strconv.Atoi(yRaw)
			path = append(path, coord{x, y})
			minX = min(minX, x)
			maxX = max(maxX, x)
			maxY = max(maxY, y)
		}
		paths = append(paths, path)
	}

	// hardcoded hack for task2 :(
	// assume that this range should be enough
	maxX = 1000
	minX = 0

	// normalize coords
	for _, path := range paths {
		for i := range path {
			path[i].x -= minX
		}
	}

	grid := make([][]byte, maxY+1+2) // +2 for task 2
	rowsC := maxX - minX + 1
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]byte, rowsC)
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = '.'
		}
	}
	start := coord{500 - minX, 0}
	grid[start.y][start.x] = '+'
	for j := 0; j < len(grid[0]); j++ {
		grid[len(grid)-1][j] = '#'
	}

	for _, path := range paths {
		from := path[0]
		for k := 1; k < len(path); k++ {
			to := path[k]
			// draw line
			if from.y == to.y {
				for j := min(from.x, to.x); j <= max(from.x, to.x); j++ {
					grid[from.y][j] = '#'
				}
			} else if from.x == to.x {
				for i := min(from.y, to.y); i <= max(from.y, to.y); i++ {
					grid[i][from.x] = '#'
				}
			}
			from = to
		}
	}

	print(grid)
	res := simulateSand(grid, start)
	print(grid)
	return res
}

func simulateSand(grid [][]byte, start coord) int {
	res := 0
L:
	for {
		cur := start
		for {
			if cur.y+1 >= len(grid) {
				break L
			}
			if grid[cur.y+1][cur.x] == '.' {
				cur = coord{cur.x, cur.y + 1}
				continue
			}
			if cur.x-1 < 0 {
				break L
			}
			if grid[cur.y+1][cur.x-1] == '.' {
				cur = coord{cur.x - 1, cur.y + 1}
				continue
			}
			if cur.x+1 >= len(grid[0]) {
				break L
			}
			if grid[cur.y+1][cur.x+1] == '.' {
				cur = coord{cur.x + 1, cur.y + 1}
				continue
			}
			break
		}
		grid[cur.y][cur.x] = 'o'
		res++
		if cur == start {
			break L
		}
	}

	return res
}

func print(grid [][]byte) {
	time.Sleep(10 * time.Millisecond)
	fmt.Println()
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
