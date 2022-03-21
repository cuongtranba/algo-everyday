// Given a string s and a character c that occurs in s, return an array of integers answer where answer.length == s.length and answer[i] is the distance from index i to the closest occurrence of character c in s.

// The distance between two indices i and j is abs(i - j), where abs is the absolute value function.
package twopointer

func shortestToChar(s string, c byte) []int {
	res := make([]int, len(s))

	curr := len(s)
	// shortest distance from left
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			res[i] = 0
			curr = 1
		} else {
			res[i] = curr
			curr++
		}
	}

	curr = len(s)
	// shortest distance from right
	for i := len(s) - 1; i >= 0; i-- {
		if res[i] == 0 {
			curr = 1
		} else {
			res[i] = min(res[i], curr)
			curr++
		}
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
