package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

func part1(adapters []int) int {

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
	sort.Ints(numbers)

	fmt.Println(numbers)
}
