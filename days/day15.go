package days

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type sensor struct {
	sensorCoord coord
	beaconCoord coord
	distance    int
}

func newSensor(sensorC, beaconC coord) *sensor {
	return &sensor{
		sensorCoord: sensorC,
		beaconCoord: beaconC,
		distance:    getDistance(sensorC, beaconC),
	}
}

func (s *sensor) isOverlap(c coord) bool {
	return getDistance(c, s.sensorCoord) <= s.distance
}

func getDistance(c1, c2 coord) int {
	return abs(c1.x-c2.x) + abs(c1.y-c2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Day15() int {
	reader := bufio.NewReader(os.Stdin)
	sensors := make([]*sensor, 0)
	beacons := make(map[coord]bool)
	minX := 0
	maxX := 0
	for {
		var x1, y1, x2, y2 int
		_, err := fmt.Fscanf(reader, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d\n", &x1, &y1, &x2, &y2)
		if err != nil {
			break
		}
		s := newSensor(coord{x1, y1}, coord{x2, y2})
		sensors = append(sensors, s)
		beacons[coord{x2, y2}] = true
		minX = min(minX, x1)
		minX = min(minX, x2)
		maxX = max(maxX, x1)
		maxX = max(maxX, x2)
	}

	res := 0

	// y := 2000000
	// y := 10
	sort.Slice(sensors, func(i, j int) bool {
		return sensors[i].distance > sensors[j].distance
	})
L:
	for x := 2_000_000; x <= 3_000_000; x++ {
		if x > 10 && x%10 == 0 {
			fmt.Println(x)
		}
		for y := 0; y <= 4_000_000; y++ {
			overlaps := false
			for _, s := range sensors {
				c := coord{x, y}
				if s.isOverlap(c) {
					overlaps = true
					break
				}
			}
			if !overlaps {
				fmt.Println(x, y)
				res = x*4000000 + y
				break L
			}
		}
	}

	return res
}
