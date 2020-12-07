package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// This code to remain largely unaltered from the state it was used to solve,
// mostly for historical purposes.

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	var treeMap [][]bool
	var width int

	var scan = bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		width = len(line)

		trees := make([]bool, width)

		for i, c := range line {
			trees[i] = c == '#'
		}

		treeMap = append(treeMap, trees)
	}

	var pairs = [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	var bumps = 1

	for _, pair := range pairs {
		bumps *= findBump(treeMap, pair[0], pair[1])
	}

	fmt.Println(bumps)
}

func findBump(treeMap [][]bool, right, down int) int {
	var bump = 0
	var downStep = down
	var rightStep = right

	for down < len(treeMap) {
		if treeMap[down][right] {
			bump++
		}

		right = (right + rightStep) % len(treeMap[0])
		down += downStep
	}

	return bump
}
