package main

import (
	"fmt"
)

// Project Euler Problem 1
// Sum of multiples of 3 or 5 up to 1000
// Author: Connor Keenum
// Date-Time: 26 Feb 2022 , 2:41 pm CST
// Build: Pass
func main() {

	// O(n) programmatic solution
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum of multiples:", sum)

	// O(1) mathematical solution
	/*
		s_3(1000) + s_5(1000) - s_15(1000) = sum of muiltiples of 3 or 5 up to 1000
		s_3(1000) = 3+6+9+12+15+...+999
		s_5(1000) = 5+10+15+20+25+...+1000
		s_15(1000)= 15+30+45+...+990
		Factor out and use summation formula:
		s_3(1000) = 3*(1+2+3+4+5+6+7+8+9+...+333)
		s_5(1000) = 5*(1+2+3+4+5+6+7+8+9+...+200)
		s_15(1000)= 15*(1+2+3+4+5+6+...+66)
		s_a(n) = a(n(n+1)/2)
		s_3(1000)+s_5(1000)-s_15(1000) = 3(333(334)/2)+5(200(201)/2)-15(66(67)/2)
		s_3(1000)+s_5(1000)-s_15(1000) = 166833 + 100500 - 33165
		s_3(1000)+s_5(1000)-s_15(1000) = 233168
	*/
}
