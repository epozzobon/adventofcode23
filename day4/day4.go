package main

import (
	"fmt"
	"regexp"
	"strconv"

	"epozzobon.it/adventofcode23/utils"
)

func Day4(filepath string, problemPart int) int {

	file, err := utils.ReadLines(filepath)
	if err != nil {
		panic(err)
	}

	sum1, sum2 := 0, 0
	reLine := regexp.MustCompile(`^Card\s+(\d+):\s+([\d\s]+)\s+\|\s+([\d\s]+)\s*$`)
	ahead := []int{0, 0}

	for idx, txt := range file {
		matches := reLine.FindSubmatch([]byte(txt))
		if len(matches) == 0 {
			continue
		}

		cardID, err := strconv.Atoi(string(matches[1]))
		if err != nil {
			panic("bad card ID")
		}

		if cardID != idx+1 {
			panic("wrong card ID")
		}

		winning, err := utils.SpacedIntsList(string(matches[2]))
		if err != nil {
			panic("bad winning numbers")
		}

		mynumbers, err := utils.SpacedIntsList(string(matches[3]))
		if err != nil {
			panic("bad own numbers")
		}

		wins := 0
		for _, i := range mynumbers {
			for _, j := range winning {
				if i == j {
					wins++
					break
				}
			}
		}

		for len(ahead) <= idx+1+wins {
			ahead = append(ahead, 0)
		}

		ncards := 1 + ahead[idx+1]
		sum2 += ncards

		score := 0
		for i := 0; i < wins; i++ {
			ahead[idx+2+i] += ncards
			score *= 2
			if score == 0 {
				score = 1
			}
		}

		sum1 += score

		fmt.Println(cardID, ncards)
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
	utils.CheckSolution(Day4, "example1.txt", 1, 13)
	utils.CheckSolution(Day4, "example1.txt", 2, 30)
	utils.CheckSolution(Day4, "input.txt", 1, 19855)
	utils.CheckSolution(Day4, "input.txt", 2, 10378710)
}
