package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hash(s string) byte {
	var currentValue byte
	currentValue = 0
	for _, c := range s {
		currentValue += byte(c)
		currentValue *= 17
	}
	fmt.Println(currentValue)
	return currentValue
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	txt := scanner.Text()

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	steps := strings.Split(txt, ",")

	sum := 0
	for _, s := range steps {
		sum += int(hash(s))
	}
	fmt.Println(sum)
}
