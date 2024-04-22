package practice

import "testing"
import . "AlgoPractice/src/tree/structure"

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				root: CreateBinaryTree([]*Integer{{1}, {2}, {3}}),
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				root: CreateBinaryTree([]*Integer{{-10}, {9}, {20}, nil, nil, {15}, {7}}),
			},
			want: 42,
		},
		{
			name: "test3",
			args: args{
				root: CreateBinaryTree([]*Integer{{1}, {-2}, {-3}, {1}, {3}, {-2}, nil, {-1}}),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
