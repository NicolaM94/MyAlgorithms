//Package used to perform operations with the matrices.
//One should create a matrix object first and then add the rows.
//Columns are not defined as a base property of the object, they are calculated as a method instead.
//This is to prevent the need of updating the columns each time one modifies the rows.
package matrix

import (
	"errors"
	"fmt"
	"myalgo/utils"
)

// Base matrix type
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
func (mx *matrix) Columns() (out [][]float32) {
	for n := 0; n < len(mx.rows[0]); n++ {
		var temp []float32
		for m := range mx.rows {
			temp = append(temp, mx.rows[m][n])
		}
		out = append(out, temp)
	}
	return
} 

func (mx *matrix) PrintMatrix() {
	for r := range mx.rows {
		fmt.Println(mx.rows[r])
	}
	fmt.Println()
}

// Returns the size of the matrix by rows and columns
func (mx *matrix) Size() (int,int) {
	var cols = mx.Columns()
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

// Returns the diagonal of the matrix.
// Raise an error if the matrix is not a square one, since it's impossibile to calculate the diagonal of a non-square matrix.
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

// Returns the submatrix starting from the element a(i,j) to the element a(s,t)
func (mx *matrix) Submatrix(i,j,s,t int) matrix {
	var newMatrix matrix
	for row := i; row <= s; row++ {
		newMatrix.AddRow(mx.rows[row][j:t+1])
	}
	return newMatrix
}

// Swaps the newmatrix in the mx matrix starting from the element a(i,j).
// Returns an error if the newmatrix won't fit in mx matrix.
func (mx *matrix) ReplaceSubMatrix(i,j int, newmatrix matrix) (out matrix, err error) {

	// Checks if the newmatrix fits in the mx one
	var mxsizerow, mxsizecols = mx.Size()
	var nmxsizerow, nmxsizecols = newmatrix.Size()
	if mxsizecols-j < nmxsizecols {
		return matrix{}, fmt.Errorf("the columns of the new matrix overflow the old ones")
	}
	if mxsizerow-i < nmxsizerow {
		return matrix{}, fmt.Errorf("the rows of the new matrix overflow the old ones")
	}

	// Creates a copy of the original matrix and starts replacing elements
	out = *mx 
	for row := range newmatrix.rows {
		for col := range newmatrix.rows[row] {
			out.rows[row+i][col+j] = newmatrix.rows[row][col]
		}
	}

	return out, nil
}

// Calculate the transposed matrix from mx in place
func (mx *matrix) Transpose () {
	var out matrix
	var cols [][]float32 = mx.Columns()

	for c := range cols {
		var temp []float32
		for el := range cols[c] {
			temp = append(temp, cols[c][el])
		}
		out.AddRow(temp)
	}
	mx = &out
}
 
// Sum the matrix b in place
func (mx *matrix) Sum (b matrix) {
	for i := range mx.rows {
		for j := range mx.rows[i] {
			mx.rows[i][j] = mx.rows[i][j]+b.rows[i][j]
		}
	}
}

func (mx *matrix) ScaleRowBy (alpha float32, rowindex int) {
	for r := range mx.rows[rowindex] {
		mx.rows[rowindex][r] = mx.rows[rowindex][r] * alpha
	}
}

func (mx *matrix) ScaleBy (alpha float32) {
	for r := range mx.rows {
		mx.ScaleRowBy(alpha, r)
	}
}