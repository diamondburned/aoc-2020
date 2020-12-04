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

	var bump int

	var right = 3
	for down := 1; down < len(treeMap); down++ {
		if treeMap[down][right] {
			bump++
		}

		right = (right + 3) % width
	}

	fmt.Println(bump)
}
