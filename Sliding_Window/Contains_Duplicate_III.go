package slidingwindow

import "math"

// Given an integer array nums and two integers k and t,
//return true if there are two distinct indices i and j in the array such that abs(nums[i] - nums[j]) <= t and abs(i - j) <= k.
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := range nums {
		for j := i + 1; j < len(nums) && j-i <= k; j++ {
			if int(math.Abs(float64(nums[i]-nums[j]))) <= t {
				return true
			}
		}
	}
	return false
}
