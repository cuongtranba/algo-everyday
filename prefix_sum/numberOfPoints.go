package prefixsum

func numberOfPoints(nums [][]int) int {
	mPoints := make(map[int]struct{})
	for _, v := range nums {
		for i := v[0]; i <= v[1]; i++ {
			mPoints[i] = struct{}{}
		}
	}
	return len(mPoints)
}
