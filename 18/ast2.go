package main

import (
	"fmt"
	"log"
)

func parse2Int(tokens *tokenStream) *node {
	t, ok := tokens.pop()
	if !ok || t.typ != intToken {
		log.Fatalf("expected int, got %v\n", t)
	}

	return &node{intNode, t.int, nil, nil}
}

func parse2Paren(tokens *tokenStream) *node {
	if *verbose {
		fmt.Println("parse2Paren:", tokens.tokens)
	}

	t, ok := tokens.peek()
	if ok && t.typ == intToken {
		return parse2Int(tokens)
	}
	if !ok || t.typ != '(' {
		log.Fatalf("expected paren, got ? (%v)\n", tokens.tokens)
	}

	tokens.pop()
	factor := parse2Factor(tokens)

	t, ok = tokens.peek()
	if ok && t.typ == ')' {
		tokens.pop()
	}

	return factor
}

func parse2Sum(tokens *tokenStream) *node {
	if *verbose {
		fmt.Println("parse2Sum:", tokens.tokens)
	}

	left := parse2Paren(tokens)
	op, ok := tokens.peek()
	if !ok || !(op.typ == '+' || op.typ == '-') {
		return left
	}
	tokens.pop()
	right := parse2Paren(tokens)
	n := &node{
		left:     left,
		right:    right,
		nodeType: opConv(op.typ),
	}

	for {
		op, ok = tokens.peek()
		if !ok || !(op.typ == '+' || op.typ == '-') {
			return n
		}
		tokens.pop()

		rhs := parse2Paren(tokens)
		n = &node{
			left:     n,
			right:    rhs,
			nodeType: opConv(op.typ),
		}
	}
}

func parse2Factor(tokens *tokenStream) *node {
	if *verbose {
		fmt.Println("parse2Factor:", tokens.tokens)
	}

	left := parse2Sum(tokens)
	op, ok := tokens.peek()
	if !ok || !(op.typ == '*' || op.typ == '/') {
		return left
	}
	tokens.pop()
	right := parse2Sum(tokens)
	n := &node{
		left:     left,
		right:    right,
		nodeType: opConv(op.typ),
	}

	for {
		op, ok = tokens.peek()
		if !ok || !(op.typ == '*' || op.typ == '/') {
			return n
		}
		tokens.pop()

		rhs := parse2Sum(tokens)
		n = &node{
			left:     n,
			right:    rhs,
			nodeType: opConv(op.typ),
		}
	}
}
