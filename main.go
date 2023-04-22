package main

import (
	"fmt"
	numbertheory "myalgo/numberTheory"
	"os"
)

func main() {

	triangularExceedingN(1, 200)

}

func triangularExceedingN(sp int, target int) {
	var divisors []int
	var a []int = []int{}

	for {
		println(sp)
		if sp%2 == 0 {
			a = numbertheory.FindDivisors(sp / 2)
			divisors = numbertheory.FindDivisors(sp + 1)
		} else {
			a = numbertheory.FindDivisors(sp + 1)
		}
		fmt.Printf("Len a: %v, len div: %v   SUM: %v\n", len(a), len(divisors), len(a)+len(divisors))
		if len(a)+len(divisors) > target {
			fmt.Printf("Found more than %v div on %v", target, sp)
			os.Exit(0)
		}
		sp++
	}

}
