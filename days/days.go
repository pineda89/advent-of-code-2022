package days

import (
	"github.com/pineda89/advent-of-code-2022/days/day01"
)

var DaysArray []Day

func init() {
	addDay(&day01.Day{})
}

type Day interface {
	GetDay() string
	GetInput(filepath string) string
	GetReadme(filepath string) string
	Part1(filepath string) string
	Part2(filepath string) string
}

func addDay(day Day) {
	DaysArray = append(DaysArray, day)
}
