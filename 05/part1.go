package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// This code to remain largely unaltered from the state it was used to solve,
// mostly for historical purposes.

const (
	totalRows = 128
	totalCols = 8
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	var maxSeatID int

	var scan = bufio.NewScanner(f)
	for scan.Scan() {
		row, col := bseat(scan.Text())
		seatID := seatID(row, col)

		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	fmt.Println(maxSeatID)
}

func bseat(input string) (row, col int) {
	rowInputs := input[:7]
	colInputs := input[7:]

	rows := [2]int{0, totalRows}
	cols := [2]int{0, totalCols}

	for _, r := range rowInputs {
		switch r {
		case 'F':
			rows[1] -= (rows[1] - rows[0]) / 2
		case 'B':
			rows[0] += (rows[1] - rows[0]) / 2
		}
	}

	for _, c := range colInputs {
		switch c {
		case 'L':
			cols[1] -= (cols[1] - cols[0]) / 2
		case 'R':
			cols[0] += (cols[1] - cols[0]) / 2
		}
	}

	return rows[0], cols[0]
}

func seatID(row, col int) int {
	return row*8 + col
}
