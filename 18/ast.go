package main

import (
	"fmt"
	"log"
)

type nodeType int

const (
	intNode nodeType = iota
	add
	sub
	mul
	div
)

type node struct {
	nodeType
	int
	left  *node
	right *node
}

func opConv(tokenOp byte) nodeType {
	switch tokenOp {
	case '+':
		return add
	case '-':
		return sub
	case '*':
		return mul
	case '/':
		return div
	default:
		log.Fatalf("expected operator")
		return 0
	}
}

func parseInt(tokens *tokenStream) *node {
	t, ok := tokens.pop()
	if !ok || t.typ != intToken {
		log.Fatalf("expected int, got %v\n", t)
	}

	return &node{intNode, t.int, nil, nil}
}

func parseParen(tokens *tokenStream) *node {
	if *verbose {
		fmt.Println("parseParen:", tokens.tokens)
	}
	t, ok := tokens.peek()
	if ok && t.typ == intToken {
		return parseInt(tokens)
	}
	if !ok || t.typ != '(' {
		log.Fatalf("expected paren, got ? (%v)\n", tokens.tokens)
	}

	tokens.pop()
	factor := parseBinop(tokens)

	t, ok = tokens.peek()
	if ok && t.typ == ')' {
		tokens.pop()
	}

	return factor
}

func parseBinop(tokens *tokenStream) *node {
	if *verbose {
		fmt.Println("parseSum:", tokens.tokens)
	}

	left := parseParen(tokens)
	op, ok := tokens.peek()
	if !ok || op.typ == ')' {
		return left
	}
	tokens.pop()
	right := parseParen(tokens)
	n := &node{
		left:     left,
		right:    right,
		nodeType: opConv(op.typ),
	}

	for {
		op, ok = tokens.peek()
		if !ok || op.typ == ')' {
			return n
		}
		tokens.pop()

		rhs := parseParen(tokens)
		n = &node{
			left:     n,
			right:    rhs,
			nodeType: opConv(op.typ),
		}
	}
}
