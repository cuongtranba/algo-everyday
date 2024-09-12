package prefixsum

func findMiddleIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	prefixSumArr := makePrefixSumArr(nums)
	totalSum := prefixSumArr[len(prefixSumArr)-1]

	for i := 0; i < len(nums); i++ {
		leftSum := prefixSumArr[i]
		rightSum := totalSum - leftSum - nums[i]
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}
func makePrefixSumArr(arr []int) []int {
	n := len(arr)
	prefixSumArr := make([]int, n+1)
	prefixSumArr[0] = 0
	for i := 1; i <= n; i++ {
		prefixSumArr[i] = prefixSumArr[i-1] + arr[i-1]
	}
	return prefixSumArr
}
