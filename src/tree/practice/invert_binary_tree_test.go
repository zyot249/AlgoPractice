package practice

import (
	. "AlgoPractice/src/tree/structure"
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test1",
			args: args{root: CreateBinaryTree([]*Integer{{4}, {2}, {7}, {1}, {3}, {6}, {9}})},
			want: CreateBinaryTree([]*Integer{{4}, {7}, {2}, {9}, {6}, {3}, {1}}),
		},
		{
			name: "test2",
			args: args{root: CreateBinaryTree([]*Integer{{2}, {1}, {3}})},
			want: CreateBinaryTree([]*Integer{{2}, {3}, {1}}),
		},
		{
			name: "test3",
			args: args{root: CreateBinaryTree([]*Integer{})},
			want: CreateBinaryTree([]*Integer{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
