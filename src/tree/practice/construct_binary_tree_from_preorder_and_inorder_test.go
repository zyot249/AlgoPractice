package practice

import (
	"reflect"
	"testing"
)
import . "AlgoPractice/src/tree/structure"

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: CreateBinaryTree([]*Integer{{3}, {9}, {20}, nil, nil, {15}, {7}}),
		},
		{
			name: "test2",
			args: args{
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: CreateBinaryTree([]*Integer{{-1}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
