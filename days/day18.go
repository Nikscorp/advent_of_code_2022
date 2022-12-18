package days

import (
	"bufio"
	"fmt"
	"os"
)

type cube struct {
	x, y, z   int
	openSides int
}

type coord3 struct {
	x, y, z int
}

type coords struct {
	c1, c2 coord3
}

func Day18_1() int {
	reader := bufio.NewReader(os.Stdin)
	cubes := make([]*cube, 0)
	for {
		var x, y, z int
		_, err := fmt.Fscanf(reader, "%d,%d,%d\n", &x, &y, &z)
		if err != nil {
			break
		}
		c := &cube{x, y, z, 6}
		cubes = append(cubes, c)
	}

	for i := 0; i < len(cubes); i++ {
		for j := i + 1; j < len(cubes); j++ {
			if abs(cubes[i].x-cubes[j].x)+abs(cubes[i].y-cubes[j].y)+abs(cubes[i].z-cubes[j].z) == 1 {
				cubes[i].openSides--
				cubes[j].openSides--
			}
		}
	}

	res := 0

	for _, cube := range cubes {
		res += cube.openSides
	}

	return res
}

func Day18_2() int {
	const gridSize int = 30
	grid := make([][][]byte, gridSize)
	for i := 0; i < len(grid); i++ {
		t := make([][]byte, gridSize)
		grid[i] = t
		for j := 0; j < len(grid[i]); j++ {
			t := make([]byte, gridSize)
			grid[i][j] = t
			for k := 0; k < len(grid[i][j]); k++ {
				grid[i][j][k] = '.'
			}
		}
	}
	reader := bufio.NewReader(os.Stdin)
	cubes := make([]*cube, 0)
	for {
		var x, y, z int
		_, err := fmt.Fscanf(reader, "%d,%d,%d\n", &x, &y, &z)
		if err != nil {
			break
		}
		// +2 to add free space before border
		c := &cube{x + 2, y + 2, z + 2, 6}
		cubes = append(cubes, c)
	}

	for _, cube := range cubes {
		grid[cube.x][cube.y][cube.z] = '#'
	}

	res := 0

	seen := make(map[coord3]bool)
	var dirs = []coord3{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}
	var queue = []coord3{{}}
	seen[coord3{}] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			newCoord := coord3{cur.x + dir.x, cur.y + dir.y, cur.z + dir.z}
			if newCoord.x < 0 || newCoord.y < 0 || newCoord.z < 0 {
				continue
			}
			if newCoord.x >= len(grid) || newCoord.y >= len(grid[0]) || newCoord.z >= len(grid[0][0]) {
				continue
			}
			if seen[newCoord] {
				continue
			}
			if grid[newCoord.x][newCoord.y][newCoord.z] == '#' {
				res++
				continue
			}

			seen[newCoord] = true
			queue = append(queue, newCoord)

		}
	}

	return res
}
