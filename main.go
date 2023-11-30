package main

import (
	"fmt"
	setstheory "myalgo/setsTheory"
)

func main() {
	A := []int{2, 5, 3, 0, 7, 1, 8}
	B := []int{6, 4, 0, 3, 9, 6, 5}

	fmt.Println(setstheory.Difference(A, B))

}
