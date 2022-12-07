package day07

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day07"
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

	result := 0
	for _, v := range data {
		if v <= 100000 {
			result += v
		}
	}

	return strconv.Itoa(result)
}

func (d *Day) Part2(filepath string) string {
	data := d.parse(filepath)

	result := data["/"]
	for _, v := range data {
		if v-data["/"]+70000000 >= 30000000 &&
			v < result {
			result = v
		}
	}

	return strconv.Itoa(result)
}

func (d *Day) parse(filepath string) map[string]int {
	files := make(map[string]int)
	var cd string

	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		var size int
		var file string

		if strings.HasPrefix(line, "$ cd") {
			cd = path.Join(cd, strings.Fields(line)[2])
		} else if _, err := fmt.Sscanf(line, "%d %s", &size, &file); err == nil {
			files[cd+"/"+file] = size
		}
	}

	data := make(map[string]int)
	for f, s := range files {
		subfolder := path.Dir(f)
		for {
			data[subfolder] += s

			if newSubfolder := path.Dir(subfolder); newSubfolder != subfolder {
				subfolder = newSubfolder
			} else {
				break
			}
		}
	}

	return data
}
