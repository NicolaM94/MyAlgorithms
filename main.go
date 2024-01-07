package main

import (
	"fmt"
	"myalgo/sets"
)

func main() {
	A := []int{1,2,3}
	B := [][]int{{1,4},{2,5},{3,9}}
	fmt.Println(sets.IsAPartition(A,B))
}
