package main

import (
	"fmt"
	"myalgo/permutations"
)

func main() {

	for i := 1; i <= 20; i++ {
		fmt.Println(permutations.NthPermutation(i, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	}

}
