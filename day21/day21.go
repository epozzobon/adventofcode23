package main

import (
	"fmt"

	"epozzobon.it/adventofcode23/utils"
)

type exploration map[int]map[int]bool

func step(self utils.Matrix, x exploration) (exploration, int) {
	total := 0
	output := make(exploration)
	options := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i, v := range x {
		for j, v := range v {
			if !v {
				panic("how")
			}
			for p := 0; p < len(options); p++ {
				oi := options[p][0]
				oj := options[p][1]
				c := self.AtMod(i+oi, j+oj)
				if '#' != c {
					if output[i+oi] == nil {
						output[i+oi] = make(map[int]bool)
					}
					if !output[i+oi][j+oj] {
						output[i+oi][j+oj] = true
						total++
					}
				}
			}
		}
	}
	return output, total
}

func Day21(filename string, requiredSteps int) int {
	matrix, err := utils.ReadMatrix(filename)
	if err != nil {
		panic(err)
	}

	r0, c0 := matrix.FindChar('S')
	matrix[r0][c0] = '.'
	x := make(exploration)
	x[r0] = map[int]bool{c0: true}

	matrix.Print()
	prevSum := 1
	prevDiff := 0
	prevDiffDiff := 0
	offset := (requiredSteps - 1) % len(matrix)
	i := 0
	var sum int
	for i = 0; i < requiredSteps; i++ {
		x, sum = step(matrix, x)
		if i == 63 {
			fmt.Println(i+1, ":", sum)
		}
		if i%len(matrix) == offset {
			diff := sum - prevSum
			diffDiff := diff - prevDiff
			fmt.Println(i+1, ":", sum, ", derivative =", diff, ", second derivative = ", diffDiff)
			if prevDiffDiff == diffDiff {
				fmt.Println("Found constant second derivative", diffDiff)
				further := (requiredSteps - i) / len(matrix)
				avgDiff := diff + diffDiff*(further+1)/2 // second integral
				sum += avgDiff * further
				i += len(matrix)*further + 1
				break
			}
			prevDiffDiff = diffDiff
			prevDiff = diff
			prevSum = sum
		}
	}
	fmt.Println(i, ":", sum)
	return sum
}

func main() {
	utils.CheckSolution(Day21, "example1.txt", 6, 16)
	utils.CheckSolution(Day21, "example1.txt", 6, 16)
	utils.CheckSolution(Day21, "example1.txt", 10, 50)
	utils.CheckSolution(Day21, "example1.txt", 50, 1594)
	utils.CheckSolution(Day21, "example1.txt", 100, 6536)
	utils.CheckSolution(Day21, "example1.txt", 500, 167004)
	utils.CheckSolution(Day21, "example1.txt", 1000, 668697)
	utils.CheckSolution(Day21, "example1.txt", 5000, 16733044)
	utils.CheckSolution(Day21, "input.txt", 64, 3853)
	utils.CheckSolution(Day21, "input.txt", 26501365, 639051580070841)
}
