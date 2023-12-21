package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type matrix [][]byte
type exploration map[int]map[int]bool

func readMatrix(filename string) (matrix, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make(matrix, 0)
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

func (self matrix) findChar(char byte) (int, int) {
	for i := 0; i < len(self); i++ {
		for j := 0; j < len(self[i]); j++ {
			if self[i][j] == char {
				return i, j
			}
		}
	}
	return -1, -1
}

func (self matrix) countChar(char byte) int {
	count := 0
	for i := 0; i < len(self); i++ {
		for j := 0; j < len(self[i]); j++ {
			if self[i][j] == char {
				count++
			}
		}
	}
	return count
}

func (self matrix) at(r, c int) byte {
	rn := len(self)
	r = (r%rn + rn) % rn
	cn := len(self[r])
	c = (c%cn + cn) % cn
	return self[r][c]
}

func (self matrix) print() {
	for i := 0; i < len(self); i++ {
		for j := 0; j < len(self[i]); j++ {
			fmt.Printf("%c", self[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (self matrix) step(x exploration) (exploration, int) {
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
				c := self.at(i+oi, j+oj)
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

func main() {
	matrix, err := readMatrix("input.txt")
	if err != nil {
		panic(err)
	}

	r0, c0 := matrix.findChar('S')
	matrix[r0][c0] = '.'
	x := make(exploration)
	x[r0] = map[int]bool{c0: true}

	matrix.print()
	prevSum := 1
	prevDiff := 0
	prevDiffDiff := 0
	requiredSteps := 26501365
	offset := (requiredSteps - 1) % len(matrix)
	i := 0
	var sum int
	for i = 0; i < requiredSteps; i++ {
		x, sum = matrix.step(x)
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
}
