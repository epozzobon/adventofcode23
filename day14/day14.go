package main

import (
	"fmt"

	"epozzobon.it/adventofcode23/utils"
)

func rotate(matrix utils.Matrix, out utils.Matrix) {
	H := len(matrix[0])
	V := len(matrix)
	if len(out) != H {
		panic("bad output size")
	}
	if len(out[0]) != V {
		panic("bad output size")
	}
	for c := 0; c < H; c++ {
		out[c] = make([]byte, V)
	}
	for r := 0; r < H; r++ {
		for c := 0; c < V; c++ {
			out[r][c] = matrix[V-1-c][r]
		}
	}
}

func pushRocksNorth(matrix utils.Matrix) {
	for r := 1; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == 'O' {
				i := r - 1
				for i >= 0 && matrix[i][c] == '.' {
					i--
				}
				i++
				matrix[r][c] = '.'
				matrix[i][c] = 'O'
			}
		}
	}
}

func cycle(in utils.Matrix) utils.Matrix {
	H := len(in[0])
	V := len(in)

	tmp := make(utils.Matrix, H)
	for c := 0; c < H; c++ {
		tmp[c] = make([]byte, V)
	}

	out := make(utils.Matrix, V)
	for r := 0; r < V; r++ {
		out[r] = make([]byte, H)
		copy(out[r], in[r])
	}

	pushRocksNorth(out)
	rotate(out, tmp)
	pushRocksNorth(tmp)
	rotate(tmp, out)
	pushRocksNorth(out)
	rotate(out, tmp)
	pushRocksNorth(tmp)
	rotate(tmp, out)

	return out
}

func southSupportWeight(matrix utils.Matrix) int {
	sum := 0
	for r := 0; r < len(matrix); r++ {
		w := len(matrix) - r
		// fmt.Printf("%s %d\n", matrix[r], w)
		rowWeight := 0
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == 'O' {
				rowWeight += w
			}
		}
		sum += rowWeight
	}
	return sum
}

func Day14(filepath string, problemPart int) int {
	matrix, err := utils.ReadMatrix(filepath)
	if err != nil {
		panic(err)
	}

	if problemPart == 1 {
		pushRocksNorth(matrix)
		return southSupportWeight(matrix)
	} else if problemPart != 2 {
		panic("Unknown problem part")
	}

	type MemoLine struct {
		matrix utils.Matrix
		weight int
		next   string
		index  int
	}
	memo := make(map[string]*MemoLine)
	l := 0
	var line *MemoLine
	for l = 0; l < 1000000000; l++ {
		key := matrix.ToStr()
		if l > 0 {
			line.next = key
		}
		line = memo[key]
		if line != nil {
			break
		}
		matrix = cycle(matrix)
		sum := southSupportWeight(matrix)
		line = &MemoLine{matrix, sum, "", l}
		memo[key] = line
		fmt.Println(l+1, line.weight)
	}
	fmt.Println(l+1, line.weight)
	fmt.Println(l+1, "=", line.index+1)
	period := l - line.index
	fmt.Println("period is", period)
	fmt.Println(l+1+period, line.weight)
	fmt.Println(l+1+period*2, line.weight)
	l += ((1000000000 - l) / period) * period
	for l+1 < 1000000000 {
		fmt.Println(l+1, line.weight)
		line = memo[line.next]
		l++
	}
	fmt.Println(l+1, line.weight)
	return line.weight
}

func main() {
	utils.CheckSolution(Day14, "example1.txt", 1, 136)
	utils.CheckSolution(Day14, "example1.txt", 2, 64)
	utils.CheckSolution(Day14, "input.txt", 1, 109638)
	utils.CheckSolution(Day14, "input.txt", 2, 102657)
}
