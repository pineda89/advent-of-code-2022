package day05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1RealData(t *testing.T) {
	d := &Day{}

	result := d.Part1("input.txt")

	assert.Equal(t, "NTWZZWHFV", result)
}

func TestPart1TestData(t *testing.T) {
	d := &Day{}

	result := d.Part1("input_testdata.txt")

	assert.Equal(t, "CMZ", result)
}

func TestPart2RealData(t *testing.T) {
	d := &Day{}

	result := d.Part2("input.txt")

	assert.Equal(t, "BRZGFVBTJ", result)
}

func TestPart2TestData(t *testing.T) {
	d := &Day{}

	result := d.Part2("input_testdata.txt")

	assert.Equal(t, "MCD", result)
}
