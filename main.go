package main

import (
	"fmt"
	"myalgo/sets"
)

func main() {
	A := []int{1,2,3}
	B := []int{1,2,3}
	fmt.Println(sets.CartesianProduct(A,B))
}
