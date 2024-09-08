package prefixsum

import "testing"

func Test_minStartValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 001",
			args: args{
				nums: []int{
					-3, 2, -3, 4, 2,
				},
			},
			want: 5,
		},
		// {
		// 	name: "test 002",
		// 	args: args{
		// 		nums: []int{
		// 			1, 2,
		// 		},
		// 	},
		// 	want: 1,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStartValue2(tt.args.nums); got != tt.want {
				t.Errorf("minStartValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
