package main

import (
	"bufio"
	"fmt"
	"os"
)

func rotate(matrix [][]byte, out [][]byte) {
	H := len(matrix[0])
	V := len(matrix)
	if len(out) != H {
		panic("bad output size")
	}
	if len(out[0]) != V {
		panic("bad output size")
	}
	for c := 0; c < H; c++ {
		out[c] = make([]byte, V)
	}
	for r := 0; r < H; r++ {
		for c := 0; c < V; c++ {
			out[r][c] = matrix[V-1-c][r]
		}
	}
}

func pushRocksNorth(matrix [][]byte) {
	for r := 1; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == 'O' {
				i := r - 1
				for i >= 0 && matrix[i][c] == '.' {
					i--
				}
				i++
				matrix[r][c] = '.'
				matrix[i][c] = 'O'
			}
		}
	}
}

func matrix2str(in [][]byte) string {
	H := len(in[0])
	V := len(in)
	key := ""
	for r := 0; r < V; r++ {
		if len(in[r]) != H {
			panic("lmao")
		}
		key += string(in[r])
	}
	return key
}

func cycle(in [][]byte) [][]byte {
	H := len(in[0])
	V := len(in)

	tmp := make([][]byte, H)
	for c := 0; c < H; c++ {
		tmp[c] = make([]byte, V)
	}

	out := make([][]byte, V)
	for r := 0; r < V; r++ {
		out[r] = make([]byte, H)
		copy(out[r], in[r])
	}

	pushRocksNorth(out)
	rotate(out, tmp)
	pushRocksNorth(tmp)
	rotate(tmp, out)
	pushRocksNorth(out)
	rotate(out, tmp)
	pushRocksNorth(tmp)
	rotate(tmp, out)

	return out
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := []([]byte){}
	for scanner.Scan() {
		txt := scanner.Text()
		matrix = append(matrix, []byte(txt))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	type MemoLine struct {
		matrix [][]byte
		weight int
		next   string
		index  int
	}
	memo := make(map[string]*MemoLine)
	l := 0
	var line *MemoLine
	for l = 0; l < 1000000000; l++ {
		key := matrix2str(matrix)
		if l > 0 {
			line.next = key
		}
		line = memo[key]
		if line != nil {
			break
		}
		sum := 0
		matrix = cycle(matrix)
		for r := 0; r < len(matrix); r++ {
			w := len(matrix) - r
			// fmt.Printf("%s %d\n", matrix[r], w)
			rowWeight := 0
			for c := 0; c < len(matrix[r]); c++ {
				if matrix[r][c] == 'O' {
					rowWeight += w
				}
			}
			sum += rowWeight
		}
		line = &MemoLine{matrix, sum, "", l}
		memo[key] = line
		fmt.Println(l+1, line.weight)
	}
	fmt.Println(l+1, line.weight)
	fmt.Println(l+1, "=", line.index+1)
	period := l - line.index
	fmt.Println("period is", period)
	fmt.Println(l+1+period, line.weight)
	fmt.Println(l+1+period*2, line.weight)
	l += ((1000000000 - l) / period) * period
	for l < 1000000000 {
		fmt.Println(l+1, line.weight)
		line = memo[line.next]
		l++
	}

}
