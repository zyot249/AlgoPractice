package practice

import . "AlgoPractice/src/linkedlist/structure"

/*
Problem:
Ref: https://leetcode.com/problems/linked-list-cycle/description/

Given:
- A linked list

Return:
- True if the linked list has a cycle in it and false otherwise

Constraints:
- Node number: [0, 10^4]
- Node val: [-10^5, 10^5]
*/

/*
First approach:
I will use fast and slow pointers to iterate the linked list.
If 2 pointers meet each other before reaching end of list (nil), the list has a cycle.

Time complexity: O(n)
Space complexity: O(1)
*/

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head.Next, head
	for fast != nil {
		if fast == slow {
			return true
		}

		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next
	}

	return false
}
