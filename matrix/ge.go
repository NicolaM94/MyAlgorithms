// Provides a convenient function to perform the Gaussian elimination on the given matrix
package matrix

import (
	//"fmt"
	"fmt"
	"myalgo/utils"
)

func recursiveElimination (m matrix, startIndex int) matrix {
	fmt.Println("Stating index:", startIndex)
	out := m.Submatrix(startIndex,startIndex,len(m.rows)-1,len(m.rows[0])-1)
	fmt.Println("Considering matrix: ")
	
	out.PrintMatrix()

	if startIndex == len(m.rows) {
		return out
	}
	return recursiveElimination(out, startIndex+1)

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

	fmt.Println("Starting recursive elimination on matrix: ")
	out.PrintMatrix()
	recursiveElimination(out,0)


	return out,nil

}

