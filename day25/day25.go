package day25

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"epozzobon.it/adventofcode23/utils"
)

type v3 struct {
	x int
	y int
	z int
}
type v3f struct {
	x float64
	y float64
	z float64
}

type node struct {
	name string
	arcs []*node
}

func loadGraph(filename string) map[string]*node {

	re := regexp.MustCompile(`[a-z]+`)

	file, err := utils.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	nodes := map[string]*node{}

	getNode := func(name string) *node {
		if nodes[name] == nil {
			nodes[name] = &node{name: name}
		}
		return nodes[name]
	}
	for _, txt := range file {
		matches := re.FindAllString(txt, -1)
		src := getNode(matches[0])
		for _, dstName := range matches[1:] {
			dst := getNode(dstName)
			dst.arcs = append(dst.arcs, src)
			src.arcs = append(src.arcs, dst)
		}
	}

	return nodes
}

func Day25(filename string, problemPart int) int {

	nodes := loadGraph(filename)

	g0, g1 := findSplit(nodes, 3)
	fmt.Println(g0, g1)

	countNodes := func(g0 *node) int {
		stack := []*node{g0}
		counted := map[string]bool{}
		for len(stack) > 0 {
			src := stack[0]
			stack = stack[1:]
			for _, arc := range src.arcs {
				if !counted[arc.name] {
					counted[arc.name] = true
					stack = append(stack, arc)
				}
			}
		}
		return len(counted)
	}
	c0 := countNodes(g0)
	c1 := countNodes(g1)
	fmt.Println(c0)
	fmt.Println(c1)

	return c1 * c0
}

func disconnect(nodes map[string]*node, k0, k1 string) {
	disconnectImpl := func(k0, k1 string) {
		for i := 0; i < len(nodes[k0].arcs); i++ {
			if nodes[k0].arcs[i].name == k1 {
				nodes[k0].arcs = append(nodes[k0].arcs[:i], nodes[k0].arcs[i+1:]...)
				i--
			}
		}
	}
	disconnectImpl(k0, k1)
	disconnectImpl(k1, k0)
}

func findSplit(nodes map[string]*node, arcsCount int) (*node, *node) {

	var forbidden map[string]bool

	findGraphPath := func(src *node, dst *node) []*node {
		traversed := map[string]*node{src.name: src}
		stack := []*node{src}
		for len(stack) > 0 {
			src = stack[0]
			stack = stack[1:]
			for _, n := range src.arcs {
				if forbidden[src.name+":"+n.name] {
					continue
				}
				if traversed[n.name] == nil {
					traversed[n.name] = src
					stack = append(stack, n)
					if n.name == dst.name {
						path := []*node{dst}
						for traversed[dst.name] != dst {
							dst = traversed[dst.name]
							path = append(path, dst)
						}
						slices.Reverse(path)
						return path
					}
				}
			}
		}
		return []*node{}
	}

	keys := make([]string, 0, len(nodes))
	for k := range nodes {
		keys = append(keys, k)
	}

	n0 := nodes[keys[0]]
	for _, secondKey := range keys[1:] {
		forbidden = map[string]bool{}
		n1 := nodes[secondKey]
		for try := 0; try < arcsCount+1; try++ {
			fmt.Printf("%s -> %s : ", n0.name, n1.name)
			path := findGraphPath(n0, n1)
			if try == arcsCount && len(path) == 0 {
				fmt.Printf("%s and %s are connected by only %d wires\n", n0.name, n1.name, arcsCount)

				keys := make([]string, 0, len(forbidden))
				for k := range forbidden {
					keys = append(keys, k)
				}

				for _, k := range keys {
					if forbidden[k] == false {
						continue
					}
					ks := strings.Split(k, ":")
					k0, k1 := ks[0], ks[1]
					forbidden[k0+":"+k1] = false
					forbidden[k1+":"+k0] = false
					path := findGraphPath(n0, n1)
					if len(path) != 0 {
						disconnect(nodes, k0, k1)
						fmt.Printf("%s is essential\n", k)
					}
				}

				return n0, n1
			}
			for i := 0; i < len(path)-1; i++ {
				k := path[i]
				forbidden[path[i].name+":"+path[i+1].name] = true
				forbidden[path[i+1].name+":"+path[i].name] = true
				fmt.Printf("%s ", k.name)
			}
			fmt.Println()
		}
	}

	return nil, nil
}
