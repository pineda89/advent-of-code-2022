package day05

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day05"
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
	data := make(map[int][]string)
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		var numElements, from, to int
		if _, err := fmt.Sscanf(line, "move %d from %d to %d", &numElements, &from, &to); err == nil {
			for i := 0; i < numElements; i++ {
				data[to] = append([]string{data[from][0]}, data[to]...)
				data[from] = data[from][1:]
			}
		} else {
			d.parseLine(data, line)
		}
	}

	result := bytes.Buffer{}
	for i := 1; i <= len(data); i++ {
		result.WriteString(data[i][0])
	}

	return result.String()
}

func (d *Day) Part2(filepath string) string {
	data := make(map[int][]string)
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		var numElements, from, to int
		if _, err := fmt.Sscanf(line, "move %d from %d to %d", &numElements, &from, &to); err == nil {
			for i := 0; i < numElements; i++ {
				data[to] = append([]string{data[from][numElements-i-1]}, data[to]...)
			}
			data[from] = data[from][numElements:]
		} else {
			d.parseLine(data, line)
		}
	}

	result := bytes.Buffer{}
	for i := 1; i <= len(data); i++ {
		result.WriteString(data[i][0])
	}

	return result.String()
}

func (d *Day) parseLine(data map[int][]string, line string) {
	ctr := 1
	for i := 1; i < len(line); i = i + 4 {
		if line[i-1:i] == "[" && line[i+1:i+2] == "]" {
			data[ctr] = append(data[ctr], line[i:i+1])
		}
		ctr++
	}
}
