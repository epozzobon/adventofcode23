package day11

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay11Part2Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day11("example1.txt", 2)), 374) }
func TestDay11Part10Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day11("example1.txt", 10)), 1030)
}
func TestDay11Part100Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day11("example1.txt", 100)), 8410)
}
func TestDay11Part2Input(t *testing.T) { utils.Assert(t, utils.Wrap(Day11("input.txt", 2)), 10313550) }
func TestDay11Part1000000Input(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day11("input.txt", 1000000)), 611998089572)
}
