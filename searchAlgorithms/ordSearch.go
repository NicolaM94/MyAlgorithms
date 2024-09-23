package searchalgorithms

// Binary search implementation which looks for elem in array
// Returns true if it finds it
// Iterative implementation
// Binary search
func BinarySearchIterative(array []int, target int) bool {
	var i int = 0
	var j int = len(array)-1
	for j>i {	
		var m int = (i+j)/2
		if array[m] < target {
			i = m+1
			continue
		}
		if array[m] > target {
			j = m-1
			continue
		}
		return true
	}
	return false
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
