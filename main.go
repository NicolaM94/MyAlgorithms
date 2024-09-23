package main

import (
	"fmt"
	"myalgo/ruffini"
)

func main() {

	var grades = []int{2,1,3,4}
	var equation = ruffini.Equation{Terms: grades}
	equation.PPrint()
	var out, div = equation.RuffinisRule()
	for out != nil {
		fmt.Println(out, div)
		equation.Terms = out
		out, div = equation.RuffinisRule()
	}
}
