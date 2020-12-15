package main

import "fmt"

const max = 30_000_000

func main() {
	last := map[int]int{}

	last[0] = 1
	last[6] = 2
	last[1] = 3
	last[7] = 4
	last[2] = 5
	last[19] = 6

	prev := 20

	for i := 7; i < max; i++ {
		next := 0

		if l, ok := last[prev]; ok {
			next = i - l;
		}

		last[prev] = i
		prev = next
		// fmt.Printf("%6v %v\n", prev, last)
		if i < 11 || i > max-5 {
			fmt.Printf("%4v %6v\n", i+1, prev)
		}
	}
}
