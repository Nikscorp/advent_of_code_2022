package days

import (
	"bufio"
	"fmt"
	"os"
)

func Day17() int {
	reader := bufio.NewReader(os.Stdin)
	wind, _ := reader.ReadString('\n')
	wind = wind[:len(wind)-1]

	rock1 := [][]byte{
		{'#', '#', '#', '#'},
	}
	rock2 := [][]byte{
		{'.', '#', '.'},
		{'#', '#', '#'},
		{'.', '#', '.'},
	}
	rock3 := [][]byte{
		{'.', '.', '#'},
		{'.', '.', '#'},
		{'#', '#', '#'},
	}
	rock4 := [][]byte{
		{'#'},
		{'#'},
		{'#'},
		{'#'},
	}
	rock5 := [][]byte{
		{'#', '#'},
		{'#', '#'},
	}

	var rocks = [][][]byte{
		rock1,
		rock2,
		rock3,
		rock4,
		rock5,
	}

	grid := make([][]byte, 10000)
	for i := range grid {
		grid[i] = make([]byte, 7)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	fmt.Println(rocks)
	type mapKey struct {
		key     string
		rockInd int
		windInd int
	}

	type rockAndHeight struct {
		rocks  int
		height int
	}
	someMap := make(map[mapKey][]rockAndHeight)

	highest := len(grid)
	cnt := 0
	i := 0
	j := 0
	for {
		rock := rocks[i]
		cur := coord{2, highest - 3 - len(rock)}
		fmt.Println(cur.y, highest)

		for {
			curWind := wind[j]
			curWindDiff := -1
			if curWind == '>' {
				curWindDiff = 1
			}
			ok := true
		L1:
			for ii := 0; ii < len(rock); ii++ {
				newX := cur.x + curWindDiff
				for jj := 0; jj < len(rock[ii]); jj++ {
					if newX+jj >= len(grid[0]) || newX+jj < 0 || (rock[ii][jj] == '#' && grid[cur.y+ii][newX+jj] == '#') {
						ok = false
						break L1
					}
				}
			}

			if ok {
				cur = coord{cur.x + curWindDiff, cur.y}
			}

			ok = true
		L2:
			for ii := 0; ii < len(rock); ii++ {
				newY := cur.y + 1
				for jj := 0; jj < len(rock[ii]); jj++ {
					if newY+ii >= len(grid) || (rock[ii][jj] == '#' && grid[newY+ii][cur.x+jj] == '#') {
						ok = false
						break L2
					}
				}
			}

			if ok {
				cur = coord{cur.x, cur.y + 1}
			}

			if !ok {
				for ii := 0; ii < len(rock); ii++ {
					for jj := 0; jj < len(rock[ii]); jj++ {
						if grid[cur.y+ii][cur.x+jj] == '.' {
							grid[cur.y+ii][cur.x+jj] = rock[ii][jj]
						}
					}
				}
				highest = getHighest(grid)
				if highest < len(grid)-20 {
					key := getMapKey(grid, highest)
					mkey := mapKey{key, i, j}
					if _, ok := someMap[mkey]; ok {
						fmt.Println("pattern: ", cnt)
					}
					someMap[mkey] = append(someMap[mkey], rockAndHeight{cnt, highest})
				}
				j = (j + 1) % len(wind)
				break
			}
			j = (j + 1) % len(wind)
		}

		// printGrid(grid)
		i = (i + 1) % len(rocks)
		cnt++
		if cnt == 5000 {
			break
		}
	}

	printGrid(grid)
	res := len(grid) - getHighest(grid)

	rocksCount := 1000000000000 - 1
	for _, v := range someMap {
		if len(v) < 2 {
			continue
		}
		r1, h1 := v[len(v)-2].rocks, (len(grid) - v[len(v)-2].height)
		r2, h2 := v[len(v)-1].rocks, (len(grid) - v[len(v)-1].height)
		if (rocksCount-r1)%(r2-r1) == 0 {
			result := (rocksCount-r1)/(r2-r1)*(h2-h1) + h1
			fmt.Println(result)
		}
	}

	return res
}

func getMapKey(grid [][]byte, highest int) string {
	key := ""
	for i := highest; i <= highest+20; i++ {
		key += string(grid[i]) + "\n"
	}

	return key
}

func getHighest(grid [][]byte) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '#' {
				return i
			}
		}
	}
	return 0
}

func printGrid(grid [][]byte) {
	fmt.Println()
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
