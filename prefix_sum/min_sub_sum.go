package prefixsum

import "math"

func minSubArrayLen(target int, nums []int) int {
	prefix := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	left := 0
	right := 1

	minLen := math.MaxInt

	for right < len(prefix) {
		diff := prefix[right] - prefix[left]
		if diff < target {
			right++
		} else {
			minLen = min(minLen, right-left)
			left++
		}
	}

	if minLen == math.MaxInt {
		return 0
	}

	return minLen
}
