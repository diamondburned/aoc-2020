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

// I can't believe I didn't think of this!
// Credit to someone from the Discord Gophers server for stealing the idea off
// of someone else.

var binHack = strings.NewReplacer(
	"F", "0", "L", "0",
	"B", "1", "R", "1",
)

func bseat(input string) (int, int) {
	input = binHack.Replace(input)
	row, _ := strconv.ParseInt(input[:7], 2, 0)
	col, _ := strconv.ParseInt(input[7:], 2, 0)
	return int(row), int(col)
}

func seatID(row, col int) int {
	return row*8 + col
}
