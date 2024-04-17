package practice

import . "AlgoPractice/src/linkedlist/structure"

/*
Problem:
Ref: https://leetcode.com/problems/reverse-linked-list/description/

Given:
- A singly linked list

Return:
- Reversed list

Constraints:
- num of nodes: [0, 5000]
- node val: [-5000, 5000]

? A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

/*
First approach:
Iteratively
I will iterate all the nodes of list.
For each node, I will store the next of it.
Then, I assign the next of current node to previous node, the previous node to current node and the current node to next node

Time complexity: O(n)
Space complexity: O(1)
*/

func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

/*
Second approach:
I will use the same idea with first approach but implementing in recursive

Time complexity: O(n)
Space complexity: O(1)
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
