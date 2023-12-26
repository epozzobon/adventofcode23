package day23

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay23Partfalseexample1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day23("example1.txt", false)), 94)
}
func TestDay23Parttrueexample1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day23("example1.txt", true)), 154)
}
func TestDay23Partfalseinput(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day23("input.txt", false)), 2406)
}
func TestDay23Parttrueinput(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day23("input.txt", true)), 6630)
}
