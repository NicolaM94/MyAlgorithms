package primes

import (
	"fmt"
	"math"
)

func SieveOfErath (limit int) int {
	crosslimit := int(math.Sqrt(float64(limit)))
	sieve := make([]bool,limit)
	
	for n:=0;n<limit;n=n+2 {
		sieve[n] = true
	}

	for n:=3;n<crosslimit;n=n+2{
		if !sieve[n] {
			for m := n*n;m<=limit;m = m+2 {
				sieve[m] = true
			}
		}
	}

	println("code reached")
	fmt.Println(sieve)

	sum := 0
	for n:=0;2<limit;n++ {
		if !sieve[n] {
			sum = sum+n
		}
	}

	return sum

}