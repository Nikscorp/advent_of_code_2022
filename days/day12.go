package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Day12() int {
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

	var (
		starts []coord
		end    coord
	)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'S' || grid[i][j] == 'a' { // or just grid[i][j] == 'S' for task 1
				starts = append(starts, coord{i, j})
			} else if grid[i][j] == 'E' {
				end = coord{i, j}
			}
		}
	}

	minSteps := math.MaxInt

	// to remove E from grid
	grid[end.x][end.y] = 'z'
	for _, start := range starts {
		// to remove possible 'S' from grid
		grid[start.x][start.y] = 'a'
		curSteps := bfs(grid, start, end)
		minSteps = min(minSteps, curSteps)
	}

	return minSteps
}

func bfs(grid [][]byte, start, end coord) int {
	seen := make(map[coord]bool)
	seen[start] = true
	queue := []coord{start}
	steps := 0
	var dirs = []coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	found := false

	// simple BFS
L:
	for len(queue) > 0 {
		k := len(queue)
		for i := 0; i < k; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur == end {
				found = true
				break L
			}

			for _, dir := range dirs {
				newCoord := coord{cur.x + dir.x, cur.y + dir.y}
				// isInvalid?
				if newCoord.x < 0 || newCoord.x >= len(grid) || newCoord.y < 0 || newCoord.y >= len(grid[0]) {
					continue
				}
				// isSeen?
				if seen[newCoord] {
					continue
				}
				// isTooHigh?
				isNewCoordGreater := grid[newCoord.x][newCoord.y] > grid[cur.x][cur.y]
				if isNewCoordGreater && grid[newCoord.x][newCoord.y]-grid[cur.x][cur.y] > 1 {
					continue
				}
				seen[newCoord] = true
				queue = append(queue, newCoord)
			}
		}
		steps++
	}

	if !found {
		fmt.Println(start, "to", end, "not fund")
		return math.MaxInt
	}

	return steps
}
