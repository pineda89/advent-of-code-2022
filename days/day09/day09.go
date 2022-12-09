package day09

import (
	"math"
	"os"
	"strconv"
	"strings"
)

var directions = map[string]Point{"U": {x: 0, y: -1}, "R": {x: 1, y: 0}, "D": {x: 0, y: 1}, "L": {x: -1, y: 0}}

type Point struct {
	x int
	y int
}

type Day struct {
}

func (d *Day) GetDay() string {
	return "day09"
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
	return d.solve(filepath, 2)
}

func (d *Day) Part2(filepath string) string {
	return d.solve(filepath, 10)
}

func (d *Day) solve(filepath string, knots int) string {
	rope := make([]Point, knots)

	result := make(map[Point]int)
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		fields := strings.Fields(line)
		direction := fields[0]
		steps, _ := strconv.Atoi(fields[1])

		for step := 0; step < steps; step++ {
			rope[0] = Point{rope[0].x + directions[direction].x, rope[0].y + directions[direction].y}

			for knot := 1; knot < len(rope); knot++ {
				subpoint := Point{rope[knot-1].x - rope[knot].x, rope[knot-1].y - rope[knot].y}
				if d.abs(subpoint.x) > 1 || d.abs(subpoint.y) > 1 {
					rope[knot] = Point{rope[knot].x + d.sign(subpoint.x), rope[knot].y + d.sign(subpoint.y)}
				}
			}
			result[rope[len(rope)-1]]++ // mark it as visited
		}
	}

	return strconv.Itoa(len(result))
}

func (d *Day) abs(v int) int {
	return int(math.Abs(float64(v)))
}

func (d *Day) sign(v int) int {
	if v == 0 {
		return 0
	}
	return v / d.abs(v)
}
