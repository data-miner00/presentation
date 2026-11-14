package main

import (
	"fmt"
	"strings"
)

// Grid holds digits 1 to 9, with 0 for an empty cell.
type Grid [9][9]byte

// solve fills g in place. pick chooses which empty cell to fill next.
func solve(g *Grid, pick func(*Grid) (r, c int, ok bool), calls *int) bool {
	*calls++
	r, c, ok := pick(g)
	if !ok {
		return true // no empty cell left: solved
	}
	for d := byte(1); d <= 9; d++ {
		if !g.allowed(r, c, d) {
			continue // breaks a row, column or box: prune
		}
		g[r][c] = d                // choose
		if solve(g, pick, calls) { // explore
			return true
		}
		g[r][c] = 0 // unchoose
	}
	return false // no digit fits here: back up
}

// allowed reports whether d can go at (r, c) without repeating a digit
// in its row, column or 3x3 box.
func (g *Grid) allowed(r, c int, d byte) bool {
	br, bc := r/3*3, c/3*3
	for i := 0; i < 9; i++ {
		if g[r][i] == d || g[i][c] == d || g[br+i/3][bc+i%3] == d {
			return false
		}
	}
	return true
}

// firstEmpty picks the next empty cell in reading order.
func firstEmpty(g *Grid) (int, int, bool) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if g[r][c] == 0 {
				return r, c, true
			}
		}
	}
	return 0, 0, false
}

// mostConstrained picks the empty cell with the fewest digits still allowed.
func mostConstrained(g *Grid) (int, int, bool) {
	best, br, bc := 10, 0, 0
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if g[r][c] != 0 {
				continue
			}
			options := 0
			for d := byte(1); d <= 9; d++ {
				if g.allowed(r, c, d) {
					options++
				}
			}
			if options < best {
				best, br, bc = options, r, c
			}
		}
	}
	return br, bc, best < 10
}

// parse reads 81 cells, digits or dots, ignoring anything else.
func parse(s string) Grid {
	var g Grid
	i := 0
	for _, ch := range s {
		switch {
		case ch == '.':
			i++
		case ch >= '1' && ch <= '9':
			g[i/9][i%9] = byte(ch - '0')
			i++
		}
	}
	return g
}

func (g Grid) String() string {
	var b strings.Builder
	for _, row := range g {
		for _, d := range row {
			fmt.Fprintf(&b, "%d ", d)
		}
		b.WriteString("\n")
	}
	return b.String()
}

func main() {
	puzzle := parse(`
		8........ ..36..... .7..9.2..
		.5...7... ....457.. ...1...3.
		..1....68 ..85...1. .9....4..`)
	a, b := puzzle, puzzle
	var naive, smart int
	solve(&a, firstEmpty, &naive)
	solve(&b, mostConstrained, &smart)
	fmt.Print(b)
	fmt.Println("\nfirst empty cell:", naive, "calls")
	fmt.Println("most constrained:", smart, "calls")
}
