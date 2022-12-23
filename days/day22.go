package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day22() int {
	file, _ := os.Open("input/input22.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	var cmd string
	lines := make([]string, 0)
	maxLen := 0
	starts := make([]int, 0)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if line == "\n" {
			line, _ := reader.ReadString('\n')
			cmd = line[:len(line)-1]
		}
		line = line[:len(line)-1]
		lines = append(lines, line)
		maxLen = max(maxLen, len(line))
	}

	grid := make([][]byte, len(lines))
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]byte, maxLen)
		startFound := false
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != ' ' && !startFound {
				starts = append(starts, j)
				startFound = true
			}
			grid[i][j] = lines[i][j]
		}
	}

	cmdS := make([]string, 0)

	prev := ""
	for i := 0; i < len(cmd); i++ {
		if cmd[i] >= '0' && cmd[i] <= '9' {
			prev += string(cmd[i])
		} else {
			if prev != "" {
				cmdS = append(cmdS, prev)
				prev = ""
			}
			cmdS = append(cmdS, string(cmd[i]))
		}
	}
	if prev != "" {
		cmdS = append(cmdS, prev)
	}
	fmt.Println(cmdS)

	start := coord{starts[0], 0}

	var dirs = []coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var facing = map[coord]int{
		{1, 0}:  0,
		{0, 1}:  1,
		{-1, 0}: 2,
		{0, -1}: 3,
	}
	dirInd := 0
	cur := start
	printGrid(grid)
	for i := 0; i < len(grid[11]); i++ {
		fmt.Println(i, string(grid[11][i]))
	}
	fmt.Println(string(grid[11]))

	for _, cmd := range cmdS {
		if cmd == "R" {
			dirInd = (dirInd + 1) % len(dirs)
			continue
		}
		if cmd == "L" {
			dirInd = (dirInd - 1) % len(dirs)
			if dirInd == -1 {
				dirInd = len(dirs) - 1
			}
			continue
		}
		n, _ := strconv.Atoi(cmd)
		dir := dirs[dirInd]
		for k := 0; k < n; k++ {
			newCoord := coord{cur.x + dir.x, cur.y + dir.y}
			if newCoord.y < 0 || newCoord.x < 0 || newCoord.y >= len(grid) || newCoord.x >= len(grid[newCoord.y]) || grid[newCoord.y][newCoord.x] == ' ' || grid[newCoord.y][newCoord.x] == 0 {
				// tmpDirInd := (dirInd + 2) % len(dirs)
				// tmpDir := dirs[tmpDirInd]
				// for {
				// 	tmpCoord := coord{newCoord.x + tmpDir.x, newCoord.y + tmpDir.y}
				// 	if tmpCoord.y < 0 || tmpCoord.x < 0 || tmpCoord.y >= len(grid) || tmpCoord.x >= len(grid[tmpCoord.y]) || grid[tmpCoord.y][tmpCoord.x] == ' ' || grid[tmpCoord.y][tmpCoord.x] == 0 {
				// 		break
				// 	}
				// 	newCoord = tmpCoord
				// }
				promt := fmt.Sprintf("i am at (%d, %d). Going to (%d, %d). Where to go?\n", newCoord.x, newCoord.y, dir.x, dir.y)
				fmt.Println(promt)
				tmpDir := coord{}
				_, err := fmt.Scanf("%d %d %d %d\n", &newCoord.x, &newCoord.y, &tmpDir.x, &tmpDir.y)
				if err != nil {
					panic(err)
				}
				if grid[newCoord.y][newCoord.x] == '#' {
					break
				}
				if grid[newCoord.y][newCoord.x] == '.' {
					cur = newCoord
					dir = tmpDir
					dirInd = facing[dir]
					continue
				}
				// wrap around
			}
			if grid[newCoord.y][newCoord.x] == '#' {
				break
			}
			if grid[newCoord.y][newCoord.x] == '.' {
				cur = newCoord
				continue
			}
		}
	}

	fmt.Println(cur)
	fmt.Println(starts)

	return 1000*(cur.y+1) + 4*(cur.x+1) + facing[dirs[dirInd]]
}
