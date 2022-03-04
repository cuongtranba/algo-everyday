// A string is good if there are no repeated characters.

// Given a string s​​​​​, return the number of good substrings of length three in s​​​​​​.

// Note that if there are multiple occurrences of the same substring, every occurrence should be counted.

// A substring is a contiguous sequence of characters in a string.

package slidingwindow

import "testing"

//aababcabc
func Test_countGoodSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 4",
			args: args{
				s: "aababcabc",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodSubstrings(tt.args.s); got != tt.want {
				t.Errorf("countGoodSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
