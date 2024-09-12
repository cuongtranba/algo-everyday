package prefixsum

func isCovered(ranges [][]int, left int, right int) bool {
	coverd := make(map[int]bool, left-right+1)

	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			coverd[i] = true
		}
	}
	for i := left; i <= right; i++ {
		if !coverd[i] {
			return false
		}
	}
	return true
}
