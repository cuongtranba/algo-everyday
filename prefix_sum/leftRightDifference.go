package prefixsum

func leftRightDifference(nums []int) []int {
	if len(nums) == 1 {
		return []int{0}
	}
	leftPrefixSum := make([]int, len(nums))
	rightPrefixSum := make([]int, len(nums))

	leftPrefixSum[0] = nums[0]
	rightPrefixSum[len(nums)-1] = nums[len(nums)-1]
	result := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		leftPrefixSum[i] = leftPrefixSum[i-1] + nums[i]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		rightPrefixSum[i] = rightPrefixSum[i+1] + nums[i]
	}

	for i := 0; i < len(nums); i++ {
		result[i] = abs(rightPrefixSum[i] - leftPrefixSum[i])
	}
	return result
}

func leftRightDifferencev2(nums []int) []int {
	if len(nums) == 1 {
		return []int{0}
	}
	result := make([]int, len(nums))
	sumleft, sumRight := 0, 0

	for i := 0; i < len(nums); i++ {
		sumRight += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		sumRight -= nums[i]
		result[i] = abs(sumRight - sumleft)
		sumleft += nums[i]
	}
	return result

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
