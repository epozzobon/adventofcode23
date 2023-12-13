package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func toIntList(txt string) ([]int, error) {
	vs := strings.Split(txt, " ")
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

func derive(line []int) []int {
	derivative := make([]int, len(line)-1)
	for c := 0; c < len(line)-1; c++ {
		derivative[c] = line[c+1] - line[c]
	}
	return derivative
}

func predict(line []int) (int, int) {
	if len(line) == 0 {
		return 0, 0
	}
	derivative := derive(line)
	d0, d1 := predict(derivative)
	return line[0] - d0, line[len(line)-1] + d1
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum1, sum2 := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		row, err := toIntList(txt)
		if err != nil {
			panic("integer expected")
		}
		prev, next := predict(row)
		fmt.Println(prev, ",", row, ",", next)
		sum1 += prev
		sum2 += next
	}

	err = scanner.Err()
	if err != nil {
		panic("input read error")
	}

	fmt.Println(sum1, sum2)
}
