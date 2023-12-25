package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"epozzobon.it/adventofcode23/utils"
)

const (
	cN = 3
	cW = 2
	cS = 1
	cE = 0
)

type instruction struct {
	reg int
	cmp string
	imm int
	dst string
}

type regs [4]int

type workflow struct {
	name     string
	ins      []instruction
	fallback string
}

type regsRange struct {
	min regs
	max regs
}

func (rr *regsRange) isPossible() bool {
	for i := range rr.max {
		if rr.max[i] <= rr.min[i] {
			return false
		}
	}
	return true
}

func (rr *regsRange) size() int {
	size := 1
	for i := range rr.max {
		if rr.max[i] <= rr.min[i] {
			return 0
		}
		size *= rr.max[i] - rr.min[i] + 1
	}
	return size
}

func Day19(filename string, problemPart int) int {
	lines, err := utils.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	re1 := regexp.MustCompile(`^([a-z]+)\{(.+)\}$`)
	re2 := regexp.MustCompile(`^([xmas])([><])(\d+):?([ARa-z]+)$`)
	workflows := make(map[string]workflow)
	sep := slices.Index(lines, "")
	if sep == -1 {
		panic("Could not find empty line")
	}

	for _, txt := range lines[:sep] {
		pieces := re1.FindSubmatch([]byte(txt))
		workflowName := string(pieces[1])
		workflowString := string(pieces[2])
		spieces := strings.Split(workflowString, ",")

		instructions := make([]instruction, len(spieces)-1)
		for i, s := range spieces[:len(spieces)-1] {
			m := re2.FindSubmatch([]byte(s))
			reg := map[string]int{"x": 0, "m": 1, "a": 2, "s": 3}[string(m[1])]
			cmp := string(m[2])
			imm, err := strconv.Atoi(string(m[3]))
			if err != nil {
				panic("not an int")
			}
			dst := string(m[4])
			instructions[i] = instruction{reg, cmp, imm, dst}
		}
		fallback := spieces[len(spieces)-1]
		w := workflow{workflowName, instructions, fallback}
		workflows[workflowName] = w
	}

	fmt.Println(workflows)

	re3 := regexp.MustCompile(`^{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}$`)
	sum := 0
	for _, txt := range lines[sep+1:] {
		pieces := re3.FindSubmatch([]byte(txt))
		x, _ := strconv.Atoi(string(pieces[1]))
		m, _ := strconv.Atoi(string(pieces[2]))
		a, _ := strconv.Atoi(string(pieces[3]))
		s, _ := strconv.Atoi(string(pieces[4]))
		state := regs{x, m, a, s}
		workflow := workflows["in"]
		for true {
			dst := workflow.fallback
			for _, ins := range workflow.ins {
				v := state[ins.reg]
				if (v > ins.imm && ins.cmp == ">") || (v < ins.imm && ins.cmp == "<") {
					dst = ins.dst
					break
				}
			}
			if dst == "A" {
				fmt.Println("Accepted", x, m, a, s)
				sum += x + m + a + s
				break
			} else if dst == "R" {
				fmt.Println("Rejected", x, m, a, s)
				break
			} else {
				workflow = workflows[dst]
				if workflow.name == "" {
					panic("bad workflow")
				}
			}
		}
	}

	if problemPart == 1 {
		return sum
	}

	rRange := regsRange{regs{1, 1, 1, 1}, regs{4000, 4000, 4000, 4000}}
	var rrr func(regsRange, workflow)

	accepted := []regsRange{}
	rejected := []regsRange{}

	rrr = func(rr regsRange, w workflow) {
		for _, ins := range w.ins {
			if rr.size() == 0 {
				return
			}
			nr := rr
			if ins.cmp == ">" {
				// example: x > 1000 ->
				// if (rr.min.x <= 1000)  <-- checks if it is even possible
				// nr = x > 1000  -> nr.min.x = 1001
				// rr = x <= 1000 -> rr.max.x = 1000
				if rr.max[ins.reg] <= ins.imm {
					continue // condition never matched
				}
				rr.max[ins.reg] = ins.imm
				nr.min[ins.reg] = ins.imm + 1
			} else if ins.cmp == "<" {
				// example: x < 1000 ->
				// if (rr.max.x >= 1000)  <-- checks if it is even possible
				// nr = x < 1000  -> nr.max.x = 999
				// rr = x >= 1000 -> rr.min.x = 1000
				if rr.min[ins.reg] >= ins.imm {
					continue // condition never matched
				}
				rr.min[ins.reg] = ins.imm
				nr.max[ins.reg] = ins.imm - 1
			} else {
				panic("wtf")
			}
			if ins.dst == "A" {
				accepted = append(accepted, nr)
			} else if ins.dst == "R" {
				rejected = append(rejected, nr)
			} else {
				rrr(nr, workflows[ins.dst])
			}
		}
		if rr.size() > 0 {
			if w.fallback == "A" {
				accepted = append(accepted, rr)
			} else if w.fallback == "R" {
				rejected = append(rejected, rr)
			} else {
				rrr(rr, workflows[w.fallback])
			}
		}
	}
	rrr(rRange, workflows["in"])

	fmt.Println(accepted)

	possibilities := 0
	for _, a := range accepted {
		possibilities += a.size()
	}
	fmt.Println(possibilities)

	if problemPart == 2 {
		return possibilities
	}

	panic("Unknown problem part")
}

func main() {
	utils.CheckSolution(Day19, "example1.txt", 1, 19114)
	utils.CheckSolution(Day19, "example1.txt", 2, 167409079868000)
	utils.CheckSolution(Day19, "input.txt", 1, 495298)
	utils.CheckSolution(Day19, "input.txt", 2, 132186256794011)
}
