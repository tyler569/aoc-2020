package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func eval(n *node) int {
	if n.nodeType == intNode {
		return n.int
	}
	if n.left == n {
		log.Fatal("cycle detected")
	}
	l := eval(n.left)
	r := eval(n.right)
	switch n.nodeType {
	case add:
		if *verbose {
			fmt.Println(l, "+", r)
		}
		return l + r
	case sub:
		if *verbose {
			fmt.Println(l, "-", r)
		}
		return l - r
	case mul:
		if *verbose {
			fmt.Println(l, "*", r)
		}
		return l * r
	case div:
		if *verbose {
			fmt.Println(l, "/", r)
		}
		return l / r
	}
	log.Fatal("invalid eval")
	return 0
}

var file = flag.String("file", "input", "Input file")
var verbose = flag.Bool("verbose", false, "Be verbose")

func main() {
	flag.Parse()

	bcontent, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	content := string(bcontent)
	lines := strings.Split(content, "\n")
	sum := 0

	for _, line := range lines {
		tokens := tokenize(line)
		if len(tokens.tokens) == 0 {
			continue
		}
		if *verbose {
			fmt.Println(tokens.tokens)
		}
		ast := parseBinop(tokens)
		// fmt.Println(ast)
		result := eval(ast)
		if *verbose {
			fmt.Println("result:", result)
		}
		sum += result
	}

	fmt.Println("P1:", sum)

	sum = 0
	for _, line := range lines {
		tokens := tokenize(line)
		if len(tokens.tokens) == 0 {
			continue
		}
		if *verbose {
			fmt.Println(tokens.tokens)
		}
		ast := parse2Factor(tokens)
		// fmt.Println(ast)
		result := eval(ast)
		if *verbose {
			fmt.Println("result:", result)
		}
		sum += result
	}

	fmt.Println("P2:", sum)
}
