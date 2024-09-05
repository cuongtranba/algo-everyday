package prefixsum

func subarraySum(nums []int, k int) int {
	var count int
	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// we need to find k = sum - x then we need to find x
		// x = sum - k
		if _, ok := prefixSum[sum-k]; ok {
			count += prefixSum[sum-k]
		}
		prefixSum[sum]++
	}
	return count
}
