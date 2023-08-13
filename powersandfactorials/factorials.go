package powersandfactorials

func Factorial (n int) int {
	out := 1
	for a:=2;a<=n;a++ {
		out = out * a
	}
	return out
}