package day01

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day01"
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
	var highest = -1

	var current = 0
	for _, v := range strings.Split(d.GetInput(filepath), "\n") {
		if v == "" {
			if current > highest {
				highest = current
			}
			current = 0
		} else {
			newest, _ := strconv.Atoi(v)
			current = current + newest
		}
	}

	return strconv.Itoa(highest)
}

func (d *Day) Part2(filepath string) string {
	var highest = make([]int, 0)

	var current = 0
	for _, v := range strings.Split(d.GetInput(filepath), "\n") {
		if v == "" {
			highest = append(highest, current)
			current = 0
		} else {
			newest, _ := strconv.Atoi(v)
			current = current + newest
		}
	}

	sort.Ints(highest)

	return strconv.Itoa(highest[len(highest)-1] + highest[len(highest)-2] + highest[len(highest)-3])
}
