package lib

import (
	"testing"
)

/*
Tests covered:
	Fact -> tests Passed to n = 20
	ProbabilyIsPrime -> tests Passed to prime <= 4000000007
	RandRange -> tests Passed
*/

// Fact tests --> PASSING
type FactResult struct {
	n        uint64
	expected uint64
}

var factResults = []FactResult{
	{0, 1},
	{1, 1},
	{2, 2},
	{3, 6},
	{4, 24},
	{8, 40320},
	{10, 3628800},
	{12, 479001600},
	{20, 2432902008176640000},
}

func TestFactResults(t *testing.T) {
	for _, test := range factResults {
		result := Fact(test.n)
		if result != test.expected {
			t.Fatal("ProbabilyIsPrime Test FAILED. Expected:", test.expected, " got:", result)
		}
	}
}

// End of Fact tests

// IsPrime tests --> PASSING
type ProbabilyIsPrimeResult struct {
	n        uint32
	expected bool
}

var probablyIsPrimeResults = []ProbabilyIsPrimeResult{
	{0, false},
	{1, false},
	{2, true},
	{3, true},
	{4, false},
	{5, true},
	{6, false},
	{7, true},
	{9, false},
	{11, true},
	{25, false},
	{53, true},
	{59, true},
	{60, false},
	{97, true},
	{541, true},
	{7919, true},
	{7920, false},
	{104729, true},
	{1299709, true},
	{1299710, false},
	{4000000006, false},
	{4000000007, true},
	{3045000031, true},
}

func TestProbabilyIsPrime(t *testing.T) {
	for _, test := range probablyIsPrimeResults {
		result := ProbabilyIsPrime(test.n)
		if result != test.expected {
			t.Fatal("ProbabilyIsPrime Test FAILED. Expected:", test.expected, " got:", result)
		}
	}
}

// End of IsPrime tests

// RandRange tests (seed = 1) --> PASSING
type RandRangeResult struct {
	start    uint32
	stop     uint32
	expected uint32
}

var randRangeResults = []RandRangeResult{
	{0, 5, 5},
	{0, 10, 1},
	{0, 1, 1},
	{0, 2, 2},
	{0, 4, 1},
	{2, 4, 4},
}

func TestRandRange(t *testing.T) {
	for _, test := range randRangeResults {
		result := RandRange(test.start, test.stop)
		if result != test.expected {
			t.Fatal("RandRange Test FAILED. Expected:", test.expected, " got:", result)
		}
	}
}

// End of RandRange tests
