package prefixsum

func minStartValue(nums []int) int {
	k := 1
	for !checkPrefixSum(nums, k) {
		k += 1
	}
	return k
}

func checkPrefixSum(nums []int, k int) bool {
	currentSum := k
	for _, v := range nums {
		currentSum += v
		if currentSum <= 0 {
			return false
		}
	}
	return true
}

func minStartValue2(nums []int) int {
	prefixSum, result := int(0), int(1)

	for _, num := range nums {
		prefixSum += num
		//start value + prefixsum >= 1 --> start value = 1 - prefixsum
		// then we must use max to ensure every step must be > 1
		result = max(result, 1-prefixSum)
	}

	return result
}
