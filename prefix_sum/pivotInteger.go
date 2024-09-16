package prefixsum

func pivotInteger(n int) int {
	totalSum := n * (n + 1) / 2
	for i := range n {
		currentSum := i * (i + 1) / 2
		if totalSum-currentSum == currentSum+(i+1) {
			return i + 1
		}
	}
	return -1
}

func pivotInteger2(n int) int {
	totalSum := n * (n + 1) / 2
	left := 0
	right := n
	for left < right {
		mid := (left + right) / 2
		currentSum := mid * (mid + 1) / 2
		if totalSum-currentSum == currentSum+(mid+1) {
			return mid + 1
		} else if totalSum-currentSum > currentSum+(mid+1) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}
