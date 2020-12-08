package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type op int

const (
	NOP op = iota
	ACC
	JMP
)

func (o op) String() string {
	switch o {
	case NOP:
		return "NOP"
	case ACC:
		return "ACC"
	case JMP:
		return "JMP"
	}
	return "invalid"
}

type instr struct {
	op
	arg     int
	visited bool
}

func (i instr) String() string {
	return fmt.Sprintf("[%v %+5d]", i.op, i.arg)
}

func parseInstr(s string) (i instr, err error) {
	split := strings.Split(s, " ")
	switch split[0] {
	case "nop":
		i.op = NOP
	case "acc":
		i.op = ACC
	case "jmp":
		i.op = JMP
	default:
		err = errors.New("Invalid instruction: " + s)
		return
	}

	i.arg, err = strconv.Atoi(split[1])
	return
}

func runProgram(program []instr) (acc int, term bool, err error) {
	pc := 0

	for {
		if pc == len(program) {
			term = true
			return
		}
		i := program[pc]
		if i.visited {
			return
		}
		program[pc].visited = true
		switch i.op {
		case NOP:
			pc += 1
		case ACC:
			acc += i.arg
			pc += 1
		case JMP:
			pc += i.arg
		default:
			err = errors.New("Invalid instruction")
			return
		}
	}
}

func main() {
	fmt.Println("Hello World")

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var instrs []instr

	for scanner.Scan() {
		s := scanner.Text()
		i, err := parseInstr(s)
		if err != nil {
			log.Fatal(err)
		}
		instrs = append(instrs, i)
	}

	for i, in := range instrs {
		if i > 0 && i%8 == 0 {
			fmt.Println()
		}
		fmt.Printf("%v ", in)
	}
	fmt.Println()

	cpy := make([]instr, len(instrs))
	copy(cpy, instrs)

	acc, term, err := runProgram(cpy)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("P1:", acc)

	for i := 0; i < len(instrs); i++ {
		copy(cpy, instrs)

		if cpy[i].op == JMP {
			cpy[i].op = NOP
		} else if cpy[i].op == NOP {
			cpy[i].op = JMP
		}

		acc, term, err = runProgram(cpy)
		if err != nil {
			log.Fatal(err)
		}
		if term {
			fmt.Println("P2:", acc)
		}
	}
}
