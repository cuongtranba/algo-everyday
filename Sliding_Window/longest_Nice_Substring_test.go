// A string s is nice if, for every letter of the alphabet that s contains, it appears both in uppercase and lowercase.
// For example, "abABB" is nice because 'A' and 'a' appear, and 'B' and 'b' appear. However, "abA" is not because 'b' appears, but 'B' does not.
// Given a string s, return the longest substring of s that is nice.
// If there are multiple, return the substring of the earliest occurrence. If there are none, return an empty string.
package slidingwindow

import "testing"

func Test_longestNiceSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "should return aAa",
			args: args{
				s: "YazaAay",
			},
			want: "aAa",
		},
		{
			name: "should return dD",
			args: args{
				s: "dDzeE",
			},
			want: "dD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestNiceSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestNiceSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
