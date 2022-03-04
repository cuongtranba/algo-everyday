// A string is good if there are no repeated characters.

// Given a string s​​​​​, return the number of good substrings of length three in s​​​​​​.

// Note that if there are multiple occurrences of the same substring, every occurrence should be counted.

// A substring is a contiguous sequence of characters in a string.

package slidingwindow

func countGoodSubstrings(s string) int {
	result := 0

	for i := 2; i < len(s); i++ {
		if s[i-2] != s[i-1] && s[i-2] != s[i] && s[i-1] != s[i] {
			result++
		}
	}

	return result
}
