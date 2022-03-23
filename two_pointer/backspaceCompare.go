package twopointer

func backspaceCompare(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) {
		if s[i] == '#' {
			if i == 0 {
				s = s[i+1:]
			} else {
				s = s[:i-1] + s[i+1:]
				i--
			}
		} else {
			i++
		}
	}
	for j < len(t) {
		if t[j] == '#' {
			if j == 0 {
				t = t[j+1:]
			} else {
				t = t[:j-1] + t[j+1:]
				j--
			}
		} else {
			j++
		}
	}
	return s == t
}
