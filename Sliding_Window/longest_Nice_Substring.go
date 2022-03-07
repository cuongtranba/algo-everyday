// A string s is nice if, for every letter of the alphabet that s contains, it appears both in uppercase and lowercase.
// For example, "abABB" is nice because 'A' and 'a' appear, and 'B' and 'b' appear. However, "abA" is not because 'b' appears, but 'B' does not.
// Given a string s, return the longest substring of s that is nice.
// If there are multiple, return the substring of the earliest occurrence. If there are none, return an empty string.
package slidingwindow

import (
	"math"
	"strings"
)

// A = 65
// a = 97
// B = 66
// b = 98
// YazaAay
func longestNiceSubstring(s string) string {
	if len(s) == 1 {
		return ""
	}
	min := 0
	max := 1
	subStr := ""
	isHave := false
	for min < max && max < len(s) {
		if isHave {
			s1 := strings.ToUpper(string(s[max]))
			s2 := strings.ToUpper(string(s[max-1]))
			if s1 == s2 {
				subStr += string(s[max])
			} else {
				isHave = false
			}
		} else {
			sub := int(s[min]) - int(s[max])
			isUpper := math.Abs(float64(sub)) == 32
			if isUpper {
				subStr += s[min : max+1]
				isHave = true
			}
		}
		min++
		max++
	}

	if subStr == "" {
		return ""
	}

	return subStr

	m := []string{}
	current := ""
	mp := map[string]string{
		strings.ToLower(string(subStr[0])): string(subStr[0]),
	}
	for i := 0; i < len(subStr); i++ {
		_, ok := mp[strings.ToLower(string(subStr[i]))]
		if ok {
			current += string(subStr[i])
		} else {
			m = append(m, current)
			current = ""
			mp = map[string]string{
				strings.ToLower(string(subStr[i])): string(subStr[i]),
			}
		}
	}

	if current != "" {
		m = append(m, current)
	}

	currentMax := 0
	for _, v := range m {
		if len(v) > currentMax {
			currentMax = len(v)
		}
	}

	for _, v := range m {
		if len(v) == currentMax {
			return v
		}
	}

	return ""
}
