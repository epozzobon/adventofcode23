package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func readMatrix(filename string) []([]byte) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	columns := 0
	rows := 0
	for scanner.Scan() {
		rows++
		txt := scanner.Text()
		if len(lines) == 0 {
			columns = len(txt)
		} else {
			if len(txt) != columns {
				panic("Different line lengths")
			}
		}
		lines = append(lines, txt)
	}

	err = scanner.Err()
	if err != nil {
		panic("input read error")
	}

	matrix := make([]([]byte), len(lines))
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]byte, len(lines[0]))
		for j := range matrix[i] {
			matrix[i][j] = lines[i][j]
		}
	}

	return matrix
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

func findStars(matrix []([]byte)) []star {
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

func transpose(matrix []([]byte)) []([]byte) {
	output := make([]([]byte), len(matrix[0]))
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

func main() {

	/* Step 1: input processing */

	matrix := readMatrix("input.txt")

	printMatrix(matrix)

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
					stars[i].rr += 1000000 - 1
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
					stars[i].rc += 1000000 - 1
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
	fmt.Println(sum1)
}
