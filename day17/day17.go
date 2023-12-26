package day17

import (
	"fmt"

	"epozzobon.it/adventofcode23/utils"
)

const (
	cN = iota
	cW = iota
	cS = iota
	cE = iota
)

var offX = []int{0, -1, 0, 1}
var offY = []int{-1, 0, 1, 0}

type state struct {
	heat      int
	direction int
	row       int
	col       int
	straight  int
}

func Day17(filepath string, problemPart int) int {

	matrix, err := utils.ReadMatrix(filepath)
	if err != nil {
		panic(err)
	}

	matrix.Print()
	intMatrix := utils.Map2D[byte, int](matrix, func(a byte) int { return int(a - '0') })

	if problemPart == 1 {
		return solve(intMatrix, 0, 3)
	} else if problemPart == 2 {
		return solve(intMatrix, 4, 10)
	} else {
		panic("Unknown problem part")
	}
}

func solve(matrix [][]int, minStraights, maxStraights int) int {

	var states = []state{}

	V := len(matrix)
	H := len(matrix[0])
	bound := 9 * (H + V)

	var bestHeats = make([][][][]int, V)
	for i := range bestHeats {
		bestHeats[i] = make([][][]int, H)
		for j := range bestHeats[i] {
			bestHeats[i][j] = make([][]int, 4)
			for c := range bestHeats[i][j] {
				bestHeats[i][j][c] = make([]int, maxStraights)
				for s := range bestHeats[i][j][c] {
					bestHeats[i][j][c][s] = bound
				}
			}
		}
	}

	starting1 := state{col: 0, row: 0, heat: 0, direction: cE, straight: 0}
	starting2 := state{col: 0, row: 0, heat: 0, direction: cS, straight: 0}
	states = append(states, starting1, starting2)

	for len(states) > 0 {
		s := states[len(states)-1]
		states = states[:len(states)-1]
		for d := 0; d < 4; d++ {
			if d == (s.direction+2)%4 {
				// can't turn back
				continue
			}
			dx, dy := offX[d], offY[d]
			x, y := s.col+dx, s.row+dy
			if x >= H || x < 0 || y >= V || y < 0 {
				// stay inside the matrix
				continue
			}
			heat := s.heat + matrix[y][x]
			straight := 1
			if d == s.direction {
				straight = s.straight + 1
			} else if s.straight < minStraights {
				// Must go straight at least `minStraights` tiles
				continue
			}
			if straight > maxStraights {
				// Can't go straight more than `maxStraights` tiles
				continue
			}

			if heat >= bound {
				// We went above the high bound -> we are spinning around
				continue
			}

			bestHeat := bestHeats[y][x][d][straight-1]
			if heat >= bestHeat {
				// I was already in this state with a better heat
				continue
			} else {
				// Save the best heat found in this state
				bestHeats[y][x][d][straight-1] = heat
			}

			newstate := state{heat, d, y, x, straight}
			if x == H-1 && y == V-1 && straight >= minStraights {
				bound = heat
				fmt.Println("Found new bound:", heat, newstate)
			} else {
				states = append(states, newstate)
			}
		}
	}

	return bound
}
