package practice

import (
	. "AlgoPractice/src/heap/structure"
	"container/heap"
)

/*
Problem:
Ref: https://leetcode.com/problems/kth-largest-element-in-an-array/description/

Given:
- An integer array (nums)
- An integer k

Return:
- The kth largest element

Constraints:
- Note that it is the kth largest element in the sorted order, not the kth distinct element.
- solve it without sorting
- k <= nums.length
- nums.length: [1, 10^5]
- nums[i]: [-10^4, 10^4]
*/

/*
First approach:
I will heapify the array into a max heap
After that, I will pop the largest element from them heap until counter reaches k

Time complexity: O(klogn)
Space complexity: O(n)
*/

func findKthLargest(nums []int, k int) int {
	maxHeap := IntMaxHeap(nums)
	heap.Init(&maxHeap)
	count := 1
	for {
		if count == k {
			return heap.Pop(&maxHeap).(int)
		} else {
			heap.Pop(&maxHeap)
			count++
		}
	}
}
