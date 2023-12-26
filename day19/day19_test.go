package day19

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay19Part1Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day19("example1.txt", 1)), 19114)
}
func TestDay19Part2Example1(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day19("example1.txt", 2)), 167409079868000)
}
func TestDay19Part1Input(t *testing.T) { utils.Assert(t, utils.Wrap(Day19("input.txt", 1)), 495298) }
func TestDay19Part2Input(t *testing.T) {
	utils.Assert(t, utils.Wrap(Day19("input.txt", 2)), 132186256794011)
}
