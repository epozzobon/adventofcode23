package day14

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay14Part1Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day14("example1.txt", 1)), 136) }
func TestDay14Part2Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day14("example1.txt", 2)), 64) }
func TestDay14Part1Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day14("input.txt", 1)), 109638) }
func TestDay14Part2Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day14("input.txt", 2)), 102657) }
