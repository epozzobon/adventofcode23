package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main1() {
	digits := "0123456789"

	file, err := os.Open("example1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		txt := scanner.Text()
		firstDigit := ' '
		lastDigit := ' '
		for i := range txt {
			if strings.ContainsRune(digits, rune(txt[i])) {
				firstDigit = rune(txt[i])
				break
			}
		}
		for i := range txt {
			j := len(txt) - 1 - i
			if strings.ContainsRune(digits, rune(txt[j])) {
				lastDigit = rune(txt[j])
				break
			}
		}
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
