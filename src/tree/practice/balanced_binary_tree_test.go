package practice

import "testing"
import . "AlgoPractice/src/tree/structure"

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{root: CreateBinaryTree([]*Integer{})},
			want: true,
		},
		{
			name: "test2",
			args: args{root: CreateBinaryTree([]*Integer{{3}, {9}, {20}, nil, nil, {15}, {7}})},
			want: true,
		},
		{
			name: "test3",
			args: args{root: CreateBinaryTree([]*Integer{{1}, {2}, {2}, {3}, {3}, nil, nil, {4}, {4}})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
