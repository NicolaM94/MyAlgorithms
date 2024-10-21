package main

import (
	"myalgo/matrix"
)

func main() {
	var row1 = []float32{2,4,0,-2,-6}
	var row2 = []float32{0,8,0,-4,-2}
	var row3 = []float32{1,0,0,2,1}
	var row4 = []float32{0,0,0,0,0}
	
	var mx = matrix.CreateMatrix()
	mx.AddRow(row1)
	mx.AddRow(row2)
	mx.AddRow(row3)
	mx.AddRow(row4)

	
}
