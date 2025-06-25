package arr

import "testing"

func Test_mergeAlternately(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// word1 = "abc", word2 = "pqr"
		{
			name: "abc-pqr",
			args: args{
				word1: "abc",
				word2: "pqr",
			},
			want: "apbqcr",
		},
		{
			name: "ab-pqrs",
			args: args{
				word1: "ab",
				word2: "pqrs",
			},
			want: "apbqrs",
		},
		{
			name: "cdf-a",
			args: args{
				word1: "cdf",
				word2: "a",
			},
			want: "cadf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeAlternately(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("mergeAlternately() = %v, want %v", got, tt.want)
			}
		})
	}
}
