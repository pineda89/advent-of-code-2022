package day03

import (
	"os"
	"strconv"
	"strings"
)

const priorities = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Day struct {
}

func (d *Day) GetDay() string {
	return "day03"
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
	total := 0
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		marked := make(map[string]bool)
		left := line[:len(line)/2]
		right := line[len(line)/2:]
		for i := range left {
			if strings.Contains(right, string(left[i])) {
				if _, ok := marked[string(left[i])]; !ok {
					marked[string(left[i])] = true
					total = total + strings.Index(priorities, string(left[i])) + 1
				}
			}
		}
	}

	return strconv.Itoa(total)
}

func (d *Day) Part2(filepath string) string {
	lines := strings.Split(d.GetInput(filepath), "\n")
	size := 3
	total := 0

	for i := 0; i < len(lines); i = i + size {
		marked := make(map[string]bool)
		for i2 := range lines[i] {
			if _, ok := marked[string(lines[i][i2])]; !ok {
				if d.allContains(lines, i, size, string(lines[i][i2])) {
					total = total + strings.Index(priorities, string(lines[i][i2])) + 1
					marked[string(lines[i][i2])] = true
				}
			}

		}
	}
	return strconv.Itoa(total)
}

func (d *Day) allContains(lines []string, startindex int, size int, s string) bool {
	for i := startindex + 1; i < startindex+size; i++ {
		if !strings.Contains(lines[i], s) {
			return false
		}
	}
	return true
}
