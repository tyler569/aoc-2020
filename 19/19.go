package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var file = flag.String("file", "input", "The file")

func main() {
	flag.Parse()

	content, err := ioutil.ReadFile(*file)
	if err != nil {
		fatalf("%v\n", err)
	}

	sp := strings.Split(string(content), "\n\n")
	allrules := sp[0]
	allexamples := sp[1]

	rules := strings.Split(allrules, "\n")
	examples := strings.Split(allexamples, "\n")

	ruleSet := parseRules(rules)
	for k, v := range ruleSet {
		fmt.Println(k, v)
	}

	total := 0
	for _, message := range examples {
		if match(message, ruleSet) {
			total++
		}
	}
	fmt.Println("P1:", total)

	total = 0
	for _, message := range examples {
		if match(message, ruleSet) {
			total++
		}
	}
	fmt.Println("P2:", total)
}

type rule struct {
	letter       byte
	alternatives [][]int
}

func parseRules(ruleLines []string) map[int]rule {
	ruleSet := map[int]rule{}

	for _, r := range ruleLines {
		sp := strings.Split(r, ": ")
		num := atoi(sp[0])
		ruleText := sp[1]
		if ruleText[0] == '"' {
			ruleSet[num] = rule{letter: ruleText[1]}
		} else {
			ruleSet[num] = rule{alternatives: parseAlternatives(ruleText)}
		}
	}
	return ruleSet
}

func parseAlternatives(ruleText string) [][]int {
	alternatives := strings.Split(ruleText, " | ")
	res := [][]int{}
	
	for _, a := range alternatives {
		values := strings.Split(a, " ")
		ivalues := []int{}
		for _, v := range values {
			ivalues = append(ivalues, atoi(v))
		}
		res = append(res, ivalues)
	}

	return res
}

func matchRule(s string, ruleSet map[int]rule, rule int) (bool, string) {
	// fmt.Printf("matchRule(%v, ruleSet, %v)\n", s, rule)
	if len(s) == 0 {
		return false, s
	}

	r := ruleSet[rule]
	if r.letter != 0 {
		if len(s) > 1 {
			return s[0] == r.letter, s[1:]
		} else {
			return s[0] == r.letter, ""
		}
	}

	for _, alternative := range r.alternatives {
		hit := true
		sCopy := s
		for _, ruleNum := range alternative {
			match, rest := matchRule(sCopy, ruleSet, ruleNum)
			if !match {
				hit = false
				break
			}
			sCopy = rest
		}
		if hit {
			return true, sCopy
		}
	}
	return false, s
}

func match(s string, ruleSet map[int]rule) bool {
	ok, rest := matchRule(s, ruleSet, 0)
	if len(rest) > 0 {
		return false
	}
	return ok
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		fatalf("Could not parse %v as an integer; %v\n", s, err)
	}
	return v
}

func fatalf(format string, values ...interface{}) {
	fmt.Fprintf(os.Stderr, format, values...)
	os.Exit(1)
}
