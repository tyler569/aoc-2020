package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var input = flag.String("file", "input", "AoC Input File")

func main() {
	flag.Parse()

	contents, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(contents), "\n")
	stime := lines[0]
	busses := lines[1]

	time, err := strconv.Atoi(stime)
	if err != nil {
		log.Fatal(err)
	}

	bestBus := 0
	bestTime := 100000

	busLine := strings.Split(busses, ",")
	for _, busLine := range busLine {
		if busLine == "x" {
			continue
		}

		busNumber, err := strconv.Atoi(busLine)
		if err != nil {
			log.Fatal(err)
		}

		dt := busNumber - time%busNumber
		if dt < bestTime {
			bestTime = dt
			bestBus = busNumber
		}
	}

	fmt.Println("P1:", bestBus, bestTime, bestBus*bestTime)

	/*
	The key insight is that b%t == n for all busses can be determined
	incrementally, find where this is true for a subset and then that sub
	example will be true for all increments of the LCM of all the factors.
	These factors are all prime, so that is a trivial multiplication.
	*/
	runningLcm := 1
	acc := 0

	for n, busLine := range busLine {
		if busLine == "x" {
			continue
		}

		busNumber, err := strconv.Atoi(busLine)
		if err != nil {
			log.Fatal(err)
		}
		
		for (acc + n)%busNumber != 0 {
			// Increment by (LCM of the previous) until we find an example
			// that works for this bus
			acc += runningLcm
		}

		// It will now work for all multiples of (LCM of the total)
		runningLcm *= busNumber
	}

	fmt.Println("P2:", acc)
}
