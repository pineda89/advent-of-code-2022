package day15

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type XY struct {
	x int
	y int
}

type Sensor struct {
	position XY
	beacon   XY
	strength int
}

type Day struct {
	data []*Sensor
}

func (d *Day) GetDay() string {
	return "day15"
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
	d.data = make([]*Sensor, 0)
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		s := &Sensor{}
		if _, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.position.x, &s.position.y, &s.beacon.x, &s.beacon.y); err == nil {
			s.strength = d.dist(s.position, s.beacon)
			d.data = append(d.data, s)
		}
	}
}

func (d *Day) dist(f, s XY) int {
	return int(math.Abs(float64(s.x-f.x)) + math.Abs(float64(s.y-f.y)))
}

func (d *Day) Part1(filepath string) string {
	d.parse(filepath)

	row := 2000000
	if strings.Contains(filepath, "test") {
		row = 10
	}

	minX, _, maxX, _ := d.getMinMax()

	total := 0
	for x := minX; x <= maxX; x++ {
		current := XY{x: x, y: row}

		canContainABeacon := true

		for _, s := range d.data {
			canContainABeacon = canContainABeacon && !(d.dist(s.position, current) <= s.strength && current != s.beacon)
		}

		if !canContainABeacon {
			total++
		}
	}

	return strconv.Itoa(total)
}

func (d *Day) Part2(filepath string) string {
	d.parse(filepath)

	high := 4000000
	multiplier := 4000000
	if strings.Contains(filepath, "test") {
		high = 20
	}

	for y := 0; y < high; y++ {
		for x := 0; x < high; x++ {
			anySensorInRange := false

			for _, s := range d.data {
				if distance := d.dist(s.position, XY{x, y}); distance <= s.strength {
					anySensorInRange = true
					x += s.strength - distance
					break
				}
			}

			if !anySensorInRange {
				// found!
				return strconv.Itoa((x * multiplier) + y)
			}
		}
	}

	return ""
}

func (d *Day) getMinMax() (int, int, int, int) {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt

	for _, s := range d.data {
		minX = d.min(s.position.x-s.strength, minX)
		minY = d.min(s.position.y-s.strength, minY)
		maxX = d.max(s.position.x+s.strength, maxX)
		maxY = d.max(s.position.y+s.strength, maxY)
	}

	return minX, minY, maxX, maxY
}

func (d *Day) min(x int, x2 int) int {
	return int(math.Min(float64(x), float64(x2)))
}

func (d *Day) max(x int, x2 int) int {
	return int(math.Max(float64(x), float64(x2)))
}
