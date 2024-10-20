package utils


// Checks if an array is null, meaning that all its elements are 0s.
func IsNullArray (array []float32) bool {
	for a := range array {
		if array[a] != 0.0 {
			return false
		}
	}
	return true
}

// Returns the index of the first non null element
func FirstNonNullElement (array []float32) (out int) {
	for r := range array {
		if array[r] != 0 {
			out = r
		}
	}
	return
}

// Checks if arrays a and b are the same.
// 
// It also considers the position of the elements:
// e.g. [1,2,3] != [2,1,3]
func ArrayCompare (a,b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for m := 0; m < len(a); m++ {
		if a[m] != b[m] {
			return false
		}
	}
	return true
}

// Removes the index-th element from the given array
func ArrayRemove (array []float32, index int) []float32 {
	for i := index; i < len(array)-1; i++ {
		array[i] = array[i+1]
	}
	return array[:len(array)-1]
}

// Inserts an element at index index into the array
func ArrayInsert (array []float32, index int, element float32) []float32 {
	var out []float32
	out = array[:index]
	out = append(out, element)
	for i := range array[index+1:] {
		out = append(out, array[i])
	}
	return out
}
