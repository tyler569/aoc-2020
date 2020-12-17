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

func parse(w io.Reader) (numbers []int, err error) {
	sc := bufio.NewScanner(w)

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

func part1(numbers []int) (int, error) {
	for i, a := range numbers {
		for _, b := range numbers[i+1:] {
			if a+b == 2020 {
				return a * b, nil
			}
		}
	}
	return 0, errors.New("No answer")
}

func part2(numbers []int) (int, error) {
	for i, a := range numbers {
		for j, b := range numbers[i+1:] {
			for _, c := range numbers[j+i:] {
				if a+b+c == 2020 {
					return a * b * c, nil
				}
			}
		}
	}
	return 0, errors.New("No answer")
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

	answer, err := part1(numbers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("P1:", answer)

	answer, err = part2(numbers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("P2:", answer)
}
