package main

import (
	"fmt"
	"regexp"

	"epozzobon.it/adventofcode23/utils"
)

type nodeLine struct {
	src   string
	left  string
	right string
}

type node struct {
	name string
	l    *node
	r    *node
}

type skip struct {
	node  *node
	steps int
}

var reLine = regexp.MustCompile(`^([0-9A-Z]+)\s*=\s*\(([0-9A-Z]+),\s*([0-9A-Z]+)\)$`)

func Day8(filepath string, problemPart int) int {

	texts, err := utils.ReadLines(filepath)
	if err != nil {
		panic(err)
	}

	instructions := texts[0]
	if texts[1] != "" {
		panic("Expected empty second line")
	}

	nodes := make(map[string]*node)
	lines := []nodeLine{}

	for _, txt := range texts[2:] {
		matches := reLine.FindSubmatch([]byte(txt))
		if len(matches) != 4 {
			panic("bad input line format")
		}
		src, left, right := string(matches[1]), string(matches[2]), string(matches[3])
		lines = append(lines, nodeLine{src: src, left: left, right: right})
		if nodes[src] != nil {
			panic("Node already defined")
		}
		nodes[src] = &node{name: src, l: nil, r: nil}
	}

	for _, line := range lines {
		nodes[line.src].l = nodes[line.left]
		nodes[line.src].r = nodes[line.right]
	}

	var isFinal func(name string) bool
	startNodes := []*node{}
	if problemPart == 1 {
		for _, line := range lines {
			if line.src == "AAA" {
				startNodes = append(startNodes, nodes[line.src])
			}
		}
		isFinal = func(name string) bool {
			return name == "ZZZ"
		}
	} else if problemPart == 2 {
		for _, line := range lines {
			if line.src[2] == 'A' {
				startNodes = append(startNodes, nodes[line.src])
			}
		}
		isFinal = func(name string) bool {
			return name[2] == 'Z'
		}
	} else {
		panic("Unknown problem part")
	}
	currentNodes := make([]*node, len(startNodes))
	copy(currentNodes, startNodes)

	memo := make(map[*node]*node)

	applyInstructions := func(startNode *node) *node {
		if memo[startNode] == nil {
			currentNode := startNode
			for i := 0; i < len(instructions); i++ {
				c := instructions[i]
				if c == 'L' {
					currentNode = currentNode.l
				} else if c == 'R' {
					currentNode = currentNode.r
				} else {
					panic("bad instruction")
				}
			}
			memo[startNode] = currentNode
		}
		return memo[startNode]
	}

	memoo := make(map[*node]*skip)

	applyUntilZ := func(startNode *node) (*node, int) {
		if memoo[startNode] == nil {
			currentNode := applyInstructions(startNode)
			steps := len(instructions)
			for !isFinal(currentNode.name) {
				currentNode = applyInstructions(currentNode)
				steps += len(instructions)
			}
			memoo[startNode] = &skip{node: currentNode, steps: steps}
		}
		s := memoo[startNode]
		return s.node, s.steps
	}

	steps := make([]int, len(startNodes))
	periods := make([]int, len(startNodes))
	finishes := 0

	for finishes < len(currentNodes) {
		finishes = 0
		for i := 0; i < len(currentNodes); i++ {

			j := (i + 1) % len(steps)
			if steps[i] > 0 && steps[i] == steps[j] {
				finishes++
			} else {
				p, q := j, i
				if steps[i] < steps[j] {
					p, q = i, j
				}

				if periods[q] > 0 && periods[p] > 0 {
					diff := steps[q] - steps[p]
					ceil := (diff + periods[p] - 1) / periods[p]
					steps[p] += ceil * periods[p]
				} else {
					n, s := applyUntilZ(currentNodes[p])
					steps[p] += s
					currentNodes[p] = n
					if n == currentNodes[p] {
						periods[p] = s
					}
				}

			}
		}
	}

	for _, node := range currentNodes {
		fmt.Printf("%q ", node.name)
	}
	fmt.Println(periods)
	fmt.Println(steps[0])

	return steps[0]
}

func main() {
	utils.CheckSolution(Day8, "example1.txt", 1, 2)
	utils.CheckSolution(Day8, "example2.txt", 1, 6)
	utils.CheckSolution(Day8, "example3.txt", 2, 6)
	utils.CheckSolution(Day8, "input.txt", 1, 22411)
	utils.CheckSolution(Day8, "input.txt", 2, 11188774513823)
}
