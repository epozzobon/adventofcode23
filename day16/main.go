package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func readMatrix(filename string) ([][]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := [][]byte{}
	scanner := bufio.NewScanner(file)
	columns := 0
	rows := 0
	for scanner.Scan() {
		rows++
		txt := scanner.Text()
		if len(lines) == 0 {
			columns = len(txt)
		}
		b := []byte(txt)
		if len(lines) == 0 {
			columns = len(b)
		} else if len(b) != columns {
			return nil, errors.New("Different row lengths")
		}
		lines = append(lines, b)
	}

	err = scanner.Err()
	if err != nil {
		panic("input read error")
	}

	return lines, nil
}

func printMatrix(matrix []([]byte)) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%c", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

const (
	cN = 1 << iota
	cW = 1 << iota
	cS = 1 << iota
	cE = 1 << iota
)

type tile struct {
	char   byte
	energy int
}

func main() {

	matrix, err := readMatrix("input.txt")
	if err != nil {
		panic(err)
	}

	printMatrix(matrix)

	V := len(matrix)
	H := len(matrix[0])
	tiles := make([][]tile, V)
	for i := range tiles {
		tiles[i] = make([]tile, H)
		for j := range tiles[i] {
			tiles[i][j] = tile{char: matrix[i][j], energy: 0}
		}
	}

	var step func(int, int, int)
	step = func(pX, pY, direction int) {
		if pX < 0 || pX >= H || pY < 0 || pY >= V {
			// out of borders
			return
		}
		c := tiles[pY][pX].char
		e := tiles[pY][pX].energy
		if e&direction != 0 {
			// already visited the tile in this direction
			return
		}
		tiles[pY][pX].energy |= direction

		switch c {
		case '.':
			switch direction {
			case cE:
				step(pX+1, pY, cE)
			case cW:
				step(pX-1, pY, cW)
			case cN:
				step(pX, pY-1, cN)
			case cS:
				step(pX, pY+1, cS)
			}
		case '/':
			switch direction {
			case cE:
				step(pX, pY-1, cN)
			case cW:
				step(pX, pY+1, cS)
			case cN:
				step(pX+1, pY, cE)
			case cS:
				step(pX-1, pY, cW)
			}
		case '\\':
			switch direction {
			case cE:
				step(pX, pY+1, cS)
			case cW:
				step(pX, pY-1, cN)
			case cN:
				step(pX-1, pY, cW)
			case cS:
				step(pX+1, pY, cE)
			}
		case '|':
			switch direction {
			case cE:
				step(pX, pY-1, cN)
				step(pX, pY+1, cS)
			case cW:
				step(pX, pY-1, cN)
				step(pX, pY+1, cS)
			case cN:
				step(pX, pY-1, cN)
			case cS:
				step(pX, pY+1, cS)
			}
		case '-':
			switch direction {
			case cE:
				step(pX+1, pY, cE)
			case cW:
				step(pX-1, pY, cW)
			case cN:
				step(pX-1, pY, cW)
				step(pX+1, pY, cE)
			case cS:
				step(pX-1, pY, cW)
				step(pX+1, pY, cE)
			}
		default:
			panic("unknown tile")
		}
	}

	countEnergy := func(pX, pY, direction int) int {
		step(pX, pY, direction)
		sum := 0
		for i := range tiles {
			for j := range tiles[i] {
				if tiles[i][j].energy != 0 {
					sum++
				}
				tiles[i][j].energy = 0
			}
		}
		return sum
	}

	fmt.Println(countEnergy(0, 0, cE))

	max := 0
	printEnergyIfMax := func(pX, pY, direction int) int {
		sum := countEnergy(pX, pY, direction)
		if sum > max {
			max = sum
			fmt.Println("(", pX, pY, direction, ") =", sum)
		}
		return max
	}

	for x := 0; x < H; x++ {
		printEnergyIfMax(x, 0, cE)
		printEnergyIfMax(x, 0, cW)
		printEnergyIfMax(x, 0, cN)
		printEnergyIfMax(x, 0, cS)
		printEnergyIfMax(x, V-1, cE)
		printEnergyIfMax(x, V-1, cW)
		printEnergyIfMax(x, V-1, cN)
		printEnergyIfMax(x, V-1, cS)
	}
	for y := 0; y < V; y++ {
		printEnergyIfMax(0, y, cE)
		printEnergyIfMax(0, y, cW)
		printEnergyIfMax(0, y, cN)
		printEnergyIfMax(0, y, cS)
		printEnergyIfMax(H-1, y, cE)
		printEnergyIfMax(H-1, y, cW)
		printEnergyIfMax(H-1, y, cN)
		printEnergyIfMax(H-1, y, cS)
	}
}
