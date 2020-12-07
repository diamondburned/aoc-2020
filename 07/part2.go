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

func main() {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}

	var bagRules = strings.Split(string(f), "\n")

	var ruleSets = map[string][]bagRule{}

	for _, bagRule := range bagRules {
		main, all := parse(bagRule)
		ruleSets[main] = all
	}

	shinyGoldRules := ruleSets["shiny gold bag"]

	totalBags := countBags("shiny gold bag", shinyGoldRules, ruleSets)

	// for main, all := range ruleSets {
	// 	for _, rule := range all {
	// 		if rule.name == "shiny gold bag" {
	// 			canShinyGoldBags[main] = struct{}{}
	// 		}
	// 	}
	// }

	// for i := 0; i < 594; i++ {
	// 	for main, all := range ruleSets {
	// 		for _, rule := range all {
	// 			for canShinyGoldBag := range canShinyGoldBags {
	// 				if canShinyGoldBag == rule.name {
	// 					if _, ok := canShinyGoldBags[main]; !ok {
	// 						totalBags += rule.num
	// 					}
	// 					canShinyGoldBags[main] = struct{}{}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	fmt.Println(totalBags)

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

func countBags(parent string, rules []bagRule, ruleSets map[string][]bagRule) (total int) {
	if len(rules) == 0 {
		return
	}

	log.Println(parent, "requires", len(rules), "bag types")

	for _, rule := range rules {
		total += rule.num

		log.Println(parent, "requires", rule.num, "bags of type", rule.name)

		// *total += rule.num * len(rules)

		rules := ruleSets[rule.name]
		total += rule.num * countBags(rule.name, rules, ruleSets)
	}

	return
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
