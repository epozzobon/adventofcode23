package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"epozzobon.it/adventofcode23/utils"
)

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

func Day9(filepath string) (int, int) {

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum1, sum2 := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		row, err := utils.SpacedIntsList(txt)
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
	return sum2, sum1
}
