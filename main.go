package main

import (
	
	"myalgo/matrix"
	
)

func main() {
	var row1 = []float32{0,1}
	var row2 = []float32{2,-1}
	var row3 = []float32{1,7}
	
	var mx = matrix.CreateMatrix()
	mx.AddRow(row1)
	mx.AddRow(row2)
	mx.AddRow(row3)

	var mb = matrix.CreateMatrix()
	mb.AddRow([]float32{9,-1})
	mb.AddRow([]float32{3,-2})

	mx.MultiplyBy(mb)
	mx.PrintMatrix()

	
}
