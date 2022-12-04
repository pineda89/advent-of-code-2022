package day04

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day04"
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
	var total int

	for _, s := range strings.Fields(d.GetInput(filepath)) {
		var v [4]int
		fmt.Sscanf(s, "%d-%d,%d-%d", &v[0], &v[1], &v[2], &v[3])
		if (v[0] <= v[2] && v[1] >= v[3]) || (v[0] >= v[2] && v[1] <= v[3]) {
			// firststart is lower than secondstart and firstend is higher than secondend . Then second is fully contained by first
			// firststart is higher than secondstart and firstend is lower than secondend . Then first is fully contained by second
			total++
		}
	}

	return strconv.Itoa(total)
}

func (d *Day) Part2(filepath string) string {
	var total int

	for _, s := range strings.Fields(d.GetInput(filepath)) {
		var v [4]int
		fmt.Sscanf(s, "%d-%d,%d-%d", &v[0], &v[1], &v[2], &v[3])
		if v[0] <= v[3] && v[1] >= v[2] {
			total++
		}
	}

	return strconv.Itoa(total)
}
