package ports

func And(a, b bool) bool {
	return a == b
}

func Or(a, b bool) bool {
	if a {
		return true
	}
	if b {
		return true
	}
	return false
}

func Not(a bool) bool {
	return !a
}
