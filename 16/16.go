package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input-transformed")
	if err != nil {
		log.Fatal(err)
	}

	rules, tickets := parse(string(content))

	myTicket := tickets[0]

	// for rname, rule := range rules {
	// 	fmt.Println(rname, ":", rule)
	// }
	// for _, ticket := range tickets {
	// 	fmt.Println(ticket)
	// }

	sumInvalid(tickets, rules)

	validTickets := []ticket{}
	for _, t := range tickets {
		if ticketValid(t, rules) {
			validTickets = append(validTickets, t)
		}
	}

	fmt.Println("tickets:", len(tickets), len(validTickets));

	ruleColumns := map[string][]int{}

	for rname, rule := range rules {
		for c := 0; c < len(validTickets[0]); c++ {
			found := true
			for _, ticket := range validTickets {
				if !rule.contains(ticket[c]) {
					found = false
					break
				}
			}
			if found {
				ruleColumns[rname] = append(ruleColumns[rname], c)
			}
		}
	}

	hitCols := map[int]bool{}

	for {
		var hitCol int = -1
		any := false
		for r, c := range ruleColumns {
			fmt.Println(r, c)
			if len(c) == 1 && !hitCols[c[0]] {
				any = true
				hitCol = c[0]
				hitCols[hitCol] = true
			}
		}

		if !any {
			break
		}

		for r, c := range ruleColumns {
			if len(c) > 1 {
				ruleColumns[r] = without(c, hitCol)
			}
		}
	}

	interestingColumns := []string{
		"departure location",
		"departure station",
		"departure platform",
		"departure track",
		"departure date",
		"departure time",
	}

	product := 1
	for _, col := range interestingColumns {
		product *= myTicket[ruleColumns[col][0]]
	}

	fmt.Println("P2:", product)
}

func without(a []int, r int) (b []int) {
	for _, v := range a {
		if v != r {
			b = append(b, v)
		}
	}
	return
}

func ticketValid(t ticket, rules map[string]ranges) bool {
	for _, value := range t {
		valueValid := false
		for _, rule := range rules {
			if rule.contains(value) {
				valueValid = true
				break
			}
		}
		if !valueValid {
			return false
		}
	}
	return true
}

func sumInvalid(ts []ticket, rules map[string]ranges) {
	inv := 0
	sum := 0

	for _, t := range ts {
		for _, value := range t {
			valid := false
			for _, rule := range rules {
				if rule.contains(value) {
					valid = true
					break
				}
			}
			if !valid {
				sum += value
				inv++
			}
		}
	}

	fmt.Println("P1:", sum)
	fmt.Println("invalid count:", inv)
}

type ranges struct {
	astart, aend int
	bstart, bend int
}

func (r ranges) contains(i int) bool {
	return i >= r.astart && r.aend >= i || i >= r.bstart && r.bend >= i
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v failed to convert to int\n", s)
		os.Exit(1)
	}
	return v
}

type ticket []int

func parse(content string) (map[string]ranges, []ticket) {
	rules := map[string]ranges{}
	tickets := []ticket{}

	parts := strings.Split(content, "-- tickets --\n")
	rules_s := parts[0]
	tickets_s := parts[1]

	for _, rule_string := range strings.Split(rules_s, "\n") {
		if len(rule_string) == 0 {
			break
		}
		values := strings.Split(rule_string, ",")
		astart := atoi(values[1])
		aend := atoi(values[2])
		bstart := atoi(values[3])
		bend := atoi(values[4])

		rules[values[0]] = ranges{astart, aend, bstart, bend}
	}

	for _, ticket_string := range strings.Split(tickets_s, "\n") {
		if len(ticket_string) == 0 {
			break
		}
		fields := strings.Split(ticket_string, ",")
		var t ticket
		for _, field := range fields {
			v := atoi(field)
			t = append(t, v)
		}
		tickets = append(tickets, t)
	}

	return rules, tickets
}
