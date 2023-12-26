package day22

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"epozzobon.it/adventofcode23/utils"
)

type v3 struct {
	x int
	y int
	z int
}

type brick struct {
	p0        v3
	p1        v3
	id        int
	supports  map[int]*brick
	supported map[int]*brick
	flag      bool
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (self brick) size() v3 {
	return v3{(self.p0.x - self.p1.x), (self.p0.y - self.p1.y), (self.p0.z - self.p1.z)}
}

func makeStrId(i int) string {
	return string([]rune{rune('A' + i%26), rune('A' + (i/26)%26), rune('A' + (i/26/26)%26)})
}

type plane [][]*brick

func (zplane plane) print() {
	for x := 0; x < len(zplane); x++ {
		for y := 0; y < len(zplane[x]); y++ {
			bid := "   "
			if zplane[x][y] != nil {
				bid = makeStrId(zplane[x][y].id)
			}
			fmt.Printf("%s,", bid)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func gravity(bricks []*brick) {
	maxZ := bricks[0].p1.z
	maxX := bricks[0].p1.x
	maxY := bricks[0].p1.y
	for i := 0; i < len(bricks); i++ {
		maxX = max(bricks[i].p1.x, maxX)
		maxY = max(bricks[i].p1.y, maxY)
		maxZ = max(bricks[i].p1.z, maxZ)
	}

	zplane := make(plane, maxX+1)
	for x := 0; x <= maxX; x++ {
		zplane[x] = make([]*brick, maxY+1)
	}

	sort.SliceStable(bricks, func(i, j int) bool {
		return bricks[i].p0.z < bricks[j].p0.z
	})
	for i := 0; i < len(bricks); i++ {
		b := bricks[i]
		maxZ := -1
		for x := b.p0.x; x <= b.p1.x; x++ {
			for y := b.p0.y; y <= b.p1.y; y++ {
				if zplane[x][y] != nil {
					if maxZ == zplane[x][y].p1.z {
						b.supports[zplane[x][y].id] = zplane[x][y]
					} else if maxZ < zplane[x][y].p1.z {
						maxZ = zplane[x][y].p1.z
						b.supports = map[int]*brick{zplane[x][y].id: zplane[x][y]}
					}
				}
				zplane[x][y] = b
			}
		}
		offZ := b.p0.z - maxZ - 1
		b.p0.z -= offZ
		b.p1.z -= offZ
	}

	for _, b := range bricks {
		for _, s := range b.supports {
			s.supported[b.id] = b
		}
	}
}

func resetFlags(b *brick) {
	b.flag = false
	for _, s := range b.supported {
		resetFlags(s)
	}
}

func countFallingImpl(b *brick) int {
	b.flag = true
	sum := 1
	for _, s := range b.supported {
		supported := false
		for _, z := range s.supports {
			if !z.flag {
				supported = true
			}
		}
		if !supported {
			fmt.Println("  brick", s.id, "would fall")
			sum += countFallingImpl(s)
		}
	}
	return sum
}

func countFalling(b *brick) int {
	fmt.Println("Testing disintegrating brick", b.id)
	d := countFallingImpl(b) - 1
	resetFlags(b)
	return d
}

func Day22(filename string, part2 bool) int {

	lines, err := utils.ReadLines(filename)

	bricks := []*brick{}
	i := 0
	for _, txt := range lines {
		ps := strings.Split(txt, "~")
		ps = append(strings.Split(ps[0], ","), strings.Split(ps[1], ",")...)
		p := make([]int, len(ps))
		for i, s := range ps {
			p[i], err = strconv.Atoi(s)
			if err != nil {
				panic("Not a number")
			}
		}
		b := brick{
			p0:        v3{p[0], p[1], p[2]},
			p1:        v3{p[3], p[4], p[5]},
			id:        i,
			supports:  make(map[int]*brick),
			supported: make(map[int]*brick)}
		i++
		if b.p0.z > b.p1.z || b.p0.y > b.p1.y || b.p0.x > b.p1.x {
			panic("Inversed coordinates")
		}
		bricks = append(bricks, &b)
	}

	fmt.Println(len(bricks))

	minZ, maxZ := bricks[0].p0.z, bricks[0].p1.z
	minX, maxX := bricks[0].p0.x, bricks[0].p1.x
	minY, maxY := bricks[0].p0.y, bricks[0].p1.y
	for i := 0; i < len(bricks); i++ {
		minX = min(bricks[i].p0.x, minX)
		minY = min(bricks[i].p0.y, minY)
		minZ = min(bricks[i].p0.z, minZ)
		maxX = max(bricks[i].p1.x, maxX)
		maxY = max(bricks[i].p1.y, maxY)
		maxZ = max(bricks[i].p1.z, maxZ)
	}
	for i := 0; i < len(bricks); i++ {
		bricks[i].p0.x -= minX
		bricks[i].p1.x -= minX
		bricks[i].p0.y -= minY
		bricks[i].p1.y -= minY
		bricks[i].p0.z -= minZ
		bricks[i].p1.z -= minZ
	}
	maxX -= minX
	maxY -= minY
	maxZ -= minZ

	gravity(bricks)

	for _, b := range bricks {
		fmt.Printf("%s supported by ", makeStrId(b.id))
		for _, s := range b.supports {
			fmt.Printf("%s, ", makeStrId(s.id))
		}
		fmt.Printf("\n")
	}

	for _, b := range bricks {
		fmt.Printf("%s supports ", makeStrId(b.id))
		for _, s := range b.supported {
			fmt.Printf("%s, ", makeStrId(s.id))
		}
		fmt.Printf("\n")
	}

	sum := 0
	for _, b := range bricks {
		soleSupport := false
		for _, s := range b.supported {
			if len(s.supports) == 1 {
				soleSupport = true
			}
		}
		if !soleSupport {
			sum++
			fmt.Printf("%s can be disintegrated\n", makeStrId(b.id))
		}
	}
	fmt.Println(sum)
	if !part2 {
		return sum
	}

	sum = 0
	for _, b := range bricks {
		f := countFalling(b)
		fmt.Println("Disintegrating", b.id, "would make", f, "bricks fall")
		sum += f
	}
	fmt.Println(sum)
	return sum
}
