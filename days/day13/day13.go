package day13

import (
	"encoding/json"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day struct {
	pairs []*Pair
}

type Pair struct {
	left             string
	right            string
	leftUnmarshaled  interface{}
	rightUnmarshaled interface{}
}

func (d *Day) GetDay() string {
	return "day13"
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
	d.parse(filepath)

	var total int
	for i := range d.pairs {
		if d.compare(d.pairs[i].leftUnmarshaled, d.pairs[i].rightUnmarshaled) < 1 {
			// 0 or lowest means are the correct order (left is lower, or both are equals)
			total = total + (i + 1)
		}
	}

	return strconv.Itoa(total)
}

func (d *Day) Part2(filepath string) string {
	d.parse(filepath)

	data := make([]interface{}, 0)
	for i := range d.pairs {
		data = append(data, d.pairs[i].leftUnmarshaled)
		data = append(data, d.pairs[i].rightUnmarshaled)
	}

	dividerPackets := make([]interface{}, 2)
	json.Unmarshal([]byte("[[2]]"), &dividerPackets[0])
	json.Unmarshal([]byte("[[6]]"), &dividerPackets[1])

	data = append(data, dividerPackets...)

	sort.Slice(data, func(i, j int) bool {
		return d.compare(data[i], data[j]) < 0
	})

	var result = 1
	for i := range data {
		if str, _ := json.Marshal(data[i]); string(str) == "[[2]]" || string(str) == "[[6]]" {
			result = result * (i + 1)
		}
	}

	return strconv.Itoa(result)
}

func (d *Day) parse(filepath string) {
	pairs := make([]*Pair, 0)
	pair := &Pair{}
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		if len(line) == 0 {
			pairs = append(pairs, pair)
			pair = &Pair{}
		} else {
			if len(pair.left) == 0 {
				pair.left = line
				json.Unmarshal([]byte(pair.left), &pair.leftUnmarshaled)
			} else {
				pair.right = line
				json.Unmarshal([]byte(pair.right), &pair.rightUnmarshaled)
			}
		}
	}
	d.pairs = pairs
}

func (d *Day) compare(left interface{}, right interface{}) int {
	var (
		leftValue, leftIsFloat   = left.(float64)
		rightValue, rightIsFloat = right.(float64)
	)

	if leftIsFloat && rightIsFloat {
		// end node. Both are values, comparables.
		return int(leftValue) - int(rightValue)
	}

	leftElements := d.toList(left)
	rightElements := d.toList(right)

	for i := range leftElements {
		if len(rightElements) <= i {
			// left has more elements than right. Order incorrect
			return 1
		}
		if result := d.compare(leftElements[i], rightElements[i]); result != 0 {
			// recursivity. If result is not equal, return it
			return result
		}
	}
	if len(rightElements) == len(leftElements) {
		// values are equals, and sizes of both lists are equals. The order is correct
		return 0
	}
	// right has more elements than left. The order is correct
	return -1
}

func (d *Day) toList(left interface{}) []interface{} {
	var leftElements []interface{}

	switch left.(type) {
	case []interface{}, []float64:
		// currently is a list, just assign
		leftElements = left.([]interface{})
	case float64:
		// is a unique element. Cast as a list to allow the comparision
		leftElements = []interface{}{left}
	}

	return leftElements
}
