package slidingwindow

func longestOnes(nums []int, k int) int {
	left, countZero, ans := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			countZero++
		}
		for countZero > k {
			if nums[left] == 0 {
				countZero--
			}
			left++
		}
		if tmp := right - left + 1; tmp > ans {
			ans = tmp
		}
	}
	return ans
}
