package practice

/*
Problem:
Ref: https://leetcode.com/problems/reverse-nodes-in-k-group/description/

Given:
- The head of a linked list
- An integer k

Return:
- reverse the nodes of the list k at a time
- return the modified list

Constraints:
- node value: [0, 1000]
- list length: [1, 5000]
- k: [1, list length]
- left-out nodes, in the end, should remain as it is
*/

/*
First approach:
I think I will convert the list into an array and I can easily reverse each partition of length k

Time complexity: O(n + n) ~ O(n)
Space complexity: O(n)
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	var arr []*ListNode
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}

	var prev *ListNode = nil
	for i := 0; i < len(arr); i += k {
		newHead := arr[i]
		if i+k <= len(arr) {
			for j := i; j < i+k-1; j++ {
				arr[j+1].Next = arr[j]
			}

			newHead = arr[i+k-1]
		}

		if prev != nil {
			prev.Next = newHead
		} else {
			head = newHead
		}

		if i+k <= len(arr) {
			prev = arr[i]
		} else {
			prev = nil
		}
	}

	if prev != nil {
		prev.Next = nil
	}

	return head
}

/*
Second approach:
In first iteration, I will slit original list into sub-lists contain need reverse lists and remain lists
After that, I will reverse each list that need to be reversed.
Finally, I concatenate them together

Time complexity: O(n + n) ~ O(n)
Space complexity: O(n)
*/
