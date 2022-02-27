package main

import (
	E "Solution-Code/lib"
	"fmt"
)

// Project Euler Problem 5
// Smallest positive number that n%2 to n%20 == 0
// Author: Connor Keenum
// Date-Time: 26 Feb 2022 , 6:34 pm CST
// Build: Pass

func primesProduct(n uint32) int { // correct
	product := 1 // correct
	for x := 2; uint32(x) < n+1; x++ {
		if E.ProbabilyIsPrime(uint32(x)) {
			product *= x
		}
	}
	return product
}

func optimized(a uint32, b uint32) uint32 {
	// for the number to be a candidate, it will need to be divisible by all the primes less than
	// b. Using this fact, one could search this space from that prime product up to b! at prime_Product intervals
	// This reduces the search space from atleast b! to only (b!-primesProduct)/primesProduct(b)
	theNum := E.Fact(uint64(b)) // correct
	for x := primesProduct(uint32(b)); uint64(x) < E.Fact(uint64(b)); x += primesProduct(uint32(b)) {
		// starting at b and counting to fact(b) by the product of primes less than b
		for i := b; i > a; i-- { // if it's not a candidate, prune it fast
			if uint32(x)%i != 0 {
				break
			} else if i == 2 {
				if uint64(x) < theNum {
					return uint32(x)
				} else {
					return uint32(theNum)
				}
			}
		}
	}
	return 0
}

func main() {
	fmt.Println("The Answer: ", optimized(1, 20))
}
