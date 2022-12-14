package day14

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Day struct {
	data map[XY]cellType
}

type cellType int

const (
	cellType_rock cellType = 1
	cellType_sand cellType = -1
)

var (
	startPosition = XY{x: 500, y: 0}
)

func (d *Day) GetDay() string {
	return "day14"
}

type XY struct {
	x int
	y int
}

func (d *Day) GetInput(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

func (d *Day) GetReadme(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

func (d *Day) parse(filepath string) {
	d.data = make(map[XY]cellType)
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		var lastx, lasty int
		for _, value := range strings.Split(line, " -> ") {
			var xy XY
			if _, err := fmt.Sscanf(value, "%d,%d", &xy.x, &xy.y); err == nil {
				d.data[xy] = cellType_rock

				if xy.x == lastx || xy.y == lasty {
					deltax, deltay := d.getdelta(xy.x, lastx), d.getdelta(xy.y, lasty)
					x, y := lastx, lasty

					for x != xy.x || y != xy.y {
						d.data[XY{x: x, y: y}] = cellType_rock
						x += deltax
						y += deltay
					}
				}
				lastx, lasty = xy.x, xy.y
			}
		}
	}
}

func (d *Day) Part1(filepath string) string {
	d.parse(filepath)

	return strconv.Itoa(d.runpart1())
}

func (d *Day) Part2(filepath string) string {
	d.parse(filepath)

	return strconv.Itoa(d.runpart2(d.runpart1()))
}

func (d *Day) runpart1() int {
	_, _, _, maxy := d.getMinMax()

	var restSandBlocks = 0
	sandpos := startPosition
	for sandpos.y != maxy {
		sandpos = startPosition

		for sandpos.y != maxy {
			var rest int
			sandpos, rest = d.simulateSandDrop(sandpos)
			if rest != 0 {
				restSandBlocks += rest
				break
			}
		}
	}
	return restSandBlocks
}

func (d *Day) getMinMax() (int, int, int, int) {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt

	for xy := range d.data {
		minX = d.min(xy.x, minX)
		minY = d.min(xy.y, minY)
		maxX = d.max(xy.x, maxX)
		maxY = d.max(xy.y, maxY)
	}

	return minX, minY, maxX, maxY
}

func (d *Day) simulateSandDrop(sandpos XY) (XY, int) {
	toTest := []XY{
		{sandpos.x, sandpos.y + 1},     // down
		{sandpos.x - 1, sandpos.y + 1}, // left-down
		{sandpos.x + 1, sandpos.y + 1}, // right-down
	}

	for i := range toTest {
		if _, ok := d.data[toTest[i]]; !ok {
			return toTest[i], 0
		}
	}

	d.data[sandpos] = cellType_sand // no options available. Mark it
	return sandpos, 1
}

func (d *Day) runpart2(sandResting int) int {
	_, _, _, maxY := d.getMinMax()

	currentSand := XY{startPosition.x, startPosition.y}
	maxY += 2

	for {
		for _, v := range []int{-1, 0, 1} {
			// create floor
			d.data[XY{currentSand.x + v, maxY}] = cellType_rock
		}

		if nextSand, inc := d.simulateSandDrop(currentSand); inc == 0 {
			currentSand = XY{nextSand.x, nextSand.y}
		} else {
			d.data[XY{currentSand.x, currentSand.y}] = cellType_sand
			sandResting++
			if currentSand.x == startPosition.x && currentSand.y == startPosition.y {
				return sandResting
			}
			currentSand = XY{startPosition.x, startPosition.y}
		}
	}

	return 0
}

func (d *Day) min(x int, x2 int) int {
	return int(math.Min(float64(x), float64(x2)))
}

func (d *Day) max(x int, x2 int) int {
	return int(math.Max(float64(x), float64(x2)))
}

func (d *Day) getdelta(x int, lastx int) int {
	switch {
	case x == lastx:
		return 0
	case x > lastx:
		return 1
	case x < lastx:
		return -1
	}
	return 0
}
