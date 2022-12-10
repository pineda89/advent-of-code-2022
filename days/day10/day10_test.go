package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1RealData(t *testing.T) {
	d := &Day{}

	result := d.Part1("input.txt")

	assert.Equal(t, "13740", result)
}

func TestPart1TestData(t *testing.T) {
	d := &Day{}

	result := d.Part1("input_testdata.txt")

	assert.Equal(t, "13140", result)
}

func TestPart2RealData(t *testing.T) {
	d := &Day{}

	result := d.Part2("input.txt")

	assert.Contains(t, result, "####.#..#.###..###..####.####..##..#....\n...#.#..#.#..#.#..#.#....#....#..#.#....\n..#..#..#.#..#.#..#.###..###..#....#....\n.#...#..#.###..###..#....#....#....#....\n#....#..#.#....#.#..#....#....#..#.#....\n####..##..#....#..#.#....####..##..####.")
}

func TestPart2TestData(t *testing.T) {
	d := &Day{}

	result := d.Part2("input_testdata.txt")

	assert.Contains(t, result, "##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######.....")
}
