package main

import (
	"fmt"
	"myalgo/sets"
)

func main() {
	A := []int{1,2,3,7,8,9}
	B := []int{7,8,9}
	fmt.Println(sets.Difference(A,B))
}
