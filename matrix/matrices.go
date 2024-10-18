package matrix

import (
	"errors"
	"fmt"
	"myalgo/utils"
)

// Base matrix type
type matrix struct {
	rows [][]float32
	cols [][]float32
}

// Creates an empty matrix object
func CreateMatrix() matrix {
	return matrix{nil,nil}
}

func (mx *matrix) Rows() [][]float32 {
	return mx.rows
}

func (mx *matrix) Columns() [][]float32 {
	return mx.cols
}

func (mx *matrix) Print() {
	for r := range mx.rows {
		fmt.Println(mx.rows[r])
	}
}

// Returns the size of the matrix by rows and columns
func (mx *matrix) Size () (int,int) {
	return len(mx.rows),len(mx.cols)
}

// Appends the new row at the bottom of the matrix
func (mx *matrix) AddRow (newRow []float32) error {

	// If at least one row is present, must check for the correct lenght of the new row
	if mx.rows != nil && len(mx.rows[0]) != len(newRow) {
		return errors.New("error in adding a new row: incompatible row lenght")
	}

	// Adds the new row
	mx.rows = append(mx.rows, newRow)

	// Updates the columns with the new values from the row
	if mx.cols == nil {
		for n := range newRow {
			mx.cols = append(mx.cols, []float32{newRow[n]})
		}
	} else {
		for n := range mx.cols {
			mx.cols[n] = append(mx.cols[n], newRow[n])
		}
	}

	return nil
}

// Swaps two rows ins place
func (mx *matrix) SwapRows (i int,j int) {
	var temp = mx.rows[i]
	mx.rows[i] = mx.rows[j]
	mx.rows[j] = temp
}

// Inserts the newRow at index i
func (mx *matrix) InsertNewRow (index int, newRow []float32) {
	mx.AddRow(newRow)
	mx.SwapRows(index,len(mx.rows)-1)
}

// Removes the i-th row from the matrix, in place
func (mx *matrix) RemoveRow (i int) {
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
	// Updates the columns
	for m := range mx.cols {
		mx.cols[m] = utils.ArrayRemove(mx.cols[m],i)
	}
}

// Removes the j-th column from the matrix, in place
func (mx *matrix) RemoveColumn (j int) {
	// Update the columns
	for c := j+1; c < len(mx.cols); c++ {
		mx.cols[c-1] = mx.cols[c]
	}
	mx.cols = mx.cols[:len(mx.cols)-1]
	//Update the rows
	for m := range mx.rows {
		mx.rows[m] = utils.ArrayRemove(mx.rows[m],j)
	}
}