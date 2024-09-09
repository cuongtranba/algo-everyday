package prefixsum

// Let's consider an array [1, 2, 3, 4, 5] and focus on the element 3 (at index 2) to understand how we calculate its contribution to all odd-length subarrays.

// (i + 1) - subarrays starting from the left:

// There are 3 ways to start a subarray that includes 3: [1,2,3], [2,3], and [3]
// This is equal to (2 + 1) = 3

// (n - i) - subarrays ending at the right:

// There are 3 ways to end a subarray that includes 3: [3], [3,4], and [3,4,5]
// This is equal to (5 - 2) = 3

// Their product (i + 1) * (n - i):

// This gives us all possible subarrays (both odd and even length) containing 3
// In this case, it's 3 * 3 = 9 subarrays:
// [3]
// [2,3], [3,4]
// [1,2,3], [2,3,4], [3,4,5]
// [1,2,3,4], [2,3,4,5]
// [1,2,3,4,5]

// Adding 1 and dividing by 2: ((i + 1) * (n - i) + 1) / 2

// This operation filters out the even-length subarrays
// (9 + 1) / 2 = 5
// Indeed, there are 5 odd-length subarrays containing 3:
// [3]
// [1,2,3], [2,3,4], [3,4,5]
// [1,2,3,4,5]

// Multiplying by the element's value:

// Each of these 5 odd-length subarrays will include the value 3
// So the total contribution of this element is 3 * 5 = 15

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	result := 0
	for i := 0; i < n; i++ {
		left := i + 1
		right := n - i
		totalSubarrays := left * right
		oddLengthSubarrays := (totalSubarrays + 1) / 2
		result += oddLengthSubarrays * arr[i]
	}
	return result
}
