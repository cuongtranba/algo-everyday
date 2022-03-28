package twopointer

func removePalindromeSub(s string) int {
	total := 0
	if isPalindrome(s) {
		return 1
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
			continue
		}
		total += 2
		break
	}
	return total
}
