package main

import "fmt"

func ackermann(m, n uint) uint {
	switch {
	case m == 0:
		return n + 1
	case n == 0:
		return ackermann(m-1, 1)
	default:
		return ackermann(m-1, ackermann(m, n-1))
	}
}

func main() {
	for m := uint(0); m <= 3; m++ {
		fmt.Printf("m=%d:", m)
		for n := uint(0); n <= 4; n++ {
			fmt.Printf("%5d", ackermann(m, n))
		}
		fmt.Println()
	}
}
