package sommations

// Simple sum of an array of integers
func IntArraySum(array []int) (sum int) {
	for v := range array {
		sum += array[v]
	}
	return
}
