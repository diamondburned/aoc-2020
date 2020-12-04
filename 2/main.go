package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	var first, second int
	var letter byte
	var password string

	var parts = []func(string, byte, int, int) bool{part1, part2}
	var valid = make([]int, len(parts))

	var scan = bufio.NewScanner(f)
	for scan.Scan() {
		_, err := fmt.Sscanf(scan.Text(), "%d-%d %c: %s\n", &first, &second, &letter, &password)
		if err != nil {
			log.Panicln("failed to scan:", err)
		}

		for i, part := range parts {
			if part(password, letter, first, second) {
				valid[i]++
			}
		}
	}

	for i, part := range valid {
		fmt.Printf("Part %d: %d\n", i+1, part)
	}
}

func part1(password string, letter byte, from, to int) bool {
	count := strings.Count(password, string(letter))
	return from <= count && count <= to
}

func part2(password string, letter byte, first, second int) bool {
	// One indexed.
	first--
	second--

	has1 := password[first] == letter
	has2 := password[second] == letter

	return (has1 && !has2) || (!has1 && has2)
}
