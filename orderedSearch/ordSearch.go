package orderedsearch

// Binary search implementation which looks for elem in array
// Returns true if it finds it
// Iterative implementation
func BinarySearchIter(array []int, elem int) bool {
	a := 1
	b := len(array)
	for array[int((a+b)/2)] != elem {
		m := int((a + b) / 2)
		if array[m] > elem {
			b = m - 1
		} else {
			a = m + 1
		}
		if a > b {
			return false
		}
	}
	return true
}

// Binary search implementation which looks for elem in array
// Returns true if it finds it
// Recursive implementation
func BinarySearchRec(array []int, elem, i, j int) bool {
	if i > j {
		return false
	}
	q := int((j - i + 1) / 2)
	if array[q] == elem {
		return true
	}
	if array[q] > elem {
		return BinarySearchRec(array, elem, i, q-1)
	}
	return BinarySearchRec(array, elem, q+1, j)
}
