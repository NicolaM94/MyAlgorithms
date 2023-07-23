package main

import (
	"fmt"
	orderingalgorithms "myalgo/orderingAlgorithms"
	
)

func main() {

	array := []int{8,5,2,3,1,9}

	fmt.Println(orderingalgorithms.SelectionSort(array))
}
