package day22

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay22Partfalseexample1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day22("example1.txt", false)), 5)
}
func TestDay22Parttrueexample1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day22("example1.txt", true)), 7)
}
func TestDay22Partfalseinput(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day22("input.txt", false)), 375)
}
func TestDay22Parttrueinput(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day22("input.txt", true)), 72352)
}
