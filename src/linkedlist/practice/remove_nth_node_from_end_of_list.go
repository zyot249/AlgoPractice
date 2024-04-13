package practice

/*
Problem:
Ref: https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/

Given:
- A head of linked list

Return:
- Remove the nth node from end of list and return the head

Constraints:
- size of list: [1, 30]
- node val: [1, 100]
- n: [1, size of list]
*/

/*
First approach:
Firstly, I use an array to store all nodes of list
I will remove the nth node by searching in array

Time complexity: O(n)
Space complexity: O(n)
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var arr []*ListNode
	ptr := head
	for ptr != nil {
		arr = append(arr, ptr)
		ptr = ptr.Next
	}

	if len(arr) < 1 {
		return nil
	}

	index := len(arr) - n
	if index == 0 {
		return head.Next
	} else if index == len(arr)-1 {
		arr[index-1].Next = nil
		return head
	} else {
		arr[index-1].Next = arr[index+1]
		return head
	}
}

/*
Second approach:
Count total nodes of list
Loop to find the position of removing node
Remove it

Time complexity: O(n)
Space complexity: O(1)
*/
