package main

import "fmt"

var calls int

func ackermann(m, n uint) uint {
	calls++
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
	for n := uint(0); n <= 8; n++ {
		calls = 0
		value := ackermann(3, n)
		fmt.Printf("A(3, %d) = %4d   %7d calls\n", n, value, calls)
	}
}
