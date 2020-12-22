package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var input = flag.String("input", "input", "Input file")

type recipie struct {
	ingredients []string
	allergens   []string
}

func parseRecipie(s string) recipie {
	var r recipie
	spl := strings.Split(s, " (")
	r.ingredients = strings.Split(spl[0], " ")

	if len(spl) > 1 {
		// remove "(contains " and ")"
		algns := spl[1][9 : len(spl[1])-1]
		r.allergens = strings.Split(algns, ", ")
	} else {
		r.allergens = []string{}
	}

	return r
}

func main() {
	flag.Parse()

	bcontent, err := ioutil.ReadFile(*input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	lines := strings.Split(string(bcontent), "\n")
	fmt.Println(lines)
	recipies := []recipie{}
	for _, line := range lines {
		if len(line) < 3 {
			continue
		}
		recipies = append(recipies, parseRecipie(line))
	}
	// for _, r := range recipies {
	// 	fmt.Println(r)
	// }

	// {allergen -> {ingredients}}
	allergenIngredients := map[string]map[string]bool{}

	for _, r := range recipies {
		for _, a := range r.allergens {
			if p, ok := allergenIngredients[a]; ok {
				intersection := map[string]bool{}
				for _, i := range r.ingredients {
					if v, ok := p[i]; v && ok {
						intersection[i] = true
					}
				}
				allergenIngredients[a] = intersection
			} else {
				ai := map[string]bool{}
				for _, i := range r.ingredients {
					ai[i] = true
				}
				allergenIngredients[a] = ai
			}
		}
	}

	allergenIngList := map[string][]string{}

	for a, ai := range allergenIngredients {
		allergenIngList[a] = []string{}
		for v := range ai {
			allergenIngList[a] = append(allergenIngList[a], v)
		}
	}

	hitCols := map[string]bool{}

	for {
		progress := false
		for allergen, ingredients := range allergenIngList {
			if len(ingredients) == 1 && !hitCols[allergen] {
				hitCols[allergen] = true
				progress = true

				for allergen_x, ingredients_x := range allergenIngList {
					if allergen == allergen_x {
						continue
					}
					allergenIngList[allergen_x] = without(ingredients_x, ingredients[0])
				}
			}
		}
		if !progress {
			break
		}
	}

	var allAllergens []string

	for all, v := range allergenIngList {
		for _, a := range v {
			allAllergens = append(allAllergens, a)
			fmt.Println(all, a)
		}
	}

	total := 0

	for _, r := range recipies {
		for _, i := range r.ingredients {
			if !in(allAllergens, i) {
				total++
			}
		}
	}

	fmt.Println("P1:", total)

	sort.Strings(allAllergens)
	fmt.Print("P2: ")
	for _, a := range allAllergens {
		fmt.Printf("%v,", a)
	}
	fmt.Println()
}

func without(a []string, r string) (b []string) {
	for _, v := range a {
		if v != r {
			b = append(b, v)
		}
	}
	return
}

func in(a []string, r string) bool {
	for _, v := range a {
		if v == r {
			return true
		}
	}
	return false
}
