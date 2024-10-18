package permutations


func HeapPermutations (array []int, size int, collector [][]int) [][]int {
	if size == 1 {
		collector = append(collector, array)
		return collector
	}
	for i := 0; i < size; i++ {
		HeapPermutations(array,size-1,collector)
		if size%2==1 {
			var tmp = array[size-1]
			array[size-1] = array[0]
			array[0] = tmp
		} else {
			var tmp = array[size-1]
			array[size-1] = array[i]
			array[i] = tmp
		}
	}
	return collector
}