
func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	if n%2 != 0 {
		return false
	}
	aux := 2
	for {
		if n == aux {
			return true
		}
		if aux > n {
			return false
		}
		aux *= 2
	}
}
