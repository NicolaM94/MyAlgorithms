package main

import (
	"fmt"
	"myalgo/matrix"
)

func main() {
	var row1 = []float32{3,3/4,-2,-5}
	var row2 = []float32{0,9,-4,-2}
	var row3 = []float32{1,0,2,1}
	var row4 = []float32{0,0,0,0}
	
	var mx = matrix.CreateMatrix()
	err := mx.AddRow(row1)
	if err != nil {
		panic(err)
	}
	err = mx.AddRow(row2)
	if err != nil {
		panic(err)
	}
	mx.AddRow(row3)
	mx.AddRow(row4)
	fmt.Println("Initial matrix")
	mx.Print()

	mx.RemoveColumn(2)
	
	fmt.Println("After operation")
	mx.Print()
	fmt.Println("Columns")
	fmt.Println(mx.Columns())
}
