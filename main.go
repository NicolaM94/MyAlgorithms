package main

import (
	"fmt"
	"myalgo/linkedlists"
	numbertheory "myalgo/numberTheory"
	"os"
)

func main() {

	linkedList := linkedlists.LinkedList{}
	linkedList.Add(10)

	linkedList.Add(9)
	linkedList.Add(8)
	linkedList.Add(7)
	linkedList.Add(6)

	current := linkedList.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
	fmt.Println()

	linkedList.Remove(8)

	current = linkedList.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}

	linkedList.Print()
	fmt.Println()
	linkedList.Reverse()
	linkedList.Print()

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
