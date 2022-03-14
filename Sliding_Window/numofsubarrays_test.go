package slidingwindow

import "testing"

func Test_numOfSubarrays(t *testing.T) {
	type args struct {
		arr       []int
		k         int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 3",
			args: args{
				arr:       []int{2, 2, 2, 2, 5, 5, 5, 8},
				k:         3,
				threshold: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfSubarrays(tt.args.arr, tt.args.k, tt.args.threshold); got != tt.want {
				t.Errorf("numOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
