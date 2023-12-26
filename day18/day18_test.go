package day18

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay18Part1Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day18("example1.txt", 1)), 62) }
func TestDay18Part2Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day18("example1.txt", 2)), 952408144115)
}
func TestDay18Part1Input(t *testing.T) { utils.Assert(t, utils.Wrap(Day18("input.txt", 1)), 50746) }
func TestDay18Part2Input(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day18("input.txt", 2)), 70086216556038)
}
