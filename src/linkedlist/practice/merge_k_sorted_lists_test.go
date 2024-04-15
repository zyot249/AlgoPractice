package practice

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{lists: []*ListNode{
				CreateLinkedList([]int{}),
			}},
			want: nil,
		},
		{
			name: "test2",
			args: args{lists: []*ListNode{}},
			want: nil,
		},
		{
			name: "test3",
			args: args{lists: []*ListNode{
				CreateLinkedList([]int{1, 4, 5}),
				CreateLinkedList([]int{1, 3, 4}),
				CreateLinkedList([]int{2, 6}),
			}},
			want: CreateLinkedList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
