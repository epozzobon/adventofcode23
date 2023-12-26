package day13

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay13Part1Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day13("example1.txt", 1)), 405) }
func TestDay13Part2Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day13("example1.txt", 2)), 400) }
func TestDay13Part1Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day13("input.txt", 1)), 31265) }
func TestDay13Part2Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day13("input.txt", 2)), 39359) }
