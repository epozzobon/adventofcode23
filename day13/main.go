package main

import (
	"bufio"
	"fmt"
	"os"
)

func getColumn(room []string, col int) string {
	out := make([]byte, len(room))
	for i, s := range room {
		out[i] = s[col]
	}
	return string(out)
}

func transpose(room []string) []string {
	out := make([]string, len(room[0]))
	for i := range room[0] {
		out[i] = getColumn(room, i)
	}
	return out
}

func countSmudges(r, q string) int {
	count := 0
	for i := range r {
		if r[i] != q[i] {
			count++
		}
	}
	return count
}

func evalHori(room []string) int {
	found := 0
	for i := range room[:len(room)-1] {
		j := i
		k := i + 1
		smudges := 0
		for j >= 0 && k < len(room) {
			smudges += countSmudges(room[j], room[k])
			j--
			k++
		}
		if smudges == 1 {
			if found != 0 {
				panic("found second mirror?")
			}
			found = i + 1
		}
	}
	return found
}

func evalRoom(room []string) int {
	hori := evalHori(room)
	vert := evalHori(transpose(room))
	if (hori == 0) == (vert == 0) {
		panic("either no or both hori and ver mirrors")
	}
	return hori*100 + vert
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	room := []string{}
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			sum += evalRoom(room)
			room = []string{}
		} else {
			room = append(room, txt)
		}
	}
	if len(room) > 0 {
		sum += evalRoom(room)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
