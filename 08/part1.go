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

var (
	acc = "acc"
	jmp = "jmp"
	nop = "nop"
)

func main() {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}

	var accumulator = 0
	var instructions = strings.Split(string(f), "\n")

	var executed = map[int]int{}

	for i := 0; i < len(instructions); {
		parts := strings.Fields(instructions[i])
		op, v := parts[0], parts[1]

		n, err := strconv.Atoi(v)
		if err != nil {
			log.Panicln(err)
		}

		// Second time?
		if time := executed[i]; time > 0 {
			fmt.Println(accumulator)
			break
		}

		executed[i]++

		switch op {
		case acc:
			accumulator += n
		case jmp:
			i += n
			continue
		case nop:
		}

		i++
	}
}
