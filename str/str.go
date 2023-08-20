package str

// Insert element in starting at index, not in place
func Insert(element, starting string, index int) string {
	out := ""
	out += starting[:index]
	out += element
	out += starting[index:]
	return out
}
