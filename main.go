package main

import (
	
	"myalgo/matrix"
	
)

func main() {
	var row1 = []float32{0,1}
	var row2 = []float32{2,-1}
	var row3 = []float32{1,7}
	
	var mx = matrix.CreateMatrix()
	mx.Columns()
	mx.RemoveRow(0)
	mx.GaussianElimination()

}
