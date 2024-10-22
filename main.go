package main

import(
	"myalgo/matrix"
)

func main() {
	var row1 = []float32{3,1,3}
	var row2 = []float32{2,-1,4}
	var row3 = []float32{1,7,2}
	
	var mx = matrix.CreateMatrix()
	mx.AddRow(row1)
	mx.AddRow(row2)
	mx.AddRow(row3)
	
	a, err := mx.GaussianElimination()
	if err != nil {
		panic(err)
	}
	a.PrintMatrix()
}
