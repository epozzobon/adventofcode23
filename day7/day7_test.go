package day7

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay7Part1Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day7("example1.txt", 1)), 6440) }
func TestDay7Part2Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day7("example1.txt", 2)), 5905) }
func TestDay7Part1Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day7("input.txt", 1)), 250957639) }
func TestDay7Part2Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day7("input.txt", 2)), 251515496) }
