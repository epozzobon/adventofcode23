package day24

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay24Part1Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day24p1("example1.txt", limits{7, 27})), 2)
}
func TestDay24Part1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day24p1("input.txt", limits{200000000000000, 400000000000000})), 24192)
}
func TestDay24Part2Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day24p2("example1.txt")), 47)
}
func TestDay24Part2Input(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day24p2("input.txt")), 664822352550558)
}
