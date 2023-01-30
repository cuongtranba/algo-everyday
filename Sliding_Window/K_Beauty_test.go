package slidingwindow

import "testing"

func Test_divisorSubstrings(t *testing.T) {
	type args struct {
		num int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase1",
			args: args{
				num: 430043,
				k:   2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorSubstrings(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("divisorSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
