package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// This alt file contains some effort at refactoring.

// hasBag recursively searches ruleSets for the intended bag. First, it grabs
// the set of requirements from ruleSets about the currently given bag. Then, if
// the set directly requires the intended bag, then it returns true. Else, it
// loops around the requirement set and check each of those requirement's
// requirements for the wanted bag using this same function, which is recursive.

// countBags recursively counts how many bags are in the given parent bag.
// First, it iterates over all the parent bag's requirement. Since each bag
// requires a certain number of other bags in them, those numbers are added to
// the total. Then, we find the total number of bags inside those children bags,
// multiply that to the total number of said bags required, then add that to the
// total as well.
//
// To visualize this, we have
//
//    1xParent-----------\
//    |                  |
//    1xA-------\        1xB---------\
//    |         |        |           |
//    1xC       2xD      |           |
//              \------->1xE-------->1xF
//                       |           |
//                       1xG         1xH
//
// In the above graph, A has 2 bags, C and D, so we add 2 into the total. D is
// special - it depends on E, so we must count E's children as well. The
// visualization tells us E has 3 children, but since we have 2 D bags, we must
// multiply 3 by 2, giving us 6 bags in the D branch without D itself. Adding
// them gives us 2 (E, D) + 6 (D children bags) = 8 bags in total.
//
// Since this function finds the number of bags in the given parent, we could
// use this same function to find D's children, meaning we're using the function
// recursively.
func countBags(ruleSets map[string]map[string]int, parent string) (total int) {
	for name, num := range ruleSets[parent] {
		total += num
		total += num * countBags(ruleSets, name)
	}

	return
}

func main() {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}

	var bagRules = strings.Split(string(f), "\n")

	var ruleSets = map[string]map[string]int{}

	for _, bagRule := range bagRules {
		main, all := parse(bagRule)
		ruleSets[main] = all
	}

	fmt.Println(countBags(ruleSets, "shiny gold bag"))
}

func parse(input string) (mainbag string, bagRule map[string]int) {
	parts := strings.Split(input, " contain ")
	bagRule = map[string]int{}

	// trim plural
	mainbag = strings.TrimSuffix(parts[0], "s")

	if len(parts) == 1 || parts[1] == "no other bags." {
		return
	}

	bagTypes := strings.Split(parts[1], ", ")

	for _, bagType := range bagTypes {
		// trim punct
		bagType = strings.TrimSuffix(bagType, ".")
		// trim plural
		bagType = strings.TrimSuffix(bagType, "s")

		parts := strings.SplitN(bagType, " ", 2)

		num, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		bagRule[parts[1]] = num
	}

	return
}
