package main

import (
	"fmt"
	"time"
)

func cache(fn func(int) int) func(int) int {
	c := map[int]int{}
	return func(a int) int {
		v, ok := c[a]
		if ok {
			return v
		}
		v = fn(a)
		c[a] = v
		return v
	}
}

var fibonacci = cache(func(n int) int {
	if n <= 0 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
})

func main() {
	const n = 42

	fmt.Print(".")
	t0 := time.Now()
	a := fibonacci(n)
	fmt.Print(".")
	t2 := time.Now()

	fmt.Println(a)
	fmt.Println(t2.Sub(t0))
}
