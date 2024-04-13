package practice

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CreateLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{}
	ptr := head
	for _, v := range arr {
		ptr.Next = &ListNode{Val: v}
		ptr = ptr.Next
	}

	return head.Next
}
