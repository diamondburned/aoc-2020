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
func hasBag(ruleSets map[string]map[string]int, bag, has string) bool {
	rules, ok := ruleSets[bag]
	if !ok {
		return false
	}

	if _, ok := rules[has]; ok {
		return true
	}

	for bag := range rules {
		if hasBag(ruleSets, bag, has) {
			return true
		}
	}

	return false
}

func main() {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}

	var bagRules = strings.Split(string(f), "\n")

	var canShinyGoldBags int

	var ruleSets = map[string]map[string]int{}

	for _, bagRule := range bagRules {
		main, all := parse(bagRule)
		ruleSets[main] = all
	}

	for bag := range ruleSets {
		if hasBag(ruleSets, bag, "shiny gold bag") {
			canShinyGoldBags++
		}
	}

	fmt.Println(canShinyGoldBags)
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
