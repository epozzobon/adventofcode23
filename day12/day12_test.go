package day12

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay12Part1Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day12("example1.txt", 1)), 21) }
func TestDay12Part2Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day12("example1.txt", 2)), 525152)
}
func TestDay12Part1Input(t *testing.T) { utils.Assert(t, utils.Wrap(Day12("input.txt", 1)), 8075) }
func TestDay12Part2Input(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day12("input.txt", 2)), 4232520187524)
}
