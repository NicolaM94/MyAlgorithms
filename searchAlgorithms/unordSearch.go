package searchalgorithms

// Simple algorithm which searches itm in itms and returns true if present
func UnorderedSearch(itm int, itms []int) bool {
	for i := range itms {
		if itm == i {
			return true
		}
	}
	return false
}
