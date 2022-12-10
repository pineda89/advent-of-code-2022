package day10

import (
	"bytes"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const LINE_WIDTH = 40

type Day struct {
	currentLine bytes.Buffer
}

func (d *Day) GetDay() string {
	return "day10"
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
	lines := strings.Split(d.GetInput(filepath), "\n")
	cycles := make([]int, len(lines)*2)
	cycleCounter := 0

	for _, line := range lines {
		if fields := strings.Fields(line); len(fields) > 0 {
			switch fields[0] {
			case "addx":
				toAddVal, _ := strconv.Atoi(fields[1])
				cycles[cycleCounter+3] += toAddVal
				cycleCounter += 2
			case "noop":
				cycleCounter++
			}
		}
	}

	cycles = cycles[:cycleCounter]

	var sum, value = 0, 1
	for i, cycle := range cycles {
		value += cycle
		if i%40 == 20 {
			sum += i * value
		}
	}

	return strconv.Itoa(sum)
}

func (d *Day) Part2(filepath string) string {
	lines := strings.Split(d.GetInput(filepath), "\n")

	value := 1
	for _, line := range lines {
		if fields := strings.Fields(line); len(fields) > 0 {
			switch fields[0] {
			case "addx":
				d.runCycle(value)
				d.runCycle(value)
				toAddVal, _ := strconv.Atoi(fields[1])
				value = value + toAddVal
			case "noop":
				d.runCycle(value)
			}
		}
	}

	return regexp.MustCompile(`(.{1,40})`).ReplaceAllString(d.currentLine.String(), "$1\n")
}

func (d *Day) runCycle(value int) {
	// If the sprite is positioned such that one of its three pixels is the pixel currently being drawn, the screen produces a lit pixel (#);
	// otherwise, the screen leaves the pixel dark (.).
	if d.currentLine.Len()%LINE_WIDTH < value-1 || d.currentLine.Len()%LINE_WIDTH > value+1 {
		d.currentLine.WriteString(".")
	} else {
		d.currentLine.WriteString("#")
	}
}
