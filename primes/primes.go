package primes

import (
	"fmt"
	"math"
)

// Validates a number as prime by checking all its odd divisors
func IsPrime (number int) bool {
	if number == 0 || number == 1 {
		return false
	}
	if number == 2 {
		return true
	}
	if number%2==0 {
		return false
	}
	for n:=3;n<(number/2)+1;n+=2 {
		if number%n == 0 {
			return false
		}
	}
	return true
}



// My Sieve of Erathostenes algorithm
func Sieve (upToN int) []int {
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

// Checks if a prime is truncable from the right, meaning removing one digit by one starting from the right.
// This simpy divides by ten and takes the integer part of the result, verifying it is prime.
// If it's not it returns false
func isRightTruncatable (n int) bool {
	for n > 10 {
		if !IsPrime(n) {
			return false
		}
		n = int(n/10)
	}
	return IsPrime(n)
}


// Checks if a prime is truncatable from the left, meaning removing one digit by one starting from the left.
// Uses fmt lib to calculate the length of the number.
// This calculates the highest power of 10 smaller than the number and then divides the number by it, taking the integer part left.
// The integer part is then multiplied by the power of 10 and subtracted from the original number.
// The process stops when the power reaches 0.
func isLeftTruncatable (n int) bool {
	var size = len(fmt.Sprint(n))-1
	for size > 0 {
		var power = math.Pow10(size)
		var remover = int(power) * int(n/int(power))
		var result = n - remover
		if !IsPrime(result) {
			return false
		}
		n = result
		size --
	}
	return true
}

// Does isRightTruncatable && isLeftTruncatable
func IsTruncatable (n int) bool {
	return isLeftTruncatable(n) && isRightTruncatable(n)
}