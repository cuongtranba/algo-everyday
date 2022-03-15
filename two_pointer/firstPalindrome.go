// Given an array of strings words, return the first palindromic string in the array. If there is no such string, return an empty string "".

// A string is palindromic if it reads the same forward and backward.

package twopointer

func isPalindrome(word string) bool {
	left, right := 0, len(word)-1
	for left < right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func firstPalindrome(words []string) string {
	if len(words) == 0 {
		return ""
	}
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}
