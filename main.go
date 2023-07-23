package main

import (
	"fmt"
	orderingalgorithms "myalgo/sorting"
	
)

func main() {

	array := []int{8,5,2,3,2,1,9}

	fmt.Println("Selection sort")
	fmt.Println(orderingalgorithms.SelectionSort(array))

	fmt.Println("Insertion sort")
	fmt.Println(orderingalgorithms.InsertionSort(array))
}
