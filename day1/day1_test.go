package day1

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay1Part1Example(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day1("example1.txt", 1)), 142)
}

func TestDay1Part1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day1("input.txt", 1)), 54605)
}

func TestDay1Part2Example(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day1("example2.txt", 2)), 281)
}

func TestDay1Part2(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day1("input.txt", 2)), 55429)
}
