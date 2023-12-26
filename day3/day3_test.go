package day3

import (
	"testing"

	"epozzobon.it/adventofcode23/utils"
)

func TestDay3Example(t *testing.T) { utils.Assert(t, utils.Wrap(Day3("example1.txt")), 4361, 467835) }
func TestDay3Input(t *testing.T)   { utils.Assert(t, utils.Wrap(Day3("input.txt")), 528819, 80403602) }
