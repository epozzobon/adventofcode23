package main

import (
	"epozzobon.it/adventofcode23/utils"
)

type star struct {
	id int
	r  int
	c  int
	rr int
	rc int
}

func in(needle byte, haystack []byte) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == byte(needle) {
			return true
		}
	}
	return false
}

func findStars(matrix utils.Matrix) []star {
	output := []star{}
	id := 0
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == '#' {
				id++
				output = append(output, star{id: id, r: r, c: c, rr: r, rc: c})
			}
		}
	}
	return output
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func transpose(matrix utils.Matrix) utils.Matrix {
	output := make(utils.Matrix, len(matrix[0]))
	for c := 0; c < len(matrix[0]); c++ {
		output[c] = make([]byte, len(matrix))
	}
	for c := 0; c < len(output); c++ {
		for r := 0; r < len(matrix); r++ {
			output[c][r] = matrix[r][c]
		}
	}
	return output
}

func Day11(filepath string, repetitions int) int {

	/* Step 1: input processing */

	matrix, err := utils.ReadMatrix(filepath)
	if err != nil {
		panic(err)
	}
	matrix.Print()

	stars := findStars(matrix)

	for r := 0; r < len(matrix); r++ {
		hasStar := false
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == '#' {
				hasStar = true
				break
			}
		}
		if !hasStar {
			for i := range stars {
				if stars[i].r == r {
					panic("bad")
				} else if stars[i].r > r {
					stars[i].rr += repetitions - 1
				}
			}
		}
	}

	for c := 0; c < len(matrix[0]); c++ {
		hasStar := false
		for r := 0; r < len(matrix); r++ {
			if matrix[r][c] == '#' {
				hasStar = true
				break
			}
		}
		if !hasStar {
			for i := range stars {
				if stars[i].c == c {
					panic("bad")
				} else if stars[i].c > c {
					stars[i].rc += repetitions - 1
				}
			}
		}
	}

	sum1 := 0
	for j := 0; j < len(stars); j++ {
		for i := j + 1; i < len(stars); i++ {
			s1 := stars[j]
			s2 := stars[i]
			vOff := abs(s1.rr - s2.rr)
			hOff := abs(s1.rc - s2.rc)
			sum1 += (vOff + hOff)
		}
	}
	return sum1
}

func main() {
	utils.CheckSolution(Day11, "example1.txt", 2, 374)
	utils.CheckSolution(Day11, "example1.txt", 10, 1030)
	utils.CheckSolution(Day11, "example1.txt", 100, 8410)
	utils.CheckSolution(Day11, "input.txt", 2, 10313550)
	utils.CheckSolution(Day11, "input.txt", 1000000, 611998089572)
}
