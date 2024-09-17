package prefixsum

import "testing"

func Test_returnToBoundaryCount(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				nums: []int{2, 3, -5},
			},
			want: 1,
		},
		{
			name: "Test case 2",
			args: args{
				nums: []int{3, 2, -3, -4},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := returnToBoundaryCount(tt.args.nums); got != tt.want {
				t.Errorf("returnToBoundaryCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
