package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type tile struct {
	r int
	c int
	d byte
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

func prettyMap(b byte) rune {
	if b == '-' {
		return '─'
	}
	if b == '|' {
		return '│'
	}
	if b == 'J' {
		return '┘'
	}
	if b == 'F' {
		return '┌'
	}
	if b == 'L' {
		return '└'
	}
	if b == '7' {
		return '┐'
	}
	return rune(b)
}

func printMatrix(matrix []([]byte)) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%c", prettyMap(matrix[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {

	/* Step 1: input processing */

	matrix := readMatrix("input.txt")

	printMatrix(matrix)

	/* Step 2: find loop */

	loop := []tile{}
	for r := range matrix {
		for c := range matrix[r] {
			if matrix[r][c] == 'S' {
				loop = append(loop, tile{r: r, c: c})
			}
		}
	}

	hasNorthConnection := func(r, c int) bool {
		return in(matrix[r][c], []byte{'L', 'J', '|', 'S'})
	}
	hasWestConnection := func(r, c int) bool {
		return in(matrix[r][c], []byte{'-', 'J', '7', 'S'})
	}
	hasEastConnection := func(r, c int) bool {
		return in(matrix[r][c], []byte{'-', 'L', 'F', 'S'})
	}
	hasSouthConnection := func(r, c int) bool {
		return in(matrix[r][c], []byte{'F', '|', '7', 'S'})
	}

	c1 := loop[len(loop)-1].c
	r1 := loop[len(loop)-1].r
	if hasNorthConnection(r1+1, c1) {
		loop = append(loop, tile{r: r1 + 1, c: c1})
	} else if hasWestConnection(r1, c1+1) {
		loop = append(loop, tile{r: r1, c: c1 + 1})
	} else if hasEastConnection(r1, c1-1) {
		loop = append(loop, tile{r: r1, c: c1 - 1})
	} else if hasSouthConnection(r1-1, c1) {
		loop = append(loop, tile{r: r1 - 1, c: c1})
	}

	d := 1
	for true {
		c0 := loop[len(loop)-2].c
		r0 := loop[len(loop)-2].r
		c1 := loop[len(loop)-1].c
		r1 := loop[len(loop)-1].r
		t := tile{}
		if r0 != r1+1 && hasSouthConnection(r1, c1) && hasNorthConnection(r1+1, c1) {
			t = tile{r: r1 + 1, c: c1}
		} else if c0 != c1+1 && hasEastConnection(r1, c1) && hasWestConnection(r1, c1+1) {
			t = tile{r: r1, c: c1 + 1}
		} else if c0 != c1-1 && hasWestConnection(r1, c1) && hasEastConnection(r1, c1-1) {
			t = tile{r: r1, c: c1 - 1}
		} else if r0 != r1-1 && hasNorthConnection(r1, c1) && hasSouthConnection(r1-1, c1) {
			t = tile{r: r1 - 1, c: c1}
		} else {
			panic("broken loop")
		}
		d++
		if t.r == loop[0].r && t.c == loop[0].c {
			break
		} else {
			loop = append(loop, t)
		}
	}

	fmt.Println(d / 2)

	/* Step 3: cleanup */

	cleaned := make([]([]byte), len(matrix))
	for i := 0; i < len(cleaned); i++ {
		cleaned[i] = make([]byte, len(matrix[0]))
		for j := range cleaned[i] {
			cleaned[i][j] = '.'
		}
	}

	for _, t := range loop {
		cleaned[t.r][t.c] = matrix[t.r][t.c]
	}

	for i := 0; i < len(cleaned); i++ {
		copy(matrix[i], cleaned[i])
	}

	printMatrix(matrix)

	/* Step 4: mark inner/outer tiles */

	fetch := func(r, c int) byte {
		if r < len(matrix) && c < len(matrix[0]) && r >= 0 && c >= 0 {
			return matrix[r][c]
		}
		return ' '
	}

	p := loop[0]
	for i, t := range loop {
		if i > 0 {
			xShift := p.c - t.c
			yShift := p.r - t.r
			tiles := [4]tile{
				{r: t.r + xShift, c: t.c - yShift, d: 'O'},
				{r: p.r + xShift, c: p.c - yShift, d: 'O'},
				{r: t.r - xShift, c: t.c + yShift, d: 'I'},
				{r: p.r - xShift, c: p.c + yShift, d: 'I'},
			}
			for _, t := range tiles {
				if fetch(t.r, t.c) == '.' {
					matrix[t.r][t.c] = t.d
				}
			}
		}
		p = t
	}

	printMatrix(matrix)

	var spread func(r, c int)
	spread = func(r, c int) {
		if fetch(r+1, c) == '.' {
			matrix[r+1][c] = matrix[r][c]
			spread(r+1, c)
		}
		if fetch(r, c+1) == '.' {
			matrix[r][c+1] = matrix[r][c]
			spread(r, c+1)
		}
		if fetch(r-1, c) == '.' {
			matrix[r-1][c] = matrix[r][c]
			spread(r-1, c)
		}
		if fetch(r, c-1) == '.' {
			matrix[r][c-1] = matrix[r][c]
			spread(r, c-1)
		}
	}

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == 'I' || matrix[r][c] == 'O' {
				spread(r, c)
			}
		}
	}

	sumi := 0
	sumo := 0
	sumd := 0
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == 'I' {
				sumi++
			}
			if matrix[r][c] == 'O' {
				sumo++
			}
			if matrix[r][c] == '.' {
				sumd++
			}
		}
	}

	printMatrix(matrix)
	if sumd > 0 {
		panic("Some tiles were not covered by inner or outer")
	}

	fmt.Println(sumi, sumo)
}
