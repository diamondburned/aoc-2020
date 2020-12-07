package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// This code to remain largely unaltered from the state it was used to solve,
// mostly for historical purposes.
//
// This challenge was really hard for me, and frankly I wrote all this code
// without even understanding how it works. The 594 magic number was the total
// line of my input file, which I was really desperate to solve, so I thought
// maybe I'll just brute it until I've looped through the entire file.

func main() {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}

	var bagRules = strings.Split(string(f), "\n")

	var canShinyGoldBags = map[string]struct{}{}

	var ruleSets = map[string][]bagRule{}

	for _, bagRule := range bagRules {
		main, all := parse(bagRule)
		ruleSets[main] = all
	}

	for main, all := range ruleSets {
		for _, rule := range all {
			if rule.name == "shiny gold bag" {
				canShinyGoldBags[main] = struct{}{}
			}
		}
	}

	for i := 0; i < 594; i++ {
		for main, all := range ruleSets {
			for _, rule := range all {
				for canShinyGoldBag := range canShinyGoldBags {
					if canShinyGoldBag == rule.name {
						canShinyGoldBags[main] = struct{}{}
					}
				}
			}
		}
	}

	fmt.Println(len(canShinyGoldBags))

	// var total int

	// for _, group := range groups {
	// 	var answer = map[rune]struct{}{}

	// 	for _, char := range group {
	// 		if char == '\n' {
	// 			continue
	// 		}

	// 		answer[char] = struct{}{}
	// 	}

	// 	total += len(answer)
	// }

	// fmt.Println(total)
}

type bagRule struct {
	name string
	num  int
}

func parse(input string) (mainbag string, allbags []bagRule) {
	parts := strings.Split(input, " contain ")

	// trim plural
	mainbag = strings.TrimSuffix(parts[0], "s")

	if len(parts) == 1 || parts[1] == "no other bags." {
		return mainbag, nil
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

		allbags = append(allbags, bagRule{parts[1], num})
	}

	return
}
