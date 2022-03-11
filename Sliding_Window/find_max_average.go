package slidingwindow

import "math"

func findMaxAverage(nums []int, k int) float64 {
	currentTot := 0
	for i := 0; i < k; i++ {
		currentTot += nums[i]
	}
	currentMaxTot := currentTot

	windowStart, windowEnd := 0, k-1
	for windowEnd < len(nums) {
		currentTot -= nums[windowStart]

		// Slide window
		windowStart++
		windowEnd++

		if windowEnd < len(nums) {
			currentTot += nums[windowEnd]
			currentMaxTot = int(math.Max(float64(currentMaxTot), float64(currentTot)))
		}
	}

	return float64(currentMaxTot) / float64(k)
}
