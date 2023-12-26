package day6

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay6Part1Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day6("example1.txt", 1)), 288) }
func TestDay6Part2Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day6("example1.txt", 2)), 71503) }
func TestDay6Part1Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day6("input.txt", 1)), 227850) }
func TestDay6Part2Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day6("input.txt", 2)), 42948149) }
