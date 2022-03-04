package slidingwindow

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 3",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "should return 1",
			args: args{
				s: " ",
			},
			want: 1,
		},
		{
			name: "should return 1",
			args: args{
				s: " ",
			},
			want: 1,
		},
		{
			name: "should return 3",
			args: args{
				s: "dvdf",
			},
			want: 3,
		},
		{
			name: "should return 3",
			args: args{
				s: "wobgrovw",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
