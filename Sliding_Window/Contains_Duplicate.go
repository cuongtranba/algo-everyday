//Given an integer array nums and an integer k,
//return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

package slidingwindow

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && int(math.Abs(float64(i-j))) <= k {
				return true
			}
		}
	}
	return false
}

// apply sliding windows
func containsNearbyDuplicateV2(nums []int, k int) bool {
	hash := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			return true
		}
		hash[nums[i]] = true
		if len(hash) > k {
			delete(hash, nums[i-k])
		}
	}
	return false
}
