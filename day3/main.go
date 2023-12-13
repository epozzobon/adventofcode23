package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type potgear struct {
	x    int
	y    int
	pnum int
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	schematics := []string{}
	scanner := bufio.NewScanner(file)
	columns := 0
	rows := 0
	for scanner.Scan() {
		rows++
		txt := scanner.Text()
		if len(schematics) == 0 {
			columns = len(txt)
		} else {
			if len(txt) != columns {
				panic("Different line lengths")
			}
		}
		schematics = append(schematics, txt)
	}
	fmt.Println(schematics)

	potgears := []potgear{}

	ratiosum, sum1 := 0, 0
	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			c := schematics[y][x]
			n := []rune{}
			for c >= '0' && c <= '9' {
				n = append(n, rune(c))
				if (x + len(n)) >= columns {
					break
				}
				c = schematics[y][x+len(n)]
			}

			if len(n) > 0 {
				num, err := strconv.Atoi(string(n))
				if err != nil {
					panic("bad number")
				}

				// number detected, is it a part number?
				bx1 := max(0, x-1)
				bx2 := min(x+len(n), columns-1)
				by1 := max(0, y-1)
				by2 := min(y+1, rows-1)

				for y0 := by1; y0 <= by2; y0++ {
					line := []rune{}
					for x0 := bx1; x0 <= bx2; x0++ {
						r := rune(schematics[y0][x0])
						line = append(line, r)
						if r != '.' && (r > '9' || r < '0') {
							sum1 += num
						}
						if r == '*' {
							// Is a potential gear

							for _, p := range potgears {
								if p.x == x0 && p.y == y0 {
									ratio := p.pnum * num
									ratiosum += ratio
								}
							}

							p := potgear{x0, y0, num}
							potgears = append(potgears, p)
						}
					}
				}

				x += len(n) - 1
			}
		}
	}

	err1 := scanner.Err()
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(sum1, ratiosum)
}
