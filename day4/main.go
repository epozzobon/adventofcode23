package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func makeIntList(txt string) ([]int, error) {
	re := regexp.MustCompile(`\d+`)
	vs := re.FindAll([]byte(txt), -1)
	vsm := make([]int, len(vs))
	for i, v := range vs {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			return nil, err
		}
		vsm[i] = n
	}
	return vsm, nil
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum1, sum2 := 0, 0
	reLine := regexp.MustCompile(`^Card\s+(\d+):\s+([\d\s]+)\s+\|\s+([\d\s]+)\s*$`)
	ahead := []int{0, 0}
	idx := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		idx++
		txt := scanner.Text()
		matches := reLine.FindSubmatch([]byte(txt))
		if len(matches) == 0 {
			continue
		}

		cardID, err := strconv.Atoi(string(matches[1]))
		if err != nil {
			panic("bad card ID")
		}

		if cardID != idx {
			panic("wrong card ID")
		}

		winning, err := makeIntList(string(matches[2]))
		if err != nil {
			panic("bad winning numbers")
		}

		mynumbers, err := makeIntList(string(matches[3]))
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

		for len(ahead) <= idx+wins {
			ahead = append(ahead, 0)
		}

		ncards := 1 + ahead[idx]
		sum2 += ncards

		score := 0
		for i := 0; i < wins; i++ {
			ahead[idx+1+i] += ncards
			score *= 2
			if score == 0 {
				score = 1
			}
		}

		sum1 += score

		fmt.Println(cardID, ncards)
	}

	err1 := scanner.Err()
	if err1 != nil {
		panic(err1)
	}

	fmt.Println(sum1, sum2)
}
