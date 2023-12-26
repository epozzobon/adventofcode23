package day9

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay9Example1(t *testing.T) { utils.Assert(t, utils.Wrap(Day9("example1.txt")), 114, 2) }
func TestDay9Input(t *testing.T)    { utils.Assert(t, utils.Wrap(Day9("input.txt")), 1743490457, 1053) }
