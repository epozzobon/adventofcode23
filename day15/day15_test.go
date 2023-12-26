package day15

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay15Part1Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day15("example1.txt", 1)), 1320)
}
func TestDay15Part2Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day15("example1.txt", 2)), 145) }
func TestDay15Part1Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day15("input.txt", 1)), 511343) }
func TestDay15Part2Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day15("input.txt", 2)), 294474) }
