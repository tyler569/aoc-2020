package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func parse(input io.Reader) (numbers []int, err error) {
	sc := bufio.NewScanner(input)

	for sc.Scan() {
		str := sc.Text()
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, n)
	}
	return
}

func numbersAdd(number int, window []int) bool {
	for i, x := range window {
		for _, y := range window[i+1:] {
			if x+y == number {
				return true
			}
		}
	}
	return false
}

func part1(numbers []int) (int, error) {
	for i := 25; i < len(numbers)-25; i++ {
		window := numbers[i-25 : i]
		number := numbers[i]

		if !numbersAdd(number, window) {
			return number, nil
		}
	}
	return 0, errors.New("No answer")
}

func addRange(numbers []int, answer int) ([]int, error) {
	for i := range numbers {
		sum := 0
		rng := numbers[i:]
		j := 0
		for sum < answer {
			sum += rng[j]
			j += 1
		}
		if sum == answer {
			return rng[:j], nil
		}
	}
	return nil, errors.New("No answer")
}

func linearRange(numbers []int, answer int) ([]int, error) {
	var a, b int
	b = 1
	sum := numbers[0]
	for b < len(numbers) {
		sum += numbers[b]
		b += 1

		for sum > answer {
			sum -= numbers[a]
			a += 1
		}

		if sum == answer {
			return numbers[a:b], nil
		}
	}
	return nil, errors.New("No answer")
}

func part2(window []int) int {
	min := window[0]
	max := window[0]
	for _, n := range window {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return max + min
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	numbers, err := parse(file)
	if err != nil {
		log.Fatal(err)
	}

	p1answer, err := part1(numbers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("P1:", p1answer)

	rng, err := addRange(numbers, p1answer)
	if err != nil {
		log.Fatal(err)
	}

	answer := part2(rng)
	fmt.Println("P2:", answer)

	rng, err = linearRange(numbers, p1answer)
	if err != nil {
		log.Fatal(err)
	}

	answer = part2(rng)
	fmt.Println("L2:", answer)
}
