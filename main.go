package main

import (
	"fmt"
	orderingalgorithms "myalgo/sorting"
)

func main() {

	array := []int{8, 5, 2, 3, 1, 9}

	println(len(array))

	fmt.Println(orderingalgorithms.BubbleSort(array))
}
