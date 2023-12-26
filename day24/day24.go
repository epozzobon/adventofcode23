package day24

import (
	"fmt"
	"math"
	"regexp"
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

type stone struct {
	position v3
	velocity v3
}

var re = regexp.MustCompile(`^(-?\d+), +(-?\d+), +(-?\d+) +@ +(-?\d+), +(-?\d+), +(-?\d+)$`)

func str2stone(txt string) stone {
	strings.Split(txt, "")
	matches := re.FindStringSubmatch(txt)
	ints := utils.Map(matches[1:], utils.Atoi)
	p := v3{ints[0], ints[1], ints[2]}
	v := v3{ints[3], ints[4], ints[5]}
	return stone{p, v}
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func intersection2D(s1, s2 stone) v3f {
	vx1 := s1.velocity.x
	vx2 := s2.velocity.x
	px1 := s1.position.x
	px2 := s2.position.x
	vy1 := s1.velocity.y
	vy2 := s2.velocity.y
	py1 := s1.position.y
	py2 := s2.position.y

	n := px2*vy1 - px1*vy1 - py2*vx1 + py1*vx1
	d := vy2*vx1 - vx2*vy1
	if d == 0 {
		return v3f{math.Inf(1), math.Inf(1), math.Inf(1)}
	}
	t2 := float64(n) / float64(d)
	x := float64(s2.position.x) + float64(s2.velocity.x)*t2
	y := float64(s2.position.y) + float64(s2.velocity.y)*t2
	t1 := 0.0
	if vy1 != 0 {
		t1 = (y - float64(py1)) / float64(vy1)
	} else if vx1 != 0 {
		t1 = (x - float64(px1)) / float64(vx1)
	} else {
		panic("how")
	}
	if t1 < 0 || t2 < 0 {
		return v3f{-math.Inf(1), -math.Inf(1), -math.Inf(1)}
	}
	return v3f{x, y, 0.0}
}

type limits struct {
	min float64
	max float64
}

func Day24p1(filename string, lim limits) int {

	file, err := utils.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	stones := utils.Map(file, str2stone)
	count := 0
	for i := 0; i < len(stones)-1; i++ {
		for j := i + 1; j < len(stones); j++ {
			intersection := intersection2D(stones[i], stones[j])
			if intersection.x >= lim.min && intersection.x <= lim.max {
				if intersection.y >= lim.min && intersection.y <= lim.max {
					count++
				}
			}
			fmt.Println(stones[i], stones[j], intersection)
		}
	}

	return count
}

func solveMatrix(matrix [][]int) []int {

	m := utils.Map2D(matrix, func(a int) float64 { return float64(a) })
	T := len(m) - 1

	for j := 0; j <= T; j++ {
		d := m[j][j]
		if math.Abs(d) <= 1e-4 {
			panic("zero on diagonal is bad, pivoting is not implemented")
		}

		for i := j + 1; i < len(m); i++ {
			a := m[i][j] / d
			// subtract row j times a from row i
			for k := 0; k < len(m[i]); k++ {
				m[i][k] -= m[j][k] * a
			}

			if math.Abs(m[i][j]) > 1e-4 {
				panic("row subtraction failed")
			}
			m[i][j] = 0
		}

		for k := 0; k < len(m[j]); k++ {
			m[j][k] /= d
		}
		fmt.Println(m[j])
	}

	for i := T; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			// subtract row i from row j
			m[j][T+1] -= m[i][T+1] * m[j][i]
			m[j][i] = 0
		}
	}

	solutions := utils.Map(m, func(a []float64) int {
		return int(math.Round(a[T+1]))
	})
	fmt.Println(solutions)
	return solutions
}

func Day24p2(filename string) int {

	file, err := utils.ReadLines(filename)
	if err != nil {
		panic(err)
	}

	stones := utils.Map(file, str2stone)
	px0, vx0 := stones[0].position.x, stones[0].velocity.x
	py0, vy0 := stones[0].position.y, stones[0].velocity.y
	pz0, vz0 := stones[0].position.z, stones[0].velocity.z
	px1, vx1 := stones[1].position.x, stones[1].velocity.x
	py1, vy1 := stones[1].position.y, stones[1].velocity.y
	pz1, vz1 := stones[1].position.z, stones[1].velocity.z
	px2, vx2 := stones[2].position.x, stones[2].velocity.x
	py2, vy2 := stones[2].position.y, stones[2].velocity.y
	pz2, vz2 := stones[2].position.z, stones[2].velocity.z
	px3, vx3 := stones[3].position.x, stones[3].velocity.x
	py3, vy3 := stones[3].position.y, stones[3].velocity.y
	pz3, vz3 := stones[3].position.z, stones[3].velocity.z
	px4, vx4 := stones[4].position.x, stones[4].velocity.x
	py4, vy4 := stones[4].position.y, stones[4].velocity.y
	pz4, vz4 := stones[4].position.z, stones[4].velocity.z

	fmt.Println("px0=", px0, "  vx0=", vx0)
	fmt.Println("py0=", py0, "  vy0=", vy0)
	fmt.Println("pz0=", pz0, "  vz0=", vz0)
	fmt.Println("px1=", px1, "  vx1=", vx1)
	fmt.Println("py1=", py1, "  vy1=", vy1)
	fmt.Println("pz1=", pz1, "  vz1=", vz1)
	fmt.Println("px2=", px2, "  vx2=", vx2)
	fmt.Println("py2=", py2, "  vy2=", vy2)
	fmt.Println("pz2=", pz2, "  vz2=", vz2)
	fmt.Println("px3=", px3, "  vx3=", vx3)
	fmt.Println("py3=", py3, "  vy3=", vy3)
	fmt.Println("pz3=", pz3, "  vz3=", vz3)
	fmt.Println("px4=", px4, "  vx4=", vx4)
	fmt.Println("py4=", py4, "  vy4=", vy4)
	fmt.Println("pz4=", pz4, "  vz4=", vz4)

	fmt.Printf("%5d * pxk + %18d * vxk + %5d * pyk + %18d * vyk + w = %18d\n", vy0, -py0, -vx0, px0, px0*vy0-py0*vx0)
	fmt.Printf("%5d * pxk + %18d * vxk + %5d * pyk + %18d * vyk + w = %18d\n", vy1, -py1, -vx1, px1, px1*vy1-py1*vx1)
	fmt.Printf("%5d * pxk + %18d * vxk + %5d * pyk + %18d * vyk + w = %18d\n", vy2, -py2, -vx2, px2, px2*vy2-py2*vx2)
	fmt.Printf("%5d * pxk + %18d * vxk + %5d * pyk + %18d * vyk + w = %18d\n", vy3, -py3, -vx3, px3, px3*vy3-py3*vx3)
	fmt.Printf("%5d * pxk + %18d * vxk + %5d * pyk + %18d * vyk + w = %18d\n", vy4, -py4, -vx4, px4, px4*vy4-py4*vx4)

	toSolve := [][]int{
		{vy0, -py0, -vx0, px0, 1, px0*vy0 - py0*vx0},
		{vy1, -py1, -vx1, px1, 1, px1*vy1 - py1*vx1},
		{vy2, -py2, -vx2, px2, 1, px2*vy2 - py2*vx2},
		{vy3, -py3, -vx3, px3, 1, px3*vy3 - py3*vx3},
		{vy4, -py4, -vx4, px4, 1, px4*vy4 - py4*vx4},
	}

	solutions := solveMatrix(toSolve)

	pxk := solutions[0]
	vxk := solutions[1]
	pyk := solutions[2]
	vyk := solutions[3]

	fmt.Printf("%18d * vzk + %5d * pzk + w = %18d\n", px0, vx0, pz0*vx0-px0*vz0+vz0*pxk-pz0*vxk)
	fmt.Printf("%18d * vzk + %5d * pzk + w = %18d\n", px1, vx1, pz1*vx1-px1*vz1+vz1*pxk-pz1*vxk)
	fmt.Printf("%18d * vzk + %5d * pzk + w = %18d\n", px2, vx2, pz2*vx2-px2*vz2+vz2*pxk-pz2*vxk)

	toSolve = [][]int{
		{vx0, px0, 1, pz0*vx0 - px0*vz0 + vz0*pxk - pz0*vxk},
		{vx1, px1, 1, pz1*vx1 - px1*vz1 + vz1*pxk - pz1*vxk},
		{vx2, px2, 1, pz2*vx2 - px2*vz2 + vz2*pxk - pz2*vxk},
	}

	solutions = solveMatrix(toSolve)

	pzk := solutions[0]
	vzk := solutions[1]

	fmt.Println("rock:", pxk, pyk, pzk, "@", vxk, vyk, vzk)
	return pxk + pyk + pzk
}
