package day16

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay16Part1Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day16("example1.txt", 1)), 46) }
func TestDay16Part2Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day16("example1.txt", 2)), 51) }
func TestDay16Part1Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day16("input.txt", 1)), 7472) }
func TestDay16Part2Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day16("input.txt", 2)), 7716) }
