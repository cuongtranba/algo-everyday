// Given an array of strings words, return the first palindromic string in the array. If there is no such string, return an empty string "".

// A string is palindromic if it reads the same forward and backward.

package twopointer

import "testing"

func Test_firstPalindrome(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ada",
			args: args{
				words: []string{
					"abc", "car", "ada", "racecar", "cool",
				},
			},
			want: "ada",
		},
		{
			name: "racecar",
			args: args{
				words: []string{
					"notapalindrome", "racecar",
				},
			},
			want: "racecar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstPalindrome(tt.args.words); got != tt.want {
				t.Errorf("firstPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
