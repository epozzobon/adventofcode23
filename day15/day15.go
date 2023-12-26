package day15

import (
	"regexp"
	"strconv"
	"strings"

	"epozzobon.it/adventofcode23/utils"
)

func hash(s string) byte {
	var currentValue byte
	currentValue = 0
	for _, c := range s {
		currentValue += byte(c)
		currentValue *= 17
	}
	return currentValue
}

type lens struct {
	value int
	label string
}

func Day15(filepath string, problemPart int) int {
	lines, err := utils.ReadLines(filepath)
	if err != nil {
		panic(err)
	}
	txt := lines[0]

	steps := strings.Split(txt, ",")

	if problemPart == 1 {
		sum := 0
		for _, s := range steps {
			sum += int(hash(s))
		}
		return sum
	} else if problemPart == 2 {

		boxes := make([][]lens, 256)

		re := regexp.MustCompile(`^([a-z]+)([=-])(\d*)$`)
		for _, s := range steps {
			matches := re.FindSubmatch([]byte(s))
			label := string(matches[1])
			op := matches[2][0]
			imm := 0
			if len(matches[3]) > 0 {
				imm, err = strconv.Atoi(string(matches[3]))
				if err != nil {
					panic("Not a number")
				}
			}

			boxIdx := hash(label)
			lensIdx := -1
			for i, l := range boxes[boxIdx] {
				if l.label == label {
					lensIdx = i
					break
				}
			}
			if op == '-' {
				if lensIdx != -1 {
					boxes[boxIdx] = append(boxes[boxIdx][:lensIdx], boxes[boxIdx][lensIdx+1:]...)
				}
			} else if op == '=' {
				if lensIdx == -1 {
					lensIdx = len(boxes[boxIdx])
					boxes[boxIdx] = append(boxes[boxIdx], lens{imm, label})
				} else {
					boxes[boxIdx][lensIdx].value = imm
				}
			}
		}

		sum := 0
		for i := 0; i < 256; i++ {
			for j, b := range boxes[i] {
				prod := (i + 1) * b.value * (1 + j)
				sum += prod
			}
		}
		return sum
	} else {
		panic("Unknown problem part")
	}
}
