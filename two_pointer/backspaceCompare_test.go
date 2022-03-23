package twopointer

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "test01",
		// 	args: args{
		// 		s: "ab#c",
		// 		t: "ad#c",
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "test02",
		// 	args: args{
		// 		s: "ab##",
		// 		t: "c#d#",
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "test03",
		// 	args: args{
		// 		s: "a#c",
		// 		t: "b",
		// 	},
		// 	want: false,
		// },
		{
			name: "test04",
			args: args{
				s: "a##c",
				t: "#a#c",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
