package main

import (
	"fmt"
	"myalgo/equations"
)

func main() {
	a := "-3x+9x-5+2+x"
	fmt.Println(equations.LinearSolver(a))

}
