package practice

/*
Problem:
Ref: https://leetcode.com/problems/find-the-duplicate-number/description/

Given:
- An array of integers has n + 1 numbers
- Each integer is in the range [1, n]
- There is only 1 repeated number

Return:
- The repeated number

Constraints:
- n: [1, 10^5]
- Space complexity: O(1)
- Cannot modify original arr
*/

/*
Solution:
Consider the array is a linked list, in which the value of each element is the index of next element.
Use fast and slow pointers technique to find the first meet element of slow and fast pointers in the list (slow and fast are start at the same position)
We can prove that if the list has cycle, 2 slow pointers start from the head of list and the first meet element of fast and slow pointers will meet each other at start of cycle.

Special case:
- If there exists an element that has its index and value of the same value:
	- If it is repeated at somewhere, the list will have a cycle of 1 element
	- Otherwise, it is not in the list (Don't care about it)
*/

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0
	for {
		slow2 = nums[slow2]
		slow = nums[slow]
		if slow == slow2 {
			return slow
		}
	}
}
