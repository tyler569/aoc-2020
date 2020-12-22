package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var input = flag.String("input", "input", "Input file")

type tile struct {
	lines []string
	id    int
}

func newTile(s string) tile {
	lines := strings.Split(s, "\n")
	t := tile{}
	t.lines = lines[1:11]
	_, err := fmt.Sscanf(lines[0], "Tile %v:", &t.id)
	if err != nil {
		fatalf("Failed to parse %v; %v\n", lines[0], err)
	}
	return t
}

// go around the tile, converting each edge into an integer to use
// to match up with other tiles.
func (t tile) edges() []int {
	length := len(t.lines)
	max := length-1

	addE := func(c byte, i int) int {
		r := i << 1
		if c == '#' {
			r |= 1
		}
		return r
	}

	comp := func(lines []string, ix, iy, dx, dy int) (int, int) {
		var fwd, bwd int
		for i := 0; i < length; i++ {
			fwd = addE(lines[ix+i*dx][iy+i*dy], fwd)
		}
		for i := max; i >= 0; i-- {
			bwd = addE(lines[ix+i*dx][iy+i*dy], bwd)
		}
		return fwd, bwd;
	}

	result := make([]int, 8)

	result[0], result[1] = comp(t.lines, 0, 0, 1, 0)
	result[2], result[3] = comp(t.lines, 0, 0, 0, 1)
	result[4], result[5] = comp(t.lines, 0, max, 1, 0)
	result[6], result[7] = comp(t.lines, max, 0, 0, 1)

	return result
}

func main() {
	flag.Parse()

	content, err := ioutil.ReadFile(*input)
	if err != nil {
		fatalf("Could not read file %v; %v\n", *input, err)
	}

	tilesStrings := strings.Split(string(content), "\n\n")
	tiles := []tile{}

	for _, t := range tilesStrings {
		tiles = append(tiles, newTile(t))
	}

	for _, t := range tiles {
		fmt.Printf("%+v %v\n", t, t.edges())
	}
}

func fatal(values ...interface{}) {
	fmt.Fprintln(os.Stderr, values...)
	os.Exit(1)
}

func fatalf(format string, values ...interface{}) {
	fmt.Fprintf(os.Stderr, format, values...)
	os.Exit(1)
}
