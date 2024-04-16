package practice

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{head: CreateLinkedList([]int{1, 2, 3, 4, 5}), k: 2}, want: CreateLinkedList([]int{2, 1, 4, 3, 5})},
		{name: "test1", args: args{head: CreateLinkedList([]int{1, 2, 3, 4, 5}), k: 3}, want: CreateLinkedList([]int{3, 2, 1, 4, 5})},
		{name: "test1", args: args{head: CreateLinkedList([]int{1, 2}), k: 2}, want: CreateLinkedList([]int{2, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
