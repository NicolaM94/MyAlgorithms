/*Package used to perform operations with the matrices.
One should create a matrix object first and then add the rows.
Columns are not defined as a base property of the object, they are calculated as a method instead.
This is to prevent the need of updating the columns each time one modifies the rows.*/
package matrix

import (
	"errors"
	"fmt"
	"myalgo/utils"
)

//Base matrix type
type matrix struct {
	rows [][]float32
}

// Creates an empty matrix object
func CreateMatrix() matrix {
	return matrix{nil}
}

func (mx *matrix) Rows() [][]float32 {
	return mx.rows
}

// Returns a new matrix composed by the columns of this matrix
func (mx *matrix) Columns() (out matrix) {
	for n := 0; n < len(mx.rows[0]); n++ {
		var temp []float32
		for m := range mx.rows {
			temp = append(temp, mx.rows[m][n])
		}
		out.AddRow(temp)
	}
	return
} 

func (mx *matrix) Print() {
	for r := range mx.rows {
		fmt.Println(mx.rows[r])
	}
}

// Returns the size of the matrix by rows and columns
func (mx *matrix) Size() (int,int) {
	var cols = mx.Columns().rows
	return len(mx.rows),len(cols)
}

// Appends the new row at the bottom of the matrix
func (mx *matrix) AddRow(newRow []float32) error {
	// If at least one row is present, must check for the correct lenght of the new row
	if mx.rows != nil && len(mx.rows[0]) != len(newRow) {
		return errors.New("error in adding a new row: incompatible row lenght")
	}
	// Adds the new row
	mx.rows = append(mx.rows, newRow)
	return nil
}

// Swaps two rows ins place
func (mx *matrix) SwapRows(i int,j int) {
	var temp = mx.rows[i]
	mx.rows[i] = mx.rows[j]
	mx.rows[j] = temp
}

// Inserts the newRow at index i
func (mx *matrix) InsertNewRow(index int, newRow []float32) {
	mx.AddRow(newRow)
	mx.SwapRows(index,len(mx.rows)-1)
}

// Removes the i-th row from the matrix, in place
func (mx *matrix) RemoveRow(i int) {
	// Updates the rows
	if i == 0 {
		mx.rows = mx.rows[1:]
	} else if i == len(mx.rows)-1 {
		mx.rows = mx.rows[:len(mx.rows)-1]
	} else {
		for row := i+1; row < len(mx.rows); row++ {
			mx.rows[row-1] = mx.rows[row]
		}
	}
	mx.rows = mx.rows[:len(mx.rows)-1]
}

// Removes the j-th column from the matrix, in place
func (mx *matrix) RemoveColumn(j int) {
	for m := range mx.rows {
		mx.rows[m] = utils.ArrayRemove(mx.rows[m],j)
	}
}

// Returns the submatrix starting from the element a(i,j) to the element a(s,t)
func (mx *matrix) Submatrix(i,j,s,t int) matrix {
	var newMatrix matrix
	for row := i; row <= s; row++ {
		newMatrix.AddRow(mx.rows[row][j:t+1])
	}
	return newMatrix
}

// Returns the diagonal of the matrix.
// Raise an error if the matrix is not a square one.
func (mx *matrix) Diagonal() ([]float32, error) {
	if len(mx.rows) != len(mx.rows[0]) {
		return nil, fmt.Errorf("cannot select the diagonal of a non square matrix")
	}
	var out []float32
	for m := range mx.rows {
		out = append(out, mx.rows[m][m])
	}
	return out, nil
}

//TODO: Here
func (mx *matrix) ReplaceSubMatrix(i,j int, newmatrix matrix) (matrix, error) {

	return *mx, nil
}