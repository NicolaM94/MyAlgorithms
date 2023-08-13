package main

import (
	"fmt"
	arraymanipulation "myalgo/arrayManipulation"

)

func main() {

	a := []int{1,4,5,6,9,10}
	b := []int{2,3,4,5,9,13,0}

	fmt.Println(arraymanipulation.Intersection(a,b))

}
