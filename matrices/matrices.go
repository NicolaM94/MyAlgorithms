package matrices

/**Package provides basic manipulation of matrices.*/

import (
	"errors"
	"fmt"
)

// Base matrix structure
type Matrix struct {
	Elements [][]int
}

// Prints the matrix rows
func (m *Matrix) PrintRows () {
	for e := range m.Elements {
		fmt.Println(m.Elements[e])
	}
}

// Prints the matrix columns
func (m *Matrix) PrintCols () {
	var counter = 0
	for counter < len(m.Elements[0]) {
		var collected []int
		for row := range m.Elements {
			collected = append(collected, m.Elements[row][counter])
		}
		fmt.Println(collected)
		counter++
	}
}

// Returns the dimensions of the matrix as (cols, rows)
func (m *Matrix) Dimensions() (int,int) {
	return len(m.Elements),len(m.Elements[0])
}

// Checks if the matrix is square
func (m *Matrix) IsSquare () bool {
	return len(m.Elements[0]) == len(m.Elements)
}

// Returns the diagonal of the matrix
func (m *Matrix) Diagonal () ([]int, error) {

	if !m.IsSquare() {
		return nil, errors.New("err: unable to calculate the diagonal: matrix is not squared")
	}

	var diagonal []int
	for i := 0; i < len(m.Elements); i++ {
		diagonal = append(diagonal, m.Elements[i][i])
	}
	return diagonal, nil
}

// TODO: Here
// Returns the determinant of the matrix
func (m *Matrix) Determinant () (int, error) {
	return 0, nil
}