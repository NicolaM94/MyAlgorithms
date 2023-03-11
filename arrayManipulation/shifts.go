package arraymanipulation

// Shifts an array left to delete the first item
func InsertLeft(in []int) []int {
	for e := 1; e < len(in); e++ {
		in[e-1] = in[e]
	}
	return in[:len(in)-1]
}

// Shift an array right to delete the last item
func InsertRight(in []int) []int {
	for e := len(in) - 1; e >= 1; e-- {
		in[e] = in[e-1]
	}
	return in[1:]
}

// Delete the element at index i from the array
func DeleteIndex[T comparable](i int, array []T) []T {
	out := []T{}
	for n := 0; n <= i-1; n++ {
		out = append(out, array[n])
	}
	for n := i + 1; n < len(array); n++ {
		out = append(out, array[n])
	}
	return out
}

// Deletes the first occurence of the element found in the array
func DeleteElem[T comparable](s []T, k T) []T {
	for index := 0; index < len(s); index++ {
		if s[index] == k {
			return DeleteIndex(index, s)
		}
	}
	return s
}
