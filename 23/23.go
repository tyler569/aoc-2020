package main

import (
	"flag"
	"fmt"
)

var input = flag.String("input", "389125467", "Input")

type circle struct {
	cups []int
	maxc int
}

func (c circle) iterate() circle {
	// take 3
	current := c.cups[0]
	three := c.cups[1:4]

	// find destination
	destination := current - 1
	if destination < 1 {
		destination = c.maxc
	}

	for in(destination, three) {
		destination--
		if destination < 1 {
			destination = c.maxc
		}
	}

	// move 3 to after destination
	before, after := partition(c.cups[4:], destination)

	result := []int{}
	result = append(result, before...)
	result = append(result, three...)
	result = append(result, after...)
	result = append(result, current)

	return circle{result, c.maxc}
}

func main() {
	flag.Parse()

	cups := []int{}
	for _, c := range *input {
		cups = append(cups, int(c-'0'))
	}
	circl := circle{cups, max(cups)}
	fmt.Printf("%+v\n", circl)

	for i := 1; i <= 100; i++ {
		circl = circl.iterate()
	}
	fmt.Printf("P1: %+v\n", circl)
}

func in(i int, a []int) bool {
	for _, v := range a {
		if v == i {
			return true
		}
	}
	return false
}

// partition a slice into 2 sections, one containing everything up to
// and including the first `n` in the slice, the other containing
// the rest. If n is not in the slice, the second argument will be nil
func partition(s []int, n int) ([]int, []int) {
	for i := range s {
		if s[i] == n {
			return s[:i+1], s[i+1:]
		}
	}
	return s, nil
}

func max(s []int) int {
	maximum := s[0]
	for _, v := range s {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
}
