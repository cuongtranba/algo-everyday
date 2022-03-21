// Given a string s and a character c that occurs in s, return an array of integers answer where answer.length == s.length and answer[i] is the distance from index i to the closest occurrence of character c in s.

// The distance between two indices i and j is abs(i - j), where abs is the absolute value function.
package twopointer

import (
	"reflect"
	"testing"
)

func Test_shortestToChar(t *testing.T) {
	type args struct {
		s string
		c byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test01",
			args: args{
				s: "loveleetcode",
				c: 'e',
			},
			want: []int{
				3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestToChar(tt.args.s, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
