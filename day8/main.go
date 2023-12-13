package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
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

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nodes := make(map[string]*node)
	lines := []nodeLine{}

	scanner.Scan()
	instructions := scanner.Text()

	scanner.Scan()
	empty := scanner.Text()
	if empty != "" {
		panic("expected empty line")
	}
	for scanner.Scan() {
		txt := scanner.Text()
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

	err1 := scanner.Err()
	if err1 != nil {
		log.Fatal(err1)
	}

	for _, line := range lines {
		nodes[line.src].l = nodes[line.left]
		nodes[line.src].r = nodes[line.right]
	}

	startNodes := []*node{}
	for _, line := range lines {
		if line.src[2] == 'A' {
			startNodes = append(startNodes, nodes[line.src])
		}
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
			for currentNode.name[2] != 'Z' {
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
		log.Printf("%q ", node.name)
	}
	log.Println(periods)
	log.Println(steps[0])
}
