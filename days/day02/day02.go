package day02

import (
	"os"
	"strconv"
	"strings"
)

const (
	rock     = "rock"
	scissors = "scissors"
	paper    = "paper"
	lose     = "lose"
	draw     = "draw"
	win      = "win"
)

var (
	scores   = map[string]int{rock: 1, paper: 2, scissors: 3}
	priority = map[string]string{rock: scissors, scissors: paper, paper: rock}
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day02"
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
	conversions := map[string]string{"A": rock, "X": rock, "B": paper, "Y": paper, "C": scissors, "Z": scissors}

	var totalScore = 0
	for _, v := range strings.Split(d.GetInput(filepath), "\n") {
		if len(v) > 0 {
			fields := strings.Fields(v)
			rival := conversions[fields[0]]
			me := conversions[fields[1]]

			totalScore = totalScore + scores[me]
			if rival == me {
				totalScore = totalScore + 3
			} else {
				if priority[me] == rival {
					totalScore = totalScore + 6
				}
			}
		}
	}

	return strconv.Itoa(totalScore)
}

func (d *Day) Part2(filepath string) string {
	conversions := map[string]string{"A": rock, "B": paper, "C": scissors, "X": lose, "Y": draw, "Z": win}

	var totalScore = 0
	for _, v := range strings.Split(d.GetInput(filepath), "\n") {
		if len(v) > 0 {
			fields := strings.Fields(v)
			rival := conversions[fields[0]]
			expectedResult := conversions[fields[1]]
			var myhand string

			switch expectedResult {
			case draw:
				myhand = rival
				totalScore = totalScore + 3
			case win:
				for k, v := range priority {
					if v == rival {
						myhand = k
					}
				}
				totalScore = totalScore + 6
			case lose:
				myhand = priority[rival]
			}

			totalScore = totalScore + scores[myhand]
		}
	}

	return strconv.Itoa(totalScore)
}
