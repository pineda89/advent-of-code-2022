package day06

import (
	"os"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day06"
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
	return d.resolve(filepath, 4)
}

func (d *Day) Part2(filepath string) string {
	return d.resolve(filepath, 14)
}

func (d *Day) resolve(filepath string, uniqueChars int) string {
	for _, line := range strings.Fields(d.GetInput(filepath)) {
		for i := uniqueChars; i < len(line); i++ {
			data := make(map[string]int)
			for j := 0; j < uniqueChars; j++ {
				data[string(line[i-j])]++
			}
			if len(data) == uniqueChars {
				return strconv.Itoa(i + 1)
			}
		}
	}
	return ""
}
