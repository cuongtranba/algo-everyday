package studyplan

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				s: "egg",
				t: "add",
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				s: "foo",
				t: "bar",
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				s: "paper",
				t: "title",
			},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{
				s: "ab",
				t: "aa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
