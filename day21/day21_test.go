package day21

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay21Part6Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day21("example1.txt", 6)), 16) }
func TestDay21Part10Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day21("example1.txt", 10)), 50)
}
func TestDay21Part50Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day21("example1.txt", 50)), 1594)
}
func TestDay21Part100Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day21("example1.txt", 100)), 6536)
}
func TestDay21Part500Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day21("example1.txt", 500)), 167004)
}
func TestDay21Part1000Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day21("example1.txt", 1000)), 668697)
}
func TestDay21Part5000Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day21("example1.txt", 5000)), 16733044)
}
func TestDay21Part64Input(t *testing.T) { utils.Assert(t, utils.Wrap(Day21("input.txt", 64)), 3853) }
func TestDay21Part26501365Input(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day21("input.txt", 26501365)), 639051580070841)
}
