package days

import (
	"bufio"
	"fmt"
	"os"
)

var dirsMap = map[byte]coord{
	'>': {1, 0},
	'<': {-1, 0},
	'^': {0, -1},
	'v': {0, 1},
}

type coordWithPath struct {
	c    coord
	path string
}

func Day24() int {
	file, _ := os.Open("input/input24.txt")
	reader := bufio.NewReader(file)
	grid := make([][]byte, 0)

	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\n" {
			break
		}
		line = line[:len(line)-1]
		grid = append(grid, []byte(line))
	}

	blizzards := make(map[coord][]byte)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '#' || grid[i][j] == '.' {
				continue
			}
			blizzards[coord{j, i}] = append(blizzards[coord{j, i}], grid[i][j])
		}
	}

	start := coord{1, 0}
	end := coord{len(grid[0]) - 2, len(grid) - 1}
	var a, b, c int
	a, blizzards = blizzardBFS(grid, blizzards, start, end)
	b, blizzards = blizzardBFS(grid, blizzards, end, start)
	c, _ = blizzardBFS(grid, blizzards, start, end)

	return a + b + c
}

func blizzardBFS(grid [][]byte, blizzards map[coord][]byte, start, end coord) (int, map[coord][]byte) {
	queue := []coordWithPath{{start, ""}}
	steps := 0
L:
	for len(queue) > 0 {
		// seen is here as we can move backwards
		seen := make(map[coord]bool)
		blizzards = minute(grid, blizzards)
		steps++

		// fmt.Println(steps, len(queue))
		k := len(queue)

		for i := 0; i < k; i++ {
			cur := queue[0]
			queue = queue[1:]

			for dirC, dir := range dirsMap {
				newCoord := coordWithPath{coord{cur.c.x + dir.x, cur.c.y + dir.y}, cur.path + string(dirC)}
				if newCoord.c.x < 0 || newCoord.c.x >= len(grid[0]) {
					continue
				}
				if newCoord.c.y < 0 || newCoord.c.y >= len(grid) {
					continue
				}
				if grid[newCoord.c.y][newCoord.c.x] != '.' {
					continue
				}
				if newCoord.c == end {
					fmt.Println(steps, newCoord.path)
					break L
				}
				if seen[newCoord.c] {
					continue
				}
				seen[newCoord.c] = true
				queue = append(queue, newCoord)
			}

			if grid[cur.c.y][cur.c.x] == '.' {
				queue = append(queue, coordWithPath{cur.c, cur.path + "w"})
			}
		}
	}

	return steps, blizzards
}

func minute(grid [][]byte, blizzards map[coord][]byte) map[coord][]byte {
	newBlizzards := make(map[coord][]byte, len(blizzards))
	for theCoord, dirs := range blizzards {
		for _, dirByte := range dirs {
			dir := dirsMap[dirByte]
			newCoord := coord{theCoord.x + dir.x, theCoord.y + dir.y}
			if grid[newCoord.y][newCoord.x] == '#' {
				if dir.x == 1 {
					newCoord = coord{1, newCoord.y}
				} else if dir.x == -1 {
					newCoord = coord{len(grid[0]) - 2, newCoord.y}
				} else if dir.y == 1 {
					newCoord = coord{newCoord.x, 1}
				} else if dir.y == -1 {
					newCoord = coord{newCoord.x, len(grid) - 2}
				}
			}
			newBlizzards[newCoord] = append(newBlizzards[newCoord], dirByte)
			grid[theCoord.y][theCoord.x] = '.'
		}
	}

	for theCoord, dirs := range newBlizzards {
		if len(dirs) == 1 {
			grid[theCoord.y][theCoord.x] = dirs[0]
		} else {
			grid[theCoord.y][theCoord.x] = '0' + byte(min(len(dirs), 9))
		}
	}

	return newBlizzards
}
