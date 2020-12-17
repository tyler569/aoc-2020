package main

import "fmt"

type point struct {
	x, y, z int
}

func (p point) neighbors() []point {
	x := p.x
	y := p.y
	z := p.z

	return []point{
		point{x - 1, y - 1, z - 1},
		point{x - 1, y - 1, z},
		point{x - 1, y - 1, z + 1},
		point{x - 1, y, z - 1},
		point{x - 1, y, z},
		point{x - 1, y, z + 1},
		point{x - 1, y + 1, z - 1},
		point{x - 1, y + 1, z},
		point{x - 1, y + 1, z + 1},

		point{x, y - 1, z - 1},
		point{x, y - 1, z},
		point{x, y - 1, z + 1},
		point{x, y, z - 1},
		// point{x  , y  , z  },
		point{x, y, z + 1},
		point{x, y + 1, z - 1},
		point{x, y + 1, z},
		point{x, y + 1, z + 1},

		point{x + 1, y - 1, z - 1},
		point{x + 1, y - 1, z},
		point{x + 1, y - 1, z + 1},
		point{x + 1, y, z - 1},
		point{x + 1, y, z},
		point{x + 1, y, z + 1},
		point{x + 1, y + 1, z - 1},
		point{x + 1, y + 1, z},
		point{x + 1, y + 1, z + 1},
	}
}

type board map[point]bool

func (b board) iterate() board {
	res := board{}
	neighbors := map[point]int{}

	for p, v := range b {
		if v {
			for _, n := range p.neighbors() {
				neighbors[n] += 1
			}
		}
	}

	for p, count := range neighbors {
		if count == 2 || count == 3 {
			res[p] = b[p]
		}
		if count == 3 {
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
		b[point{x, y, 0}] = c == '#'
		x += 1
	}
	return b
}

type point4 struct {
	x, y, z, w int
}

func (p point4) neighbors() []point4 {
	res := []point4{}
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					res = append(res, point4{p.x+dx, p.y+dy, p.z+dz, p.w+dw})
				}
			}
		}
	}
	return res
}

type board4 map[point4]bool

func (b board4) iterate() board4 {
	res := board4{}
	neighbors := map[point4]int{}

	for p, v := range b {
		if v {
			for _, n := range p.neighbors() {
				neighbors[n] += 1
			}
		}
	}

	for p, count := range neighbors {
		if count == 2 || count == 3 {
			res[p] = b[p]
		}
		if count == 3 {
			res[p] = true
		}
	}

	return res
}

func (b board4) count() int {
	count := 0
	for _, v := range b {
		if v {
			count++
		}
	}
	return count
}

func parse4(s string) board4 {
	b := board4{}
	x := 0
	y := 0
	for _, c := range s {
		if c == '\n' {
			x = 0
			y += 1
			continue
		}
		b[point4{x, y, 0, 0}] = c == '#'
		x += 1
	}
	return b
}

const example = `
.#.
..#
###
`

const p1 = `
.#.#..##
..#....#
##.####.
...####.
#.##..##
#...##..
...##.##
#...#.#.
`

func main() {
	b := parse(example)
	for i := 0; i < 6; i++ {
		b = b.iterate()
	}
	fmt.Println("example:", b.count())

	b = parse(p1)
	for i := 0; i < 6; i++ {
		b = b.iterate()
	}
	fmt.Println("P1:", b.count())

	b4 := parse4(example)
	for i := 0; i < 6; i++ {
		b4 = b4.iterate()
	}
	fmt.Println("example4:", b4.count())

	b4 = parse4(p1)
	for i := 0; i < 6; i++ {
		b4 = b4.iterate()
	}
	fmt.Println("P2:", b4.count())
}
