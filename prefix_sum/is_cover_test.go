package prefixsum

import "testing"

func Test_isCovered(t *testing.T) {
	type args struct {
		ranges [][]int
		left   int
		right  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{
				ranges: [][]int{{1, 2}, {3, 4}, {5, 6}},
				left:   2,
				right:  5,
			},
			want: true,
		},
		{
			name: "test02",
			args: args{
				ranges: [][]int{{1, 10}, {10, 20}},
				left:   21,
				right:  21,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCovered(tt.args.ranges, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("isCovered() = %v, want %v", got, tt.want)
			}
		})
	}
}
