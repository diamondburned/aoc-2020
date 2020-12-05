package main

import (
	"fmt"
	"strconv"
	"strings"
)

// This code to remain largely unaltered from the state it was used to solve,
// mostly for historical purposes.

const (
	totalRows = 128
	totalCols = 8
)

func main() {
	// f, err := os.Open("input")
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// defer f.Close()

	fmt.Println(bseat("BFFFBBFRRR"))

	// var maxSeatID int

	// var scan = bufio.NewScanner(f)
	// for scan.Scan() {
	// 	row, col := bseat(scan.Text())
	// 	seatID := seatID(row, col)

	// 	if seatID > maxSeatID {
	// 		maxSeatID = seatID
	// 	}
	// }

	// fmt.Println(maxSeatID)
}

// I can't believe I hadn't thought of this!

func bseat(input string) (int, int) {
	rowInputs := input[:7]
	rowInputs = strings.ReplaceAll(rowInputs, "F", "0")
	rowInputs = strings.ReplaceAll(rowInputs, "B", "1")

	row, _ := strconv.ParseInt(rowInputs, 2, 0)

	colInputs := input[7:]
	colInputs = strings.ReplaceAll(colInputs, "L", "0")
	colInputs = strings.ReplaceAll(colInputs, "R", "1")

	col, _ := strconv.ParseInt(colInputs, 2, 0)

	return int(row), int(col)
}

func seatID(row, col int) int {
	return row*8 + col
}
