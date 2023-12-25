package main

import (
	"fmt"
	"strconv"

	"epozzobon.it/adventofcode23/utils"
)

type potgear struct {
	x    int
	y    int
	pnum int
}

func Day3(filepath string, problemPart int) int {
	schematics, err := utils.ReadMatrix(filepath)
	if err != nil {
		panic(err)
	}

	columns := len(schematics[0])
	rows := len(schematics)
	schematics.Print()

	potgears := []potgear{}

	ratiosum, sum1 := 0, 0
	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			c := schematics[y][x]
			n := []byte{}
			for c >= '0' && c <= '9' {
				n = append(n, byte(c))
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
					line := []byte{}
					for x0 := bx1; x0 <= bx2; x0++ {
						r := byte(schematics[y0][x0])
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

	fmt.Println(sum1, ratiosum)
	if problemPart == 1 {
		return sum1
	} else if problemPart == 2 {
		return ratiosum
	} else {
		panic("Unknown problem part")
	}
}

func main() {
	utils.CheckSolution(Day3, "example1.txt", 1, 4361)
	utils.CheckSolution(Day3, "example1.txt", 2, 467835)
	utils.CheckSolution(Day3, "input.txt", 1, 528819)
	utils.CheckSolution(Day3, "input.txt", 2, 80403602)
}
