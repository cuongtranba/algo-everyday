package twopointer

func reverseWords(s string) string {
	if s == "" {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	newStr := ""

	left, right := 0, 0
	for right < len(s) {
		if string(s[right]) == " " {
			for i := right - 1; i >= left; i-- {
				newStr += string(s[i])
			}
			newStr += " "
			left = right + 1
		}
		right++
	}
	for left < right {
		newStr += string(s[right-1])
		right--
	}

	return newStr
}
