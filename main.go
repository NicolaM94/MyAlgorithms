package main

import (
	"fmt"
	complexnumbers "myalgo/complexNumbers"
)

func main() {
	a := complexnumbers.ComplexN{RealPart: 1, ImagPart: 1}
	b := complexnumbers.ComplexN{RealPart: 2, ImagPart: 2}
	a.PrintAlgebraic()
	fmt.Println()
	a.PrintTrigonometric()
	fmt.Println()
	b.PrintTrigonometric()
	
}
