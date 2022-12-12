package day12

import (
	"math"
	"os"
	"strconv"
	"strings"
)

var coords = []XY{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

type Day struct {
	data               [][]rune
	currentPosition    XY
	bestSignalPosition XY
	fewest             []XY
}

func (d *Day) GetDay() string {
	return "day12"
}

type XY struct {
	x int
	y int
}

type XYdist struct {
	xy   XY
	dist int
}

func (d *Day) GetInput(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

func (d *Day) GetReadme(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}

func (d *Day) Part1(filepath string) string {
	d.parse(filepath)
	return strconv.Itoa(d.pathfind(d.currentPosition, d.bestSignalPosition))
}

func (d *Day) Part2(filepath string) string {
	d.parse(filepath)
	var min = math.MaxInt
	for _, from := range d.fewest {
		steps := d.pathfind(from, d.bestSignalPosition)
		min = minG(min, steps)
	}
	return strconv.Itoa(min)
}

func (d *Day) parse(filepath string) {
	lines := strings.Split(d.GetInput(filepath), "\n")
	d.data = make([][]rune, len(lines))
	for i, line := range lines {
		if len(line) > 0 {
			d.data[i] = make([]rune, len(line))
			for j, c := range line {
				d.data[i][j] = c
				switch c {
				case 'S':
					d.currentPosition.x, d.currentPosition.y = i, j
					d.data[i][j] = 'a'
					d.fewest = append(d.fewest, XY{i, j})
				case 'E':
					d.bestSignalPosition.x, d.bestSignalPosition.y = i, j
					d.data[i][j] = 'z'
				case 'a':
					d.fewest = append(d.fewest, XY{i, j})
				}
			}
		} else {
			d.data = d.data[:len(d.data)-1]
		}
	}
}

func (d *Day) pathfind(from XY, to XY) int {
	queue := []XYdist{{from, 0}}
	visited := make(map[XY]int)

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v.xy == to {
			return v.dist
		}
		if _, ok := visited[v.xy]; ok {
			continue
		}
		visited[v.xy]++
		for _, coord := range coords {
			if d.isNeighbour(v.xy.x+coord.x, v.xy.y+coord.y) &&
				d.data[v.xy.x+coord.x][v.xy.y+coord.y]-d.data[v.xy.x][v.xy.y] <= 1 {
				// is neighbour and the height change is 1 or less
				queue = append(queue, XYdist{XY{v.xy.x + coord.x, v.xy.y + coord.y}, v.dist + 1})
			}
		}
	}
	return math.MaxInt
}

func (d *Day) isNeighbour(i int, j int) bool {
	return 0 <= i &&
		i < len(d.data) &&
		0 <= j &&
		j < len(d.data[0])
}

func minG[T int](first, second T) T {
	return T(math.Min(float64(first), float64(second)))
}
