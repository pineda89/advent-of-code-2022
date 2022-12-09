package days

import (
	"github.com/pineda89/advent-of-code-2022/days/day01"
	"github.com/pineda89/advent-of-code-2022/days/day02"
	"github.com/pineda89/advent-of-code-2022/days/day03"
	"github.com/pineda89/advent-of-code-2022/days/day04"
	"github.com/pineda89/advent-of-code-2022/days/day05"
	"github.com/pineda89/advent-of-code-2022/days/day06"
	"github.com/pineda89/advent-of-code-2022/days/day07"
	"github.com/pineda89/advent-of-code-2022/days/day08"
	"github.com/pineda89/advent-of-code-2022/days/day09"
)

var DaysArray []Day

func init() {
	addDay(&day01.Day{})
	addDay(&day02.Day{})
	addDay(&day03.Day{})
	addDay(&day04.Day{})
	addDay(&day05.Day{})
	addDay(&day06.Day{})
	addDay(&day07.Day{})
	addDay(&day08.Day{})
	addDay(&day09.Day{})
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
