package day08

import (
	"os"
	"strconv"
	"strings"
)

var directions = []Direction{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type Direction struct {
	x int
	y int
}

type Day struct {
}

func (d *Day) GetDay() string {
	return "day08"
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
	data := d.parse(filepath)
	var total = (len(data) * 2) + (len(data[0]) * 2) - 4
	for rows := 1; rows < len(data)-1; rows++ {
		for cols := 1; cols < len(data[rows])-1; cols++ {
			if d.isVisible(data, rows, cols) {
				total++
			}
		}
	}

	return strconv.Itoa(total)
}

func (d *Day) Part2(filepath string) string {
	data := d.parse(filepath)
	bestScore := 0
	for rows := 0; rows < len(data); rows++ {
		for cols := 0; cols < len(data[rows]); cols++ {
			var score = 1
			for _, direction := range directions {
				visibleTrees := 0
				tmprows := rows + direction.x
				tmpcols := cols + direction.y

				for tmprows >= 0 && tmpcols >= 0 && tmprows < len(data) && tmpcols < len(data[0]) {
					visibleTrees++
					if data[tmprows][tmpcols] >= data[rows][cols] {
						break
					}

					tmprows += direction.x
					tmpcols += direction.y
				}

				score *= visibleTrees
			}
			if score > bestScore {
				bestScore = score
			}
		}
	}

	return strconv.Itoa(bestScore)
}

func (d *Day) parse(filepath string) [][]int {
	lines := strings.Split(d.GetInput(filepath), "\n")
	parsed := make([][]int, len(lines))
	for i, line := range lines {
		parsed[i] = make([]int, len(line))
		for j := range line {
			if val, err := strconv.Atoi(string(line[j])); err == nil {
				parsed[i][j] = val
			}
		}
	}
	return parsed
}

func (d *Day) isVisible(data [][]int, rows int, cols int) bool {
	initialValue := data[rows][cols]
	for _, direction := range directions {
		tmprows := rows + direction.x
		tmpcols := cols + direction.y
		directionIsVisible := true

		for tmprows >= 0 && tmpcols >= 0 && tmprows < len(data) && tmpcols < len(data[0]) {
			if data[tmprows][tmpcols] >= initialValue {
				directionIsVisible = false
			}

			tmprows = tmprows + direction.x
			tmpcols = tmpcols + direction.y
		}

		if directionIsVisible {
			return true
		}

	}
	return false
}
