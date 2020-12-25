package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var input = flag.String("input", "input", "Input file")

func readInput(file string) (rules []string, examples []string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fatalf("%v\n", err)
	}

	sp := strings.Split(string(content), "\n\n")
	allrules := sp[0]
	allexamples := sp[1]

	rules = strings.Split(allrules, "\n")
	examples = strings.Split(allexamples, "\n")
	return
}

func main() {
	flag.Parse()
	rules, examples := readInput(*input)
	ruleSet := parseRules(rules)

	total := 0
	for _, message := range examples {
		if match(message, ruleSet) {
			total++
		}
	}
	fmt.Println("P_:", total)
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

func matchRule(s string, ruleSet map[int]rule, rule int) (bool, []string) {
	// fmt.Printf("matchRule(%v, ruleSet, %v)\n", s, rule)
	var allRests []string

	if len(s) == 0 {
		allRests = append(allRests, "")
		return false, allRests
	}

	r := ruleSet[rule]
	if r.letter != 0 {
		if len(s) > 1 {
			allRests = append(allRests, s[1:])
			return s[0] == r.letter, allRests
		} else {
			allRests = append(allRests, "")
			return s[0] == r.letter, allRests
		}
	}

	anyMatch := false
	for _, alternative := range r.alternatives {
		// altRests := []string{}
		ruleRestAcc := []string{s}
		altMatch := true
		for _, ruleNum := range alternative {
			ruleRestThis := []string{}
			anySubMatch := false
			for _, rr := range ruleRestAcc {
				match, rs := matchRule(rr, ruleSet, ruleNum)
				if match {
					anySubMatch = true
					ruleRestThis = append(ruleRestThis, rs...)
				}
			}
			if !anySubMatch {
				altMatch = false
				break
			}
			ruleRestAcc = ruleRestThis
		}
		if altMatch {
			allRests = append(allRests, ruleRestAcc...)
			anyMatch = true
		}
	}

	fmt.Printf("matchRule(%v, ...) -> %v %v\n", s, anyMatch, allRests)
	return anyMatch, allRests
}

func match(s string, ruleSet map[int]rule) bool {
	ok, rest := matchRule(s, ruleSet, 0)
	exact := false
	for _, r := range rest {
		if len(r) == 0 {
			exact = true
		}
	}
	if len(rest) == 0 {
		exact = true
	}
	return ok && exact
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
