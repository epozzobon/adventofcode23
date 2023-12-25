package main

import (
	"fmt"
	"strconv"
	"strings"

	"epozzobon.it/adventofcode23/utils"
)

const (
	cN = 3
	cW = 2
	cS = 1
	cE = 0
)

type instruction struct {
	direction int
	length    int
}

func Day18(filename string, problemPart int) int {
	instructions := []instruction{}

	lines, err := utils.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	for _, txt := range lines {
		pieces := strings.Split(txt, " ")
		if err != nil {
			panic(err)
		}

		var direction int
		var length int
		if problemPart == 2 {
			v, err := strconv.ParseInt(pieces[2][2:7], 16, 32)
			length = int(v)
			if err != nil {
				panic(err)
			}
			v, err = strconv.ParseInt(pieces[2][7:8], 16, 16)
			direction = int(v)
			if err != nil {
				panic(err)
			}
		} else if problemPart == 1 {
			direction = (map[string]int{
				"R": cE,
				"L": cW,
				"U": cN,
				"D": cS,
			})[pieces[0]]
			length, err = strconv.Atoi(pieces[1])
			if err != nil {
				panic(err)
			}
		} else {
			panic("Unknown problem part")
		}
		instr := instruction{int(direction), int(length)}
		instructions = append(instructions, instr)
	}

	fmt.Println(instructions)
	return solve(instructions)
}

func checkLoop(instructions []instruction) {
	x := 0
	y := 0
	for _, instr := range instructions {
		if instr.direction == cE {
			x += instr.length
		} else if instr.direction == cW {
			x -= instr.length
		} else if instr.direction == cN {
			y -= instr.length
		} else if instr.direction == cS {
			y += instr.length
		}
	}
	if x != 0 || y != 0 {
		fmt.Println("ERROR: not a loop")
		fmt.Println(instructions)
		panic("I did not get back to the starting point")
	}
}

func fixupInstructions(instructions []instruction) []instruction {

	for i := 0; i < len(instructions); i++ {
		ni := (i + 1) % len(instructions)
		if instructions[i].length == 0 {
			instructions = append(instructions[:i], instructions[i+1:]...)
			return fixupInstructions(instructions)
		} else if instructions[i].direction == instructions[ni].direction {
			instructions[ni].length += instructions[i].length
			instructions = append(instructions[:i], instructions[i+1:]...)
			return fixupInstructions(instructions)
		}
	}
	return instructions
}

func solve(instructions []instruction) int {
	checkLoop(instructions)

	di := instructions[0].direction
	for i := range instructions {
		instructions[i].direction += 4 - di
		instructions[i].direction %= 4
	}
	fmt.Println("pre-fix:", instructions)
	checkLoop(instructions)
	instructions = fixupInstructions(instructions)
	fmt.Println("fixed:  ", instructions)
	checkLoop(instructions)

	if len(instructions) == 0 {
		return 1
	}

	for i := 0; i < len(instructions)-1; i++ {
		if ((instructions[i].direction + 2) % 4) == instructions[i+1].direction {
			// #########    => #######++
			// ......#         ......#
			// ......#      => ......#
			// ......#         ......#
			// ......#         ......#
			//   8, 2, 4  =>      6, 4 ; (8-2)
			//   R  L  D          R  D
			if instructions[i].length > instructions[i+1].length {
				instructions[i].length -= instructions[i+1].length
				area := instructions[i+1].length
				instructions = append(instructions[:i+1], instructions[i+2:]...)
				fmt.Println(area)
				return area + solve(instructions)
			} else {
				instructions[i+1].length -= instructions[i].length
				area := instructions[i].length
				instructions = append(instructions[:i], instructions[i+1:]...)
				fmt.Println(area)
				return area + solve(instructions)
			}
		}
	}

	bestI := -1
	for i := 0; i < len(instructions); i++ {
		i1 := (i + 1) % len(instructions)
		i2 := (i + 2) % len(instructions)
		d0 := instructions[i].direction
		d1 := instructions[i1].direction
		d2 := instructions[i2].direction
		if ((d2-d0+4)%4) == 2 && ((d1-d0+4)%4) == 1 {
			if bestI == -1 || instructions[bestI].length > instructions[i1].length {
				bestI = i1
			}
		}
	}
	if bestI == -1 {
		panic("erm")
	}

	i := (bestI + len(instructions) - 1) % len(instructions)
	i1 := bestI
	i2 := (bestI + 1) % len(instructions)
	// #########          ######+++
	// ........#          .....#+++
	// ........#      =>  .....#+++
	// ........#          .....#+++
	// .....####          .....#+++
	// .....#             .....#
	//   8, 4, 3  =>      5, 4 ; ((4+1)*3)
	//   R  D  L          R  D
	if instructions[i].length > instructions[i2].length {
		area := (instructions[i1].length + 1) * instructions[i2].length
		instructions[i].length -= instructions[i2].length
		instructions = append(instructions[:i2], instructions[i2+1:]...)
		fmt.Println(area)
		return area + solve(instructions)
	} else {
		area := (instructions[i1].length + 1) * instructions[i].length
		instructions[i2].length -= instructions[i].length
		instructions = append(instructions[:i], instructions[i+1:]...)
		fmt.Println(area)
		return area + solve(instructions)
	}
}

func main() {
	utils.CheckSolution(Day18, "example1.txt", 1, 62)
	utils.CheckSolution(Day18, "example1.txt", 2, 952408144115)
	utils.CheckSolution(Day18, "input.txt", 1, 50746)
	utils.CheckSolution(Day18, "input.txt", 2, 70086216556038)
}
