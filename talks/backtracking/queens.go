package main

import (
	"fmt"
	"strings"
)

const n = 8

var (
	cols      [n]int // cols[r] is the column of the queen in row r
	first     [n]int // the first solution found
	solutions int
	visited   int // partial boards the search looked at
)

// place puts a queen in each row from row onwards, trying every safe column.
func place(row int) {
	visited++
	if row == n {
		if solutions == 0 {
			first = cols
		}
		solutions++
		return
	}
	for c := 0; c < n; c++ {
		if !safe(row, c) {
			continue // attacked: skip this whole branch
		}
		cols[row] = c  // choose
		place(row + 1) // explore
		// unchoose: nothing to do, the next column overwrites cols[row]
	}
}

// safe reports whether no queen in the rows above attacks (row, c).
func safe(row, c int) bool {
	for r := 0; r < row; r++ {
		d := row - r
		if cols[r] == c || cols[r] == c-d || cols[r] == c+d {
			return false
		}
	}
	return true
}

func main() {
	place(0)
	for _, c := range first {
		fmt.Println(strings.Repeat(". ", c) + "Q " + strings.Repeat(". ", n-c-1))
	}
	fmt.Println()
	fmt.Println("solutions:           ", solutions)
	fmt.Println("partial boards tried:", visited)
}
