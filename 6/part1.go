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

	var groups = strings.Split(string(f), "\n\n")

	var total int

	for _, group := range groups {
		var answer = map[rune]struct{}{}

		for _, char := range group {
			if char == '\n' {
				continue
			}

			answer[char] = struct{}{}
		}

		total += len(answer)
	}

	fmt.Println(total)
}
