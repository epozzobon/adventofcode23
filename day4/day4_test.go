package day4

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay4Example(t *testing.T) { utils.Assert(t, utils.Wrap(Day4("example1.txt")), 13, 30) }
func TestDay4Input(t *testing.T)   { utils.Assert(t, utils.Wrap(Day4("input.txt")), 19855, 10378710) }
