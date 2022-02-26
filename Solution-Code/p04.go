package main

import (
	"fmt"
)

// Project Euler Problem 4
// Find the largest Palindrome made from the product of two 3-digit numbers
// Author: Connor Keenum
// Date-Time: 26 Feb 2022 , 3:39 pm CST
// Build: Pass
func isPalindrome(s int) bool {
	// Test if an integer is a palindrome
	// it relies on the laws of arithmetic, and won't work for a generalized string
	// see also: https://github.com/golang/go/wiki/SliceTricks
	// see also: https://www.tutorialspoint.com/write-a-golang-program-to-check-whether-a-given-number-is-a-palindrome-or-not
	in, remainder, res := s, 0, 0
	for s > 0 {
		remainder = s % 10
		res = (res * 10) + remainder
		s /= 10
	}
	if in == res {
		return true
	} else {
		return false
	}
}

func main() {
	product := 1
	for i := 100; i < 999; i++ {
		for x := 100; x < 999; x++ {
			if isPalindrome(i*x) && (i*x) > product {
				product = i * x
			}
		}
	}
	fmt.Println("Largest Palindrome made from the product of two 3-digit numbers:", product)
}
