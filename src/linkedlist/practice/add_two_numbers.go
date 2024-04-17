package practice

import . "AlgoPractice/src/linkedlist/structure"

/*
Problem:
Ref: https://leetcode.com/problems/add-two-numbers/description/

Given:
- 2 linked lists that store the digits of 2 non-negative numbers in reverse order.

Return:
- Add 2 number and return the linked list that stores the result's digits in reverse order.

Constraints:
- length of list: [1, 100].
- It is guaranteed that the list represents a number that does not have leading zeros.
*/

/*
First approach:
I will iterate 2 lists at the same time and I will add 2 digits at the same position and store the result to response list
If sum of 2 digits is greater than 9, the digit in result will be sum mod 10 and bonus for next position will be sum / 10

Time complexity: O(n)
Space complexity: O(1)
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}

	ptr1, ptr2 := l1, l2
	cur := res
	bonus := 0
	for ptr1 != nil || ptr2 != nil {
		digit1, digit2 := 0, 0
		if ptr1 != nil {
			digit1 = ptr1.Val
			ptr1 = ptr1.Next
		}

		if ptr2 != nil {
			digit2 = ptr2.Val
			ptr2 = ptr2.Next
		}

		sum := digit1 + digit2 + bonus
		cur.Next = &ListNode{Val: sum % 10, Next: nil}
		bonus = sum / 10

		cur = cur.Next
	}

	if bonus > 0 {
		cur.Next = &ListNode{Val: bonus}
	}

	return res.Next
}
