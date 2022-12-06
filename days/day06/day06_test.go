package day06

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1RealData(t *testing.T) {
	d := &Day{}

	result := d.Part1("input.txt")

	assert.Equal(t, "1920", result)
}

func TestPart1TestData(t *testing.T) {
	d := &Day{}

	result := d.Part1("input_testdata.txt")

	assert.Equal(t, "7", result)
}

func TestPart2RealData(t *testing.T) {
	d := &Day{}

	result := d.Part2("input.txt")

	assert.Equal(t, "2334", result)
}

func TestPart2TestData(t *testing.T) {
	d := &Day{}

	result := d.Part2("input_testdata.txt")

	assert.Equal(t, "19", result)
}
