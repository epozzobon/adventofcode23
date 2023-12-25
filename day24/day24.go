package main

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

	// x = px1 + t1 * vx1 = px2 + t2 * vx2
	// y = py1 + t1 * vy1 = py2 + t2 * vy2
	//
	// t1 = (px2 + t2 * vx2 - px1) / vx1
	// t1 = (py2 + t2 * vy2 - py1) / vy1
	//
	// (px2 + t2 * vx2 - px1) / vx1 = (py2 + t2 * vy2 - py1) / vy1
	// (px2 + t2 * vx2 - px1) * vy1 = (py2 + t2 * vy2 - py1) * vx1
	// px2 * vy1 + t2 * vx2 * vy1 - px1 * vy1 = py2 * vx1 + t2 * vy2 * vx1 - py1 * vx1
	// px2 * vy1 - px1 * vy1 - py2 * vx1 + py1 * vx1 = t2 * vy2 * vx1 - t2 * vx2 * vy1
	// px2 * vy1 - px1 * vy1 - py2 * vx1 + py1 * vx1 = t2 * (vy2 * vx1 - vx2 * vy1)
	//
	// t2 = (px2 * vy1 - px1 * vy1 - py2 * vx1 + py1 * vx1) / (vy2 * vx1 - vx2 * vy1)
	// x = px2 + t2 * vx2
	// y = py2 + t2 * vy2
	// t1 = (x - px1) / vx1
	// t1 = (y - py1) / vy1

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
	subrow := func(dst []float64, scaling float64, src []float64) {
		for i := 0; i < len(dst); i++ {
			dst[i] = dst[i] - src[i]*scaling
		}
	}

	var solve func(m [][]float64) []float64
	solve = func(m [][]float64) []float64 {
		for j := 0; j < len(m[0])-1; j++ {
			for i := j + 1; i < len(m); i++ {
				a := m[i][j] / m[j][j]
				// subtract row j times a from row i
				subrow(m[i], a, m[j])
				if math.Abs(m[i][j]) > 1e-4 {
					panic("subrow failed")
				}
				m[i][j] = 0
			}
		}

		T := len(m) - 1
		solutions := make([]float64, len(m))
		for i := T; i >= 0; i-- {
			solutions[i] = m[i][T+1]
			for j := i + 1; j <= T; j++ {
				solutions[i] -= m[i][j] * solutions[j]
			}
			solutions[i] /= m[i][i]
		}
		fmt.Println(solutions)

		return solutions
	}
	result := solve(utils.Map2D(matrix, func(a int) float64 {
		return float64(a)
	}))
	return utils.Map(result, func(a float64) int {
		return int(math.Round(a))
	})
}

func Day24p2(filename string, arg int) int {

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

	// px0 + vx0 * t0 = pxk + vxk * t0
	// py0 + vy0 * t0 = pyk + vyk * t0
	// pz0 + vz0 * t0 = pzk + vzk * t0
	// px1 + vx1 * t1 = pxk + vxk * t1
	// py1 + vy1 * t1 = pyk + vyk * t1
	// pz1 + vz1 * t1 = pzk + vzk * t1
	// px2 + vx2 * t2 = pxk + vxk * t2
	// py2 + vy2 * t2 = pyk + vyk * t2
	// pz2 + vz2 * t2 = pzk + vzk * t2
	//
	// (pxk - px0) / (vx0 - vxk) = t0
	// (pyk - py0) / (vy0 - vyk) = t0
	// (pzk - pz0) / (vz0 - vzk) = t0
	//
	// (pxk - px0) * (vy0 - vyk) = (pyk - py0) * (vx0 - vxk)
	// pyk * vxk - pxk * vyk = px0 * vyk - pxk * vy0 + px0 * vy0 + pyk * vx0 - py0 * vx0 - py0 * vxk
	//
	//
	// pyk * vxk - pxk * vyk = px0 * vyk - pxk * vy0 + px0 * vy0 + pyk * vx0 - py0 * vx0 - py0 * vxk
	// pyk * vxk - pxk * vyk = px1 * vyk - pxk * vy1 + px1 * vy1 + pyk * vx1 - py1 * vx1 - py1 * vxk
	//
	//
	// px0 * vyk - pxk * vy0 + px0 * vy0 + pyk * vx0 - py0 * vx0 - py0 * vxk = |
	// px1 * vyk - pxk * vy1 + px1 * vy1 + pyk * vx1 - py1 * vx1 - py1 * vxk = |
	// px2 * vyk - pxk * vy2 + px2 * vy2 + pyk * vx2 - py2 * vx2 - py2 * vxk = |
	// px3 * vyk - pxk * vy3 + px3 * vy3 + pyk * vx3 - py3 * vx3 - py3 * vxk = |
	//
	//
	// px0 * vyk - vy0 * pxk + vx0 * pyk - py0 * vxk + px0 * vy0 - py0 * vx0 = -w
	// px1 * vyk - vy1 * pxk + vx1 * pyk - py1 * vxk + px1 * vy1 - py1 * vx1 = -w
	// px2 * vyk - vy2 * pxk + vx2 * pyk - py2 * vxk + px2 * vy2 - py2 * vx2 = -w
	// px3 * vyk - vy3 * pxk + vx3 * pyk - py3 * vxk + px3 * vy3 - py3 * vx3 = -w
	// px4 * vyk - vy4 * pxk + vx4 * pyk - py4 * vxk + px4 * vy4 - py4 * vx4 = -w
	//
	// px0 * vyk - vy0 * pxk + vx0 * pyk - py0 * vxk + w = py0 * vx0 - px0 * vy0
	// px1 * vyk - vy1 * pxk + vx1 * pyk - py1 * vxk + w = py1 * vx1 - px1 * vy1
	// px2 * vyk - vy2 * pxk + vx2 * pyk - py2 * vxk + w = py2 * vx2 - px2 * vy2
	// px3 * vyk - vy3 * pxk + vx3 * pyk - py3 * vxk + w = py3 * vx3 - px3 * vy3
	// px4 * vyk - vy4 * pxk + vx4 * pyk - py4 * vxk + w = py4 * vx4 - px4 * vy4

	fmt.Printf("%5d * vyk + %5d * pxk + %5d * pyk + %5d * vxk + w = %5d\n", px0, -vy0, vx0, -py0, py0*vx0-px0*vy0)
	fmt.Printf("%5d * vyk + %5d * pxk + %5d * pyk + %5d * vxk + w = %5d\n", px1, -vy1, vx1, -py1, py1*vx1-px1*vy1)
	fmt.Printf("%5d * vyk + %5d * pxk + %5d * pyk + %5d * vxk + w = %5d\n", px2, -vy2, vx2, -py2, py2*vx2-px2*vy2)
	fmt.Printf("%5d * vyk + %5d * pxk + %5d * pyk + %5d * vxk + w = %5d\n", px3, -vy3, vx3, -py3, py3*vx3-px3*vy3)
	fmt.Printf("%5d * vyk + %5d * pxk + %5d * pyk + %5d * vxk + w = %5d\n", px4, -vy4, vx4, -py4, py4*vx4-px4*vy4)

	toSolve := [][]int{
		{px0, -vy0, vx0, -py0, 1, py0*vx0 - px0*vy0},
		{px1, -vy1, vx1, -py1, 1, py1*vx1 - px1*vy1},
		{px2, -vy2, vx2, -py2, 1, py2*vx2 - px2*vy2},
		{px3, -vy3, vx3, -py3, 1, py3*vx3 - px3*vy3},
		{px4, -vy4, vx4, -py4, 1, py4*vx4 - px4*vy4},
	}

	solutions := solveMatrix(toSolve)

	vyk := solutions[0]
	pxk := solutions[1]
	pyk := solutions[2]
	vxk := solutions[3]

	// px0 * vzk + vx0 * pzk + w = pz0 * vx0 - px0 * vz0 + vz0 * pxk + pz0 * vxk
	// px1 * vzk + vx1 * pzk + w = pz1 * vx1 - px1 * vz1 + vz1 * pxk + pz1 * vxk
	// px2 * vzk + vx2 * pzk + w = pz2 * vx2 - px2 * vz2 + vz2 * pxk + pz2 * vxk

	fmt.Printf("%5d * vzk + %5d * pzk + w = %5d\n", px0, vx0, pz0*vx0-px0*vz0+vz0*pxk+pz0*vxk)
	fmt.Printf("%5d * vzk + %5d * pzk + w = %5d\n", px1, vx1, pz1*vx1-px1*vz1+vz1*pxk+pz1*vxk)
	fmt.Printf("%5d * vzk + %5d * pzk + w = %5d\n", px2, vx2, pz2*vx2-px2*vz2+vz2*pxk+pz2*vxk)

	toSolve = [][]int{
		{px0, vx0, 1, pz0*vx0 - px0*vz0 + vz0*pxk + pz0*vxk},
		{px1, vx1, 1, pz1*vx1 - px1*vz1 + vz1*pxk + pz1*vxk},
		{px2, vx2, 1, pz2*vx2 - px2*vz2 + vz2*pxk + pz2*vxk},
	}

	solutions = solveMatrix(toSolve)

	vzk := solutions[0]
	pzk := solutions[1]

	equal := true
	//for _, s := range stones {
	//	px0, vx0 := s.position.x, s.velocity.x
	//	py0, vy0 := s.position.y, s.velocity.y
	//	pz0, vz0 := s.position.z, s.velocity.z
	//	// px0 + vx0 * t0 = pxk + vxk * t0
	//	// px0 - pxk = (vxk - vx0) * t0
	//	tCollX := (px0 - pxk) / (vxk - vx0)
	//	tCollY := (py0 - pyk) / (vyk - vy0)
	//	tCollZ := (pz0 - pzk) / (vzk - vz0)
	//	if tCollX != tCollY || tCollX != tCollZ {
	//		panic("uh")
	//	}
	//}
	if equal {
		fmt.Println("GOOD:", pxk, pyk, pzk, "@", vxk, vyk, vzk)
		return pxk + pyk + pzk
	} else {
		panic("panik! math is wrong")
	}
}

func main() {
	utils.CheckSolution(Day24p1, "example1.txt", limits{7, 27}, 2)
	utils.CheckSolution(Day24p1, "input.txt", limits{200000000000000, 400000000000000}, 24192)
	utils.CheckSolution(Day24p2, "example1.txt", 0, 47)
	utils.CheckSolution(Day24p2, "input.txt", 0, 664822352550558)
}
