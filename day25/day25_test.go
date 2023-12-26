package day25

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay25Example(t *testing.T) { utils.Assert(t, utils.Wrap(Day25("example1.txt", 1)), 54) }
func TestDay25Input(t *testing.T)   { utils.Assert(t, utils.Wrap(Day25("input.txt", 1)), 602151) }
