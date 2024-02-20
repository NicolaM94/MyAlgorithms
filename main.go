package main

import (
	"fmt"
	"myalgo/arrayManipulation/sorting"
)

func main() {
	a := []int{7, 1, 9, 3, 5, 2}
	a = sorting.BubbleSort(a)
	fmt.Println(a)
}
