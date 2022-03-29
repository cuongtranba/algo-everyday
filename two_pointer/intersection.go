package twopointer

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	seen := make(map[int]bool)

	for _, v := range nums1 {
		seen[v] = false
	}

	for _, v := range nums2 {
		_, ok := seen[v]
		if ok {
			seen[v] = true
		}
	}

	for k, v := range seen {
		if v {
			result = append(result, k)
		}
	}

	return result
}
