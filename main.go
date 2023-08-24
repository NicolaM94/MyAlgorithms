package main

import (
	"fmt"
	"myalgo/equations"
	"myalgo/ports"
)

func main() {
	a := "-3x+9x-5+2+x"
	fmt.Println(equations.LinearSolver(a))

	as := 13
	bs := 4
	print(ports.Xor(as, bs))

}
