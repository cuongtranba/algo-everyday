package prefixsum

func returnToBoundaryCount(nums []int) int {
	position, count := 0, 0
	for _, v := range nums {
		position += v
		if position == 0 {
			count += 1
		}
	}
	return count
}
