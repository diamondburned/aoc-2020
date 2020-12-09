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

	var lines = strings.Split(string(f), "\n")

	const preambleLen = 25

	var preamble = make([]int, 0, preambleLen)

	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		preamble = append(preamble, n)

		if len(preamble) > preambleLen {
			if !isValidSum(preamble, n) {
				fmt.Println(n, "is invalid.")
				break
			}

			// Cut the first out.
			copy(preamble[0:], preamble[1:])
			preamble = preamble[:len(preamble)-1]
		}
	}
}

func isValidSum(preamble []int, n int) bool {
	for _, preambNum1 := range preamble {
		for _, preambNum2 := range preamble {
			if preambNum2+preambNum1 == n {
				return true
			}
		}
	}
	return false
}
