package day2

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay2Example(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day2("example1.txt")), 8, 2286)
}
func TestDay2Input(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day2("input.txt")), 2449, 63981)
}
