package day8

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay8Part1Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day8("example1.txt", 1)), 2) }
func TestDay8Part1Example2(t *testing.T) { utils.Assert(t, utils.Wrap(Day8("example2.txt", 1)), 6) }
func TestDay8Part2Example3(t *testing.T) { utils.Assert(t, utils.Wrap(Day8("example3.txt", 2)), 6) }
func TestDay8Part1Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day8("input.txt", 1)), 22411) }
func TestDay8Part2Input(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day8("input.txt", 2)), 11188774513823)
}
