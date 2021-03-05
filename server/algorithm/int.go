package algorithm

func Rescuive(n int) int {
	if n == 0 {
		return 1
	}

	return n * Rescuive(n-1)
}
