// Given an array of integers nums, half of the integers in nums are odd, and the other half are even.

// Sort the array so that whenever nums[i] is odd, i is odd, and whenever nums[i] is even, i is even.

// Return any answer array that satisfies this condition.

package twopointer

func sortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	for even < len(nums) && odd < len(nums) {
		if nums[even]%2 != 0 {
			nums[even], nums[odd] = nums[odd], nums[even]
			odd += 2
		} else {
			even += 2
		}
	}
	return nums
}
