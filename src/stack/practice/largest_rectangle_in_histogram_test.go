package practice

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{heights: []int{2, 1, 5, 6, 2, 3}}, want: 10},
		{name: "test2", args: args{heights: []int{2, 4}}, want: 4},
		{name: "test3", args: args{heights: []int{1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
