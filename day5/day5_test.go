package day5

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay5Part1Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day5("example1.txt", 1)), 35) }
func TestDay5Part2Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day5("example1.txt", 2)), 46) }
func TestDay5Part1Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day5("input.txt", 1)), 51752125) }
func TestDay5Part2Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day5("input.txt", 2)), 12634632) }
