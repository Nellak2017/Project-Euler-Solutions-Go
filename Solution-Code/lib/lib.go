package lib

import (
	"math/rand"
	"time"
)

// Project Euler Library
// Author: Connor Keenum
// Date-Time: 26 Feb 2022 , 4:32 pm CST
// Build: Pass

// F

// Returns factorial of n for n in range [0,20]
// For a larger factorial range, we will need to use bigInt library
func Fact(n uint64) uint64 {
	product := 1
	for i := 1; uint64(i) < n; i++ {
		product *= i + 1
	}
	return uint64(product)
}

// P

// Tests if a number is a prime number or not using Miller-Rabin algorithm
func ProbabilyIsPrime(n uint32) bool {
	// Implementation uses the Miller-Rabin Primality Test
	// The optimal number of rounds for this test is 40
	// See http://stackoverflow.com/questions/6325576/how-many-iterations-of-rabin-miller-should-i-use-for-cryptographic-safe-primes
	// bases of 2, 7, 61 are sufficient to cover 2^32
	switch n {
	case 0, 1:
		return false
	case 2, 7, 61:
		return true
	}
	// compute s, d where 2^s * d = n-1
	nm1 := n - 1
	d := nm1
	s := 0
	for d&1 == 0 {
		d >>= 1
		s++
	}
	n64 := uint64(n)
	for _, a := range []uint32{2, 7, 61} {
		// compute x := a^d % n
		x := uint64(1)
		p := uint64(a)
		for dr := d; dr > 0; dr >>= 1 {
			if dr&1 != 0 {
				x = x * p % n64
			}
			p = p * p % n64
		}
		if x == 1 || uint32(x) == nm1 {
			continue
		}
		for r := 1; ; r++ {
			if r >= s {
				return false
			}
			x = x * x % n64
			if x == 1 {
				return false
			}
			if uint32(x) == nm1 {
				break
			}
		}
	}
	return true
}

// R

// Go implementation of Python's random number range generator
// Returns a random number in range [start, stop] inclusive
// if start > stop, it returns 0 as a default value
func RandRange(start uint32, stop uint32) uint32 {
	rand.Seed(time.Now().UnixNano()) // tests are written for rand.Seed(1), time.Now().UnixNano()
	if start <= stop {
		return uint32(rand.Intn(int(stop-start+1))) + start
	} else {
		return 0
	}
}
