package day1

import (
	"fmt"
	"strconv"
	"strings"

	"epozzobon.it/adventofcode23/utils"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func FindDigit(txt string, reverse bool, findText bool) rune {
	digits := "0123456789"
	letterdigits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	if reverse {
		txt = Reverse(txt)
		for j, v := range letterdigits {
			letterdigits[j] = Reverse(v)
		}
	}
	for i := range txt {
		if strings.ContainsRune(digits, rune(txt[i])) {
			return rune(txt[i])
		}
		if findText {
			for j, v := range letterdigits {
				if i+len(v) < len(txt) {
					if txt[i:i+len(v)] == v {
						return rune('1' + j)
					}
				}
			}
		}
	}
	return ' '
}

func Day1(filepath string, problemPart int) int {

	file, err := utils.ReadLines(filepath)
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, txt := range file {
		firstDigit := FindDigit(txt, false, problemPart == 2)
		lastDigit := FindDigit(txt, true, problemPart == 2)
		i, _ := strconv.Atoi(string(firstDigit) + string(lastDigit))
		sum += i
		fmt.Println(i)
	}
	fmt.Println(sum)
	return sum
}
