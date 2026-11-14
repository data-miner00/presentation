package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// group matches ^(a+)+b from position i the way a backtracking engine does:
// take a run of a's as one group, then try another group, then try b.
// When both fail, give back one a and try again.
func group(s string, i int, steps *int) bool {
	*steps++
	end := i
	for end < len(s) && s[end] == 'a' {
		end++
	}
	for j := end; j > i; j-- { // a+ is greedy: longest group first
		if group(s, j, steps) || (j < len(s) && s[j] == 'b') {
			return true
		}
	}
	return false
}

func main() {
	for _, n := range []int{4, 8, 12, 16, 20, 24} {
		s := strings.Repeat("a", n) + "c" // no b, so every split fails
		steps := 0
		group(s, 0, &steps)
		fmt.Printf("%2d a's: %10d steps\n", n, steps)
	}

	s := strings.Repeat("a", 100000) + "c"
	start := time.Now()
	matched := regexp.MustCompile(`^(a+)+b`).MatchString(s)
	fmt.Printf("\nGo's regexp, 100000 a's: %v in %v\n", matched, time.Since(start).Round(time.Microsecond))
}
