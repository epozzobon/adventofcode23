package main

import (
	"fmt"

	"epozzobon.it/adventofcode23/utils"
)

func countSmudges(r, q []byte) int {
	count := 0
	for i := range r {
		if r[i] != q[i] {
			count++
		}
	}
	return count
}

func evalHori(room utils.Matrix, smudgesNeeded int) int {
	found := 0
	for i := range room[:len(room)-1] {
		j := i
		k := i + 1
		smudges := 0
		for j >= 0 && k < len(room) {
			smudges += countSmudges(room[j], room[k])
			j--
			k++
		}
		if smudges == smudgesNeeded {
			if found != 0 {
				panic("found second mirror?")
			}
			found = i + 1
		}
	}
	return found
}

func evalRoom(room utils.Matrix, smudgesNeeded int) int {
	hori := evalHori(room, smudgesNeeded)
	vert := evalHori(room.Transpose(), smudgesNeeded)
	if (hori == 0) == (vert == 0) {
		panic("either no or both hori and ver mirrors")
	}
	return hori*100 + vert
}

func Day13(filepath string, problemPart int) int {

	var smudgesNeeded int
	if problemPart == 1 {
		smudgesNeeded = 0
	} else if problemPart == 2 {
		smudgesNeeded = 1
	} else {
		panic("Unknown problem part")
	}

	lines, err := utils.ReadLines(filepath)
	if err != nil {
		panic(err)
	}
	sum := 0
	room := utils.Matrix{}
	for _, txt := range lines {
		if txt == "" {
			sum += evalRoom(room, smudgesNeeded)
			room = utils.Matrix{}
		} else {
			room = append(room, []byte(txt))
		}
	}
	if len(room) > 0 {
		sum += evalRoom(room, smudgesNeeded)
	}

	fmt.Println(sum)
	return sum
}

func main() {
	utils.CheckSolution(Day13, "example1.txt", 1, 405)
	utils.CheckSolution(Day13, "example1.txt", 2, 400)
	utils.CheckSolution(Day13, "input.txt", 1, 31265)
	utils.CheckSolution(Day13, "input.txt", 2, 39359)
}
