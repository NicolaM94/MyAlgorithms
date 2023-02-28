package main

import (
	"myalgo/fibonacci"
	orderedsearch "myalgo/orderedSearch"
)

func main() {
	println(fibonacci.FibonacciMath(10))
	println(fibonacci.FibonacciRec(10))
	println(fibonacci.FibonacciIter(10))
	println(orderedsearch.BinarySearchIter([]int{1, 2, 4, 5, 7, 8, 10, 12}, 9))
}
