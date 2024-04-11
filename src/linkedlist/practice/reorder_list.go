package practice

/*
Problem:
Ref: https://leetcode.com/problems/reorder-list/description/

Given:
- You are given the head of a singly linked-list. The list can be represented as:
- L0 → L1 → … → Ln - 1 → Ln

Return:
- Reorder the list to be on the following form:
- L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

Constraints:
- You may not modify the values in the list's nodes. Only nodes themselves may be changed.
- Length of list: [1, 5*10^4]
- Node value: [1, 1000]
*/

/*
First approach:
I use a pointer to traverse the list.
For each position, I use another pointer to find the last element in list and insert the last element between the element visited by the first pointer and its next element
If the last element is visited by first pointer, the algorithm will stop

Time complexity: O(n^2)
Space complexity: O(1)
*/

func reorderList1(head *ListNode) {
	curr := head

	for curr != nil {
		last := curr
		for last.Next != nil {
			if last.Next.Next == nil {
				temp := last.Next
				last.Next = nil
				last = temp
			} else {
				last = last.Next
			}
		}

		if curr == last {
			break
		} else {
			temp := curr.Next
			curr.Next = last
			last.Next = temp
			curr = temp
		}
	}
}

/*
Solution 1:
Add all node of list to array
Use 2 pointers to iterate the array from left and right side to form new list

Time complexity: O(n)
Space complexity: O(n)
*/

func reorderList2(head *ListNode) {
	curr := head

	var arr []*ListNode
	for curr != nil {
		arr = append(arr, curr)
		curr = curr.Next
	}

	res := &ListNode{}
	curr = res
	left, right := 0, len(arr)-1
	for left <= right {
		curr.Next = arr[left]
		curr = curr.Next
		if left != right {
			curr.Next = arr[right]
			curr = curr.Next
		}

		left++
		right--
	}

	curr.Next = nil
	head = res.Next
}

/*
Solution 2:
Use fast and slow pointers to find the middle of list
Reverse the second half and merge them together

Time complexity: O(n)
Space complexity: O(1)
*/

func reorderList(head *ListNode) {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	second := slow.Next
	slow.Next = nil

	var prev *ListNode
	for second != nil {
		temp := second.Next
		second.Next = prev
		prev = second
		second = temp
	}

	second = prev

	curr := head
	for curr.Next != nil {
		temp := curr.Next
		curr.Next = second
		second = second.Next
		curr.Next.Next = temp
		curr = temp
	}

	curr.Next = second
}
