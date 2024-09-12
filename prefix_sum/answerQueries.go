package prefixsum

import "sort"

func answerQueries(nums []int, queries []int) []int {
	result := make([]int, len(queries))
	sort.Ints(nums)
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}

	for i := 0; i < len(queries); i++ {
		left, right := 0, n-1
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] > queries[i] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		result[i] = left
	}
	return result
}
