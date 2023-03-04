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
