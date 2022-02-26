package main

import (
	"fmt"
)

// Project Euler Problem 2
// Sum of Even Fibbonacci numbers
// Author: Connor Keenum
// Date-Time: 26 Feb 2022 , 3:08 pm CST
// Build: Pass

func fibEven(n int64) int64 {
	a, b, res := 0, 1, 2
	for int64(b) <= n {
		if int64(b%2) == int64(0) {
			a += b
		}
		b, res = res, b+res
	}
	return int64(a)
}

func main() {
	fmt.Println("fibEven(4000000) =", fibEven(4000000))
}
