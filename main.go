package main

import (
	"fmt"
	numbertheory "myalgo/numberTheory"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(numbertheory.CreateTriangular(i))
	}
}
