package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func FindDigit(txt string, reverse bool) rune {
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
		for j, v := range letterdigits {
			if i+len(v) < len(txt) {
				if txt[i:i+len(v)] == v {
					return rune('1' + j)
				}
			}
		}
	}
	return ' '
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		txt := scanner.Text()
		firstDigit := FindDigit(txt, false)
		lastDigit := FindDigit(txt, true)
		i, _ := strconv.Atoi(string(firstDigit) + string(lastDigit))
		sum += i
		fmt.Println(i)
	}
	fmt.Println(sum)

	err1 := scanner.Err()
	if err1 != nil {
		log.Fatal(err1)
	}
}
