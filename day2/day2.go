package main

import (
	"fmt"
	"strconv"
	"strings"

	"epozzobon.it/adventofcode23/utils"
)

func Day2(filepath string, problemPart int) int {

	file, err := utils.ReadLines(filepath)
	if err != nil {
		panic(err)
	}

	sum1, sum2 := 0, 0
	for _, txt := range file {
		fmt.Println(txt)

		gameParts := strings.Split(txt, ":")
		if len(gameParts) != 2 {
			panic("bad game string")
		}

		gameName, gameString := gameParts[0], gameParts[1]

		gameNameParts := strings.Split(gameName, " ")
		if len(gameNameParts) != 2 || gameNameParts[0] != "Game" {
			panic("bad game name")
		}

		gameID, err := strconv.Atoi(gameNameParts[1])
		if err != nil {
			panic("bad game ID")
		}

		subgames := strings.Split(gameString, ";")
		maxR, maxG, maxB := 0, 0, 0
		for _, subgame := range subgames {
			r, g, b := 0, 0, 0
			// A single subgame, made up of at most 3 draws of different colors
			draws := strings.Split(subgame, ",")
			for _, draw := range draws {
				// A single draw of a specific color in a game
				drawParts := strings.Split(strings.Trim(draw, " "), " ")
				if len(drawParts) != 2 {
					panic("bad draw string")
				}
				color := drawParts[1]
				n, err := strconv.Atoi(drawParts[0])
				if err != nil {
					panic("bad number")
				}
				if color == "red" {
					r += n
				} else if color == "green" {
					g += n
				} else if color == "blue" {
					b += n
				} else {
					panic("bad color")
				}
			}

			maxR = max(r, maxR)
			maxG = max(g, maxG)
			maxB = max(b, maxB)
		}
		_ = gameID
		sum2 += maxR * maxG * maxB

		if maxR <= 12 && maxG <= 13 && maxB <= 14 {
			sum1 += gameID
		}
	}
	fmt.Println(sum1, sum2)

	if problemPart == 1 {
		return sum1
	} else if problemPart == 2 {
		return sum2
	} else {
		panic("Unknown problem part")
	}
}

func main() {
	utils.CheckSolution(Day2, "example1.txt", 1, 8)
	utils.CheckSolution(Day2, "example1.txt", 2, 2286)
	utils.CheckSolution(Day2, "input.txt", 1, 2449)
	utils.CheckSolution(Day2, "input.txt", 2, 63981)
}
