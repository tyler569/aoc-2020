package main

import (
	"flag"
	"fmt"
)

type point struct {
	x, y, z, w int
}

type board map[point]bool

func (b board) iterate(live uint, born uint) board {
	res := board{}
	neighbors := map[point]int{}

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					for p, v := range b {
						pPrime := point{p.x + dx, p.y + dy, p.z + dz, p.w + dw}
						if v {
							neighbors[pPrime] += 1
						}
					}
				}
			}
		}
	}

	for p, count := range neighbors {
		if count > 64 {
			continue
		}
		if 1<<count&live != 0 && b[p] {
			res[p] = true
		}
		if 1<<count&born != 0 {
			res[p] = true
		}
	}

	return res
}

func (b board) count() int {
	count := 0
	for _, v := range b {
		if v {
			count++
		}
	}
	return count
}

func parse(s string) board {
	b := board{}
	x := 0
	y := 0
	for _, c := range s {
		if c == '\n' {
			x = 0
			y += 1
			continue
		}
		b[point{x, y, 0, 0}] = c == '#'
		x += 1
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (b board) print() {
	var firstPoint point
	for p := range b {
		firstPoint = p
		break
	}

	xmin := firstPoint.x
	xmax := firstPoint.x
	ymin := firstPoint.y
	ymax := firstPoint.y
	zmin := firstPoint.z
	zmax := firstPoint.z
	wmin := firstPoint.w
	wmax := firstPoint.w
	// var xmin, xmax, ymin, ymax, zmin, zmax, wmin, wmax int

	for p := range b {
		xmax = max(p.x, xmax)
		ymax = max(p.y, ymax)
		zmax = max(p.z, zmax)
		wmax = max(p.w, wmax)
		xmin = min(p.x, xmin)
		ymin = min(p.y, ymin)
		zmin = min(p.z, zmin)
		wmin = min(p.w, wmin)
	}

	xmin--
	ymin--

	xmax++
	ymax++

	// w columns
	// z rows
	// xy squares

	//       w=-1 w=0  w=1
	//       xxxx yyyy zzzz
	// z = -1xxxx yyyy zzzz
	//       xxxx yyyy zzzz
	//       xxxx yyyy zzzz

	//       xxxx yyyy zzzz
	// z = 0 xxxx yyyy zzzz
	//       xxxx yyyy zzzz
	//       xxxx yyyy zzzz

	// for z {
	// 	for y {
	// 		for w {
	// 			for x {
	// 			}
	// 			print ' '
	// 		}
	// 		print '\n'
	// 	}
	// 	print '\n'
	// }

	for pz := zmin; pz <= zmax; pz++ {
		for py := ymin; py <= ymax; py++ {
			for pw := wmin; pw <= wmax; pw++ {
				for px := xmin; px <= xmax; px++ {
					p := point{px, py, pz, pw}
					if b[p] {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Print(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

const stable4_5 = `
.##.#
..###
#####
`

const example = `
.#.
..#
###
`

func r(counts ...int) (o uint) {
	for _, c := range counts {
		o |= 1 << c
	}
	return
}

var iters = flag.Int("iters", 6, "Iterations to run")

func main() {
	flag.Parse()
	// b := parse(example)
	// for i := 0; i < 6; i++ {
	// 	// fmt.Printf("%3v: %v\n", i, b.count())
	// 	b.print()
	// 	fmt.Println("---")
	// 	b = b.iterate(r(2), r(3))
	// }

	b := parse(example)
	for i := 0; i <= *iters; i++ {
		b.print()
		fmt.Printf("%3v: %v\n", i, b.count())
		fmt.Println("---")
		b = b.iterate(r(2), r(3))
	}
}
