package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func readMatrix(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := [][]int{}
	scanner := bufio.NewScanner(file)
	columns := 0
	rows := 0
	for scanner.Scan() {
		rows++
		txt := scanner.Text()
		if len(lines) == 0 {
			columns = len(txt)
		} else if len(txt) != columns {
			return nil, errors.New("Different row lengths")
		}
		var b []int
		b = make([]int, columns)
		for h := 0; h < columns; h++ {
			b[h] = int(txt[h] - '0')
		}
		lines = append(lines, b)
	}

	err = scanner.Err()
	if err != nil {
		panic("input read error")
	}

	return lines, nil
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

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

var states = []state{}

func main() {

	matrix, err := readMatrix("input.txt")
	if err != nil {
		panic(err)
	}

	printMatrix(matrix)

	fmt.Println(solve(matrix, 0, 3))
	fmt.Println(solve(matrix, 4, 10))
}

func solve(matrix [][]int, minStraights, maxStraights int) int {

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
