package day11

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day struct {
}

func (d *Day) GetDay() string {
	return "day11"
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
	monkeys := d.parse(filepath)
	return d.simulate(monkeys, 20, monkeys.part1checkDamage)
}

func (m *Monkeys) part1checkDamage(w int) int {
	return w / 3
}

func (d *Day) Part2(filepath string) string {
	monkeys := d.parse(filepath)
	return d.simulate(monkeys, 10000, monkeys.part2checkDamage)
}

func (m *Monkeys) part2checkDamage(w int) int {
	return w % m.mcm
}

type Monkeys struct {
	monkeys []*Monkey
	mcm     int
}

type Monkey struct {
	id                int
	objects           []int
	operation         string
	operationModifier string // that can be a number or the text "old"
	divisible         int
	TrueThrowTo       int
	FalseThrowTo      int
}

func (d *Day) parse(filepath string) *Monkeys {
	monkeys := &Monkeys{
		mcm: 1,
	}
	monkeys.monkeys = make([]*Monkey, 0)
	monkey := &Monkey{}
	for _, line := range strings.Split(d.GetInput(filepath), "\n") {
		switch {
		case line == "":
			monkey.id = len(monkeys.monkeys)
			monkeys.monkeys = append(monkeys.monkeys, monkey)
			monkey = &Monkey{}
		case strings.Contains(line, "Starting items:"):
			for _, v := range strings.Split(strings.Split(line, ":")[1], ",") {
				newV, _ := strconv.Atoi(strings.TrimSpace(v))
				monkey.objects = append(monkey.objects, newV)
			}
		case strings.Contains(line, "Operation: "):
			fmt.Sscanf(line, "  Operation: new = old %s %s", &monkey.operation, &monkey.operationModifier)
		case strings.Contains(line, "Test: "):
			fmt.Sscanf(line, "  Test: divisible by %d", &monkey.divisible)
			monkeys.mcm *= monkey.divisible // calculate the MCM for the part 2
		case strings.Contains(line, "throw to monkey"):
			var boolValue bool
			var monkeyNumToThrow int
			fmt.Sscanf(line, "    If %t: throw to monkey %d", &boolValue, &monkeyNumToThrow)
			if boolValue {
				monkey.TrueThrowTo = monkeyNumToThrow
			} else {
				monkey.FalseThrowTo = monkeyNumToThrow
			}
		}
	}

	return monkeys
}

func (m *Monkey) doOperation(object int) int {
	var mod int
	if m.operationModifier == "old" {
		mod = object
	} else {
		mod, _ = strconv.Atoi(m.operationModifier)
	}

	switch m.operation {
	case "+":
		return object + mod
	case "*":
		return object * mod
	default:
		return 0
	}
}

func (m *Monkey) mustThrowTo(v int) int {
	if v%m.divisible == 0 {
		return m.TrueThrowTo
	}
	return m.FalseThrowTo
}

func (m *Monkey) add(value int) {
	m.objects = append(m.objects, value)
}

func (d *Day) simulate(monkeys *Monkeys, numOfRounds int, inspectDamages func(int) int) string {
	inspected := make([]int, len(monkeys.monkeys))
	for currentRound := 0; currentRound < numOfRounds; currentRound++ {
		for i, monkey := range monkeys.monkeys {
			for _, object := range monkey.objects {
				newValue := inspectDamages(monkey.doOperation(object))
				monkeys.monkeys[monkey.mustThrowTo(newValue)].add(newValue)
				inspected[i]++
			}
			monkey.objects = make([]int, 0)
		}
	}
	sort.Ints(inspected)
	return strconv.Itoa(inspected[len(inspected)-1] * inspected[len(inspected)-2])
}
