package day10

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay10Part1Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day10("example1.txt", 1)), 4) }
func TestDay10Part1Example2(t *testing.T) { utils.Assert(t, utils.Wrap(Day10("example2.txt", 1)), 8) }
func TestDay10Part2Example3(t *testing.T) { utils.Assert(t, utils.Wrap(Day10("example3.txt", 2)), 8) }
func TestDay10Part2Example4(t *testing.T) { utils.Assert(t, utils.Wrap(Day10("example4.txt", 2)), 10) }
func TestDay10Part1Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day10("input.txt", 1)), 6951) }
func TestDay10Part2Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day10("input.txt", 2)), 563) }
