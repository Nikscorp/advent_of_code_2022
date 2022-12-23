package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type direction struct {
	dirsToWatch   []coord
	dirsToProceed coord
}

func Day23() int {
	const (
		N    int = 300
		diff int = 70
	)
	grid := make([][]byte, N)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]byte, N)
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = '.'
		}
	}

	file, _ := os.Open("input/input23.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	i := diff
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\n" {
			break
		}
		line = line[:len(line)-1]
		for j := 0; j < len(line); j++ {
			grid[i][j+diff] = line[j]
		}
		i++
	}

	elves := make(map[coord]bool)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '#' {
				elves[coord{j, i}] = true
			}
		}
	}

	var dirs = []direction{
		{
			dirsToWatch:   []coord{{0, -1}, {-1, -1}, {1, -1}},
			dirsToProceed: coord{0, -1},
		},
		{
			dirsToWatch:   []coord{{0, 1}, {-1, 1}, {1, 1}},
			dirsToProceed: coord{0, 1},
		},
		{
			dirsToWatch:   []coord{{-1, 0}, {-1, 1}, {-1, -1}},
			dirsToProceed: coord{-1, 0},
		},
		{
			dirsToWatch:   []coord{{1, 0}, {1, 1}, {1, -1}},
			dirsToProceed: coord{1, 0},
		},
	}
	dirInd := 0

	for i := 0; i < 1000; i++ {
		activeElves := make(map[coord]bool)
		moved := false

		// 1st part of round
		for elf := range elves {
			isActive := false
		L:
			for ii := -1; ii <= 1; ii++ {
				for jj := -1; jj <= 1; jj++ {
					if ii == 0 && jj == 0 {
						continue
					}
					if grid[elf.y+ii][elf.x+jj] == '#' {
						isActive = true
						break L
					}
				}
			}
			if isActive {
				activeElves[elf] = true
			}
		}

		proposed := make(map[coord]int, len(activeElves))
		actions := make([]func(), 0, len(activeElves))
		for elf := range activeElves {
			curElf := elf
			curDirInd := dirInd

			for {
				dir := dirs[curDirInd]
				ok := true
				for _, dirToWatch := range dir.dirsToWatch {
					curCoord := coord{elf.x + dirToWatch.x, elf.y + dirToWatch.y}
					if grid[curCoord.y][curCoord.x] == '#' {
						ok = false
						break
					}
				}

				if ok {
					destCoord := coord{elf.x + dir.dirsToProceed.x, elf.y + dir.dirsToProceed.y}
					proposed[destCoord]++
					actions = append(actions, func() {
						if proposed[destCoord] != 1 {
							return
						}
						delete(elves, curElf)
						grid[curElf.y][curElf.x] = '.'
						elves[destCoord] = true
						grid[destCoord.y][destCoord.x] = '#'
						moved = true
					})
					break
				}

				curDirInd = (curDirInd + 1) % len(dirs)
				if curDirInd == dirInd {
					break
				}
			}
		}

		for _, action := range actions {
			action()
		}
		if !moved {
			fmt.Println("stopped:", i+1)
			break
		}

		dirInd = (dirInd + 1) % len(dirs)
	}

	minX := math.MaxInt
	minY := math.MaxInt
	maxX := 0
	maxY := 0
	for elf := range elves {
		minX = min(minX, elf.x)
		minY = min(minY, elf.y)
		maxX = max(maxX, elf.x)
		maxY = max(maxY, elf.y)
	}

	res := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if grid[y][x] == '.' {
				res++
			}
		}
	}

	return res
}
