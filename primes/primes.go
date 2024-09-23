package primes

import (
	"math"
)

// Returns an array of primes lower than limit by implementing
// the sieve of Erathostenes
func Sieve(limit int) []int {
	controlArray := make([]bool, limit)
	cappedLimit := int(math.Sqrt(float64(limit)))
	for n := 0; n < limit; n = n + 2 {
		controlArray[n] = true
	}
	for n := 3; n < cappedLimit; n = n + 2 {
		if !controlArray[n] {
			for m := n * n; m < limit; m = m + 2*n {
				controlArray[m] = true
			}
		}
	}
	outarray := []int{}
	for v := range controlArray {
		if !controlArray[v] {
			outarray = append(outarray, v)
		}
	}
	return outarray
}

// My Sieve of Erathostenes algorithm
func MySieve (upToN int) []int {
	// Creates a pool of all the integers from 0 to upToN
	var pool = []int{}
	for x:=0;x<=upToN;x++ {
		pool = append(pool, x)
	}

	// Ranging from 2 to the the last integer, if the value is not already 0
	// sets all the multiples of that value to 0, up to the biggest multiple below the last integer.
	for _,p := range pool[2:] {
		if p != 0 {
			var divisor = 2
			for p*divisor <= upToN {
				pool[p*divisor] = 0
				divisor++
			}
		}
	}

	// Collects all the values from pool which arer still different from 0
	var collector = []int{}
	for p := range pool {
		if pool[p] != 0{
			collector = append(collector, pool[p])
		}
	}
	// Returns these values from index 1 to the end: index 0 is always 1.
	return collector[1:]
}

