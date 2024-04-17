package practice

import . "AlgoPractice/src/linkedlist/structure"

/*
Problem:
Ref: https://leetcode.com/problems/merge-two-sorted-lists/description/

Given:
- The heads of 2 sorted linked lists (list1) and (list2)

Return:
- Merge into 1 sorted list and return the head of it

Constraints:
- length of list: [0, 50]
- node val: [-100, 100]
- both lists are sorted in non-decreasing order
- The merged list should be made by splicing together 2 first lists
*/

/*
First approach:
Firstly, I choose the head of response list. It will be the head of list having the head with smaller element.
Then, I iterate 2 lists at the same time.
The node of the rest list will be merged to response list if its value is in range of visiting node of response list and its next one.
And move pointer otherwise

Time complexity: O(m + n)
Space complexity: O(1)
*/

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var head, small, high *ListNode
	if list1.Val <= list2.Val {
		head = list1
		small, high = list1, list2
	} else {
		head = list2
		small, high = list2, list1
	}

	for high != nil {
		if high.Val >= small.Val {
			if small.Next == nil {
				small.Next = high
				high = nil
			} else {
				if high.Val <= small.Next.Val {
					smallNext := small.Next
					highNext := high.Next
					small.Next = high
					high.Next = smallNext
					high = highNext
				} else {
					small = small.Next
				}
			}
		}
	}

	return head
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	result := &ListNode{}
	curr := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
		if list1 == nil {
			curr.Next = list2
			list2 = nil
		} else if list2 == nil {
			curr.Next = list1
			list1 = nil
		}
	}

	return result.Next
}
