package practice

import (
	. "AlgoPractice/src/linkedlist/structure"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{l1: CreateLinkedList([]int{2, 4, 3}), l2: CreateLinkedList([]int{5, 6, 4})},
			want: CreateLinkedList([]int{7, 0, 8}),
		},
		{
			name: "test2",
			args: args{l1: CreateLinkedList([]int{0}), l2: CreateLinkedList([]int{0})},
			want: CreateLinkedList([]int{0}),
		},
		{
			name: "test3",
			args: args{l1: CreateLinkedList([]int{9, 9, 9, 9, 9, 9, 9}), l2: CreateLinkedList([]int{9, 9, 9, 9})},
			want: CreateLinkedList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
