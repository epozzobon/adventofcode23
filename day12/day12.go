package day12

import (
	"log"
	"strconv"
	"strings"

	"epozzobon.it/adventofcode23/utils"
)

type onsen struct {
	line    []rune
	numbers []rune
}

var memo = make(map[string]*int)

func fulfil(line []rune, numbers []rune, o int) int {

	key := string(line[o:]) + ":" + string(numbers)
	if memo[key] != nil {
		return *memo[key]
	}

	if len(line) == o {
		if len(numbers) == 0 {
			return 1
		}
		return 0
	}

	if line[o] == '.' {
		return fulfil(line, numbers, o+1)
	}

	if len(numbers) == 0 {
		if line[o] == '#' {
			return 0
		}
		return fulfil(line, numbers, o+1)
	}

	if len(line)-o < int(numbers[0]) {
		return 0
	}

	if line[o] == '#' {
		i := o
		for i = o + 1; i < o+int(numbers[0]); i++ {
			if line[i] == '.' {
				return 0
			}
		}
		if len(numbers) == 1 {
			return fulfil(line, numbers[1:], i)
		}
		if i < len(line) && line[i] == '#' {
			return 0
		}
		if i < len(line)-1 {
			return fulfil(line, numbers[1:], i+1)
		}
		return 0
	}

	i1 := fulfil(line, numbers, o+1)
	line[o] = '#'
	i2 := fulfil(line, numbers, o)
	line[o] = '?'
	s := i1 + i2
	memo[key] = &s
	return s
}

func bito(line string, numbers []rune, reps int) int {
	oldline := line
	d := len(numbers)
	for i := 0; i < reps; i++ {
		line += "?"
		line += oldline
		for j := 0; j < d; j++ {
			numbers = append(numbers, numbers[j])
		}
	}

	o := onsen{line: []rune(line), numbers: numbers}

	return fulfil((o.line), (o.numbers), 0)
}

func Day12(filepath string, problemPart int) int {

	lines, err := utils.ReadLines(filepath)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, txt := range lines {
		parts := strings.Split(txt, " ")
		line := parts[0]
		parts = strings.Split(parts[1], ",")
		numbers := make([]rune, len(parts))
		for i, n := range parts {
			num, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			numbers[i] = rune(num)
		}
		//fmt.Println(line, numbers)

		var reps int
		if problemPart == 1 {
			reps = 0
		} else if problemPart == 2 {
			reps = 4
		} else {
			panic("Unknown problem part")
		}
		sum += bito(line, numbers, reps)
	}
	return sum
}
