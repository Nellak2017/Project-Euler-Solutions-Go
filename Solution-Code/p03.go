package main

import (
	"fmt"
	"math"
)

// Project Euler Problem 2
// Largest Prime Factor of 600851475143
// Author: Connor Keenum
// Date-Time: 26 Feb 2022 , 3:30 pm CST
// Build: Pass
func maxPrimeFactor(n int) int {
	maxPrime := -1

	for n%2 == 0 {
		maxPrime = 2
		n /= 2
	}

	for i := 3; i < int(math.Sqrt(float64(n))+1); i += 2 {
		for n%i == 0 {
			maxPrime = i
			n = n / i
		}
	}

	if n > 2 {
		maxPrime = n
	}

	return int(maxPrime)
}

func main() {
	fmt.Println("Max Prime Factor of 600851475143:", maxPrimeFactor(600851475143))
}
