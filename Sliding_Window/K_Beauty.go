package slidingwindow

import "strconv"

// The k-beauty of an integer num is defined as the number of substrings of num when it is read as a string that meet the following conditions:

// It has a length of k.
// It is a divisor of num.
// Given integers num and k, return the k-beauty of num.

// Note:

// Leading zeros are allowed.
// 0 is not a divisor of any value.
// A substring is a contiguous sequence of characters in a string.

func divisorSubstrings(num int, k int) int {
	result, s := 0, strconv.Itoa(num)

	for i := 0; i < len(s); i++ {
		if k+i > len(s) {
			return result
		}
		sub := s[i : k+i]
		subNum, _ := strconv.Atoi(sub)
		if subNum > 0 && num%subNum == 0 {
			result++
		}
	}

	return result
}
