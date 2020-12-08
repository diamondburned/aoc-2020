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

	type ins struct {
		op string
		v  int
	}

	var lines = strings.Split(string(f), "\n")
	var instructions = make([]ins, 0, len(lines))

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		op, v := parts[0], parts[1]

		n, err := strconv.Atoi(v)
		if err != nil {
			log.Panicln(err)
		}

		instructions = append(instructions, ins{op, n})
	}

	var accumulator = 0
	var executed = map[int]int{}

	var lastJmp int

	var allJmps []int

	for i, ins := range instructions {
		if ins.op == jmp {
			allJmps = append(allJmps, i)
		}
	}

	var jmpTested int

NopTest:
	for i, ijmp := range allJmps {
		jmpTested = i
		instructions := append([]ins(nil), instructions...)
		instructions[ijmp].op = "nop"

		accumulator = 0
		lastJmp = -1
		executed = map[int]int{}

		for i := 0; i < len(instructions); {
			if executed[i] > 10 && lastJmp >= 0 {
				continue NopTest
			}

			executed[i]++

			switch ins := instructions[i]; ins.op {
			case acc:
				accumulator += ins.v
			case jmp:
				lastJmp = i
				i += ins.v
				continue
			case nop:
			}

			i++
		}

		break
	}

	fmt.Println("jmpTested:", jmpTested, "out of", len(allJmps))
	fmt.Println(accumulator)
}
