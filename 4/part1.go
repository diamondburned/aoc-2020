package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// This code to remain largely unaltered from the state it was used to solve,
// mostly for historical purposes.

func main() {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}

	var fields = []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		"cid",
	}

	var valid int

	var passports = strings.Split(string(f), "\n\n")

	for _, passport := range passports {
		words := strings.Fields(passport)

		var present int
		var hasCID bool

		for _, word := range words {
			for _, field := range fields {
				if strings.HasPrefix(word, field+":") {
					present++

					if field == "cid" {
						hasCID = true
					}
				}
			}
		}

		if present == 8 || (present == 7 && !hasCID) {
			valid++
		}
	}

	fmt.Println(valid)
}
