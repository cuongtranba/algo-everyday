package slidingwindow

func numOfSubarrays(arr []int, k int, threshold int) int {
	if len(arr) < k {
		return 0
	}
	var count int
	total := 0
	for i := 0; i < k; i++ {
		total += arr[i]
	}
	if (total / k) >= threshold {
		count++
	}
	windowsStart, windowsEnd := 0, k-1
	for windowsEnd < len(arr) {
		total -= arr[windowsStart]
		windowsStart++
		windowsEnd++
		if windowsEnd < len(arr) {
			total += arr[windowsEnd]
			if (total / k) >= threshold {
				count++
			}
		}
	}
	return count
}

func numOfSubarrays2(arr []int, k int, threshold int) int {
	sum := 0
	res := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	if float64(sum)/float64(k) >= float64(threshold) {
		res++
	}

	for i := k; i < len(arr); i++ {
		sum += arr[i] - arr[i-k]
		if float64(sum)/float64(k) >= float64(threshold) {
			res++
		}
	}
	return res
}
