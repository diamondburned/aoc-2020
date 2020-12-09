package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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

	var numbers = make([]int, 0, len(lines))

	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		numbers = append(numbers, n)
	}

	var illegalN int
	var preamble = make([]int, 0, preambleLen)

	for _, n := range numbers {
		preamble = append(preamble, n)

		if len(preamble) > preambleLen {
			if !isValidSum(preamble, n) {
				illegalN = n
				break
			}

			// Cut the first out.
			copy(preamble[0:], preamble[1:])
			preamble = preamble[:len(preamble)-1]
		}
	}

	from, to := arraySum(numbers, illegalN)
	cont := append([]int(nil), numbers[from:to]...)
	sort.Ints(cont)

	fmt.Println(cont[0] + cont[len(cont)-1])
}

func arraySum(numbers []int, sum int) (from, to int) {
	var currentSum int

	for i := 0; i < len(numbers); i++ {
		currentSum = numbers[i]

		for j := i + 1; j < len(numbers); j++ {
			if currentSum == sum {
				return i, j - 1
			}

			if currentSum > sum || j == len(numbers) {
				break
			}

			currentSum += numbers[j]
		}
	}

	return -1, -1
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
