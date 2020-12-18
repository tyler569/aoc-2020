package main

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	noneToken byte = 0
	intToken  byte = 1
)

type token struct {
	typ byte
	int
}

func (t token) String() string {
	if t.typ == intToken {
		return fmt.Sprintf("#%v", t.int)
	} else {
		return fmt.Sprintf(".%c", t.typ)
	}
}

type tokenStream struct {
	tokens []token
}

func (ts *tokenStream) peek() (token, bool) {
	if len(ts.tokens) == 0 {
		return token{}, false
	}
	return ts.tokens[0], true
}

func (ts *tokenStream) pop() (token, bool) {
	if len(ts.tokens) == 0 {
		return token{}, false
	}
	tok := ts.tokens[0]
	ts.tokens = ts.tokens[1:]
	// fmt.Printf("pop: %v ", tok)
	return tok, true
}

var tokenRegex = regexp.MustCompile(`'.*'|".*"|[a-zA-Z_]\S*|\S`)

func tokenSplit(s string) []string {
	return tokenRegex.FindAllString(s, -1)
}

func tokenize(s string) *tokenStream {
	split := tokenSplit(s)
	tokens := []token{}
	var err error
	for _, s := range split {
		t := token{}
		t.int, err = strconv.Atoi(s)
		if err == nil {
			t.typ = intToken
		} else {
			t.typ = s[0]
		}
		tokens = append(tokens, t)
	}
	return &tokenStream{tokens}
}
