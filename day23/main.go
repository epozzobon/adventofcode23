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

type dir struct {
	r int
	c int
}

type point struct {
	r       int
	c       int
	visited bool
	arcs    []*path
	tile    byte
}

type path struct {
	destination *point
	length      int
}

func removeEdgePoint(pt *point) bool {
	if pt == nil || 2 != len(pt.arcs) || pt.tile != '.' {
		// This function only works on points that have 2 connections
		return false
	}
	n0 := pt.arcs[0].destination
	n1 := pt.arcs[1].destination

	findArc := func(arcs []*path, pt *point) *path {
		var found *path
		for _, k := range arcs {
			if k.destination == pt {
				if found != nil {
					panic("Duplicate found")
				}
				found = k
			}
		}
		return found
	}

	a0 := findArc(n0.arcs, pt)
	a1 := findArc(n1.arcs, pt)
	if a0 == nil || a1 == nil {
		return false
	}

	//   ____                 ____                 ____
	//  |    |-------a0----->|    |<------a0------|    |
	//  | n0 |               | pt |               | n1 |
	//  |____|<--pt.arcs[0]--|____|--pt.arcs[1]-->|____|
	//
	//  We want to delete pt

	a0.destination = n1
	a1.destination = n0
	a0.length += pt.arcs[1].length
	a1.length += pt.arcs[0].length
	return true
}

func day23(filename string, part2 bool) int {
	matrix, err := readMatrix(filename)
	if err != nil {
		panic(err)
	}

	pmatrix := make([][]*point, len(matrix))
	var entrance *point
	var exit *point
	for i, p := range matrix {
		pmatrix[i] = make([]*point, len(p))
		for j, q := range p {
			if q != '#' {
				pt := point{r: i, c: j, tile: q}
				if part2 {
					pt.tile = '.'
				}
				pmatrix[i][j] = &pt
				if i == 0 {
					entrance = &pt
				} else if i == len(matrix)-1 {
					exit = &pt
				}
			}
		}
	}
	for i := range pmatrix {
		for j, x := range pmatrix[i] {
			if x == nil {
				continue
			}

			dirMap := map[byte]([]dir){
				'>': []dir{{0, 1}},
				'<': []dir{{0, -1}},
				'v': []dir{{1, 0}},
				'^': []dir{{-1, 0}},
				'.': []dir{{1, 0}, {0, 1}, {0, -1}, {-1, 0}},
			}

			options := dirMap[x.tile]
			if len(options) == 0 {
				panic("Unknown tile")
			}

			r, c := 0, 0
			for _, d := range options {
				offr, offc := d.r, d.c
				r, c = i+offr, j+offc
				if r >= 0 && c >= 0 && r < len(pmatrix) && c < len(pmatrix[r]) {
					if pmatrix[r][c] != nil {
						p := path{pmatrix[r][c], 1}
						x.arcs = append(x.arcs, &p)
					}
				}
			}
			if 0 == len(x.arcs) {
				panic("Orphaned node")
			}
		}
	}

	for i := 0; i < len(pmatrix); i++ {
		for j := 0; j < len(pmatrix[i]); j++ {
			pt := pmatrix[i][j]
			if pt != nil && len(pt.arcs) == 0 {
				panic("Orphaned node")
			}
			if removeEdgePoint(pt) {
				pmatrix[i][j] = nil
				i, j = 0, 0
			}
		}
	}

	plist := []*point{}
	for i := 0; i < len(pmatrix); i++ {
		for j := 0; j < len(pmatrix[i]); j++ {
			if pmatrix[i][j] != nil {
				plist = append(plist, pmatrix[i][j])
			}
		}
	}

	var explore func(steps int, p *point) int
	explore = func(steps int, pt *point) int {

		if pt.visited {
			return -1
		}
		if pt == exit {
			return steps
		}

		pt.visited = true
		out := -1
		for _, path := range pt.arcs {
			out = max(out, explore(steps+path.length, path.destination))
		}
		pt.visited = false
		return out
	}
	return explore(0, entrance)
}

func main() {
	fmt.Println("Example, Part 1 =", day23("example1.txt", false))
	fmt.Println("Example, Part 2 =", day23("example1.txt", true))
	fmt.Println("Real Input, Part 1 =", day23("input.txt", false))
	fmt.Println("Real Input, Part 2 =", day23("input.txt", true))
}
