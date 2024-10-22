// Provides a convenient function to perform the Gaussian elimination on the given matrix
package matrix

import (
	"myalgo/utils"
)

// TODO: Non funziona se nella riga l'elemento ij è zero, perchè da n/0 -> inf
func recursiveElimination (m matrix, startIndex int) matrix {
	// Pivot the first row
	var prepivot float32 = m.rows[startIndex][startIndex]
	for entry := range m.rows[startIndex] {
		var newEntry = m.rows[startIndex][entry]
		m.rows[startIndex][entry] = newEntry/prepivot
	}

	// Returns if only one row was modified since that must be the last one
	if startIndex == len(m.rows)-1 {
		return m
	}

	// Works on other rows
	for row := startIndex+1; row < len(m.rows); row++ {
		var pivotal float32 = m.rows[row][startIndex]
		
		for s := range m.rows[row] {
			if s <= startIndex {
				m.rows[row][s] = 0
			} else {
				m.rows[row][s] = m.rows[row][s] - m.rows[startIndex][s] * pivotal
			}
			
		}
	}
	return recursiveElimination(m, startIndex+1)
}

// Returns the minimimum form of the matrix by applying the G.E. algorithm.
func (mx *matrix) GaussianElimination () (out matrix, err error) {

	out = *mx

	// Clear matrix from null rows
	var temp matrix
	for r := range out.rows {
		if !utils.IsNullArray(out.rows[r]) {
			temp.AddRow(out.rows[r])
		}
	}
	out = temp
	
	// Clear matrix from null columns
	for j,r := range out.Columns() {
		if utils.IsNullArray(r) {
			for m := range out.rows {
				out.rows[m] = utils.ArrayRemove(out.rows[m],j)
			}
		}
	}

	var a matrix = recursiveElimination(out,0)

	return a,nil
}

