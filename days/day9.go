package days

import (
	"bufio"
	"fmt"
	"os"
)

type coord struct {
	x, y int
}

var dirs = map[string]coord{
	"L": {0, -1},
	"R": {0, 1},
	"U": {1, 0},
	"D": {-1, 0},
}

var diags = []coord{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

func Day9() int {
	reader := bufio.NewReader(os.Stdin)
	coordSet := make(map[coord]struct{})

	curCoordHead := coord{0, 0}
	curCoordTails := make([]coord, 9)
	coordSet[curCoordTails[len(curCoordTails)-1]] = struct{}{}

	for {
		var (
			dirS string
			k    int
		)
		_, err := fmt.Fscanf(reader, "%s %d\n", &dirS, &k)
		if err != nil {
			break
		}

		dir := dirs[dirS]

		for t := 0; t < k; t++ {
			curCoordHead = coord{curCoordHead.x + dir.x, curCoordHead.y + dir.y}
			prev := curCoordHead
			for ii := 0; ii < len(curCoordTails); ii++ {
				if !isAdjacent(prev, curCoordTails[ii]) {
					if curCoordTails[ii].x == prev.x && curCoordTails[ii].y-prev.y == 2 {
						curCoordTails[ii] = coord{curCoordTails[ii].x, curCoordTails[ii].y - 1}
					} else if curCoordTails[ii].x == prev.x && curCoordTails[ii].y-prev.y == -2 {
						curCoordTails[ii] = coord{curCoordTails[ii].x, curCoordTails[ii].y + 1}
					} else if curCoordTails[ii].y == prev.y && curCoordTails[ii].x-prev.x == 2 {
						curCoordTails[ii] = coord{curCoordTails[ii].x - 1, curCoordTails[ii].y}
					} else if curCoordTails[ii].y == prev.y && curCoordTails[ii].x-prev.x == -2 {
						curCoordTails[ii] = coord{curCoordTails[ii].x + 1, curCoordTails[ii].y}
					} else {
						found := false
						for _, diag := range diags {
							tmp := coord{curCoordTails[ii].x + diag.x, curCoordTails[ii].y + diag.y}
							if isAdjacentDirect(prev, tmp) {
								curCoordTails[ii] = tmp
								found = true
								break
							}
						}
						if !found {
							panic("some error")
						}
					}
				}
				prev = curCoordTails[ii]
			}
			coordSet[curCoordTails[len(curCoordTails)-1]] = struct{}{}
		}
	}

	return len(coordSet)
}

func isAdjacent(coord1, coord2 coord) bool {
	ok := isAdjacentDirect(coord1, coord2)
	if ok {
		return ok
	}

	for _, diag := range diags {
		if (coord1.x+diag.x == coord2.x) && (coord1.y+diag.y == coord2.y) {
			return true
		}
	}

	return false
}

func isAdjacentDirect(coord1, coord2 coord) bool {
	return sqr(coord1.x-coord2.x)+sqr(coord1.y-coord2.y) <= 2
}

func sqr(x int) int {
	return x * x
}
