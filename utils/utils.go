package utils

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Matrix [][]byte

func SpacedIntsList(txt string) ([]int, error) {
	vs := strings.Split(txt, " ")
	vsm := []int{}
	for _, v := range vs {
		if len(v) == 0 {
			continue
		}
		n, err := strconv.Atoi(string(v))
		if err != nil {
			return nil, err
		}
		vsm = append(vsm, n)
	}
	return vsm, nil
}

func MakeMatrix(rows, cols int) Matrix {
	out := make(Matrix, rows)
	for r := 0; r < rows; r++ {
		out[r] = make([]byte, cols)
	}
	return out
}

func (in Matrix) ToStr() string {
	V := len(in)
	key := ""
	for r := 0; r < V; r++ {
		key += string(in[r]) + "\n"
	}
	return key
}

func (self Matrix) FindChar(char byte) (int, int) {
	for i := 0; i < len(self); i++ {
		for j := 0; j < len(self[i]); j++ {
			if self[i][j] == char {
				return i, j
			}
		}
	}
	return -1, -1
}

func (self Matrix) CountChar(char byte) int {
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

func (self Matrix) At(r, c int, def byte) byte {
	if r < 0 || c < 0 || r >= len(self) || c >= len(self[r]) {
		return def
	}
	rn := len(self)
	r = (r%rn + rn) % rn
	cn := len(self[r])
	c = (c%cn + cn) % cn
	return self[r][c]
}

func (self Matrix) AtMod(r, c int) byte {
	rn := len(self)
	r = (r%rn + rn) % rn
	cn := len(self[r])
	c = (c%cn + cn) % cn
	return self[r][c]
}

func (self Matrix) Print() {
	for i := 0; i < len(self); i++ {
		for j := 0; j < len(self[i]); j++ {
			fmt.Printf("%c", self[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (self Matrix) GetColumn(col int) []byte {
	out := make([]byte, len(self))
	for i, s := range self {
		out[i] = s[col]
	}
	return out
}

func (self Matrix) Transpose() Matrix {
	out := make(Matrix, len(self[0]))
	for i := range self[0] {
		out[i] = self.GetColumn(i)
	}
	return out
}

func ReadLines(filename string) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		txt := scanner.Text()
		lines = append(lines, txt)
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func ReadMatrix(filename string) (Matrix, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make(Matrix, 0)
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

func Map[INPUT any, OUTPUT any](s []INPUT, f func(INPUT) OUTPUT) []OUTPUT {
	o := make([]OUTPUT, len(s))
	for i, e := range s {
		o[i] = f(e)
	}
	return o
}

func Map2D[INPUT any, OUTPUT any](matrix [][]INPUT, f func(INPUT) OUTPUT) [][]OUTPUT {
	return Map(matrix, func(row []INPUT) []OUTPUT {
		return Map(row, f)
	})
}

func CheckSolution[K comparable, V any](fn func(string, V) K, filename string, arg V, solution K) {
	sol := fn(filename, arg)
	log.Printf("Input \"%s\" (arg=%#v), solution %#v\n", filename, arg, sol)
	if solution != sol {
		panic("Wrong solution")
	}
}
