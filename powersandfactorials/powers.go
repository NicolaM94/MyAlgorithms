// Managing operations with powers
package powersandfactorials

// Returns b**e leveraging the repeated power process
func RepeatedPower (b,e int) (int) {
	if e <= 1 {
		return b
	}
	out := RepeatedPower(b,e/2)
	out = out * out
	if e % 2 == 1 {
		out = out * b
	}
	return out
}