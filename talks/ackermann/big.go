package main

import (
	"fmt"
	"math/big"
)

func main() {
	// A(4, 2) = 2^65536 - 3
	a := new(big.Int).Lsh(big.NewInt(1), 65536)
	a.Sub(a, big.NewInt(3))

	digits := a.String()
	fmt.Println("digits:", len(digits))
	fmt.Println("starts:", digits[:20])
	fmt.Println("ends:  ", digits[len(digits)-20:])
}
