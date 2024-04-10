package slidingwindow

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{"ADOBECODEBANC", "ABC"}, want: "BANC"},
		{name: "test2", args: args{"a", "a"}, want: "a"},
		{name: "test3", args: args{"a", "aa"}, want: ""},
		{name: "test3", args: args{"cabefgecdaecf", "cae"}, want: "aec"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
