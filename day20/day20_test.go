package day20

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay20Part1Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day20("example1.txt", 1)), 32000000)
}
func TestDay20Part1Example2(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day20("example2.txt", 1)), 11687500)
}
func TestDay20Part1Input(t *testing.T) { utils.Assert(t, utils.Wrap(Day20("input.txt", 1)), 763500168) }
func TestDay20Part2Input(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day20("input.txt", 2)), 207652583562007)
}
