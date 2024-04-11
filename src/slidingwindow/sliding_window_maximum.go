package slidingwindow

import (
	"slices"
	"sort"
)

/*
Problem:
Ref: https://leetcode.com/problems/sliding-window-maximum/description/

Given:
- An integer array (nums)
- A sliding window of size (k) moves from very left to very right of array
- Only k numbers in window can be seen
- Window moves right by 1 position

Return:
- Max of window at each position

Constraints:
- array length: [1, 10^5]
- nums[i]: [-10^4, 10^4]
- k: [1, array length]
*/

/*
First approach:
Easiest, I can iterate all the numbers in the window to find the maximum
But I think this solution will be timeout when running in leetcode

Time complexity: O(n*k)
Space complexity: O(1)
*/

func maxSlidingWindow1(nums []int, k int) []int {
	var res []int

	for i := 0; i < len(nums)-k+1; i++ {
		maxVal := nums[i]
		for j := i + 1; j < i+k; j++ {
			if nums[j] > maxVal {
				maxVal = nums[j]
			}
		}
		res = append(res, maxVal)
	}

	return res
}

/*
Second approach:
Sort first window, each time the window moves, find the position of removed element and added element and then form new sorted array

Time complexity: O(nlogk + k + klogk)
Space complexity: O(k)
*/

func maxSlidingWindow2(nums []int, k int) []int {
	var res []int

	window := make([]int, k)
	for i := 0; i < k; i++ {
		window[i] = nums[i]
	}

	sort.Ints(window)

	res = append(res, window[k-1])
	left := nums[0]
	for i := 1; i < len(nums)-k+1; i++ {
		leftIndex := sort.SearchInts(window, left)
		window = slices.Delete[[]int](window, leftIndex, leftIndex+1)
		rightIndex := sort.SearchInts(window, nums[i+k-1])
		window = slices.Insert[[]int](window, rightIndex, nums[i+k-1])

		res = append(res, window[k-1])
		left = nums[i]
	}

	return res
}

/*
Solution:
Use decreasing deque:
- Each time an element is pushed to deque, pop every element has smaller value at the top of deque before pushing
- We push the index of element to deque only

Time complexity: O(n)
Space complexity: O(k)
*/

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	deque := make([]int, 0)
	for i := 0; i < k; i++ {
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}

	res = append(res, nums[deque[0]])

	for i := 1; i < len(nums)-k+1; i++ {
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i+k-1] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i+k-1)

		if deque[0] < i {
			deque = deque[1:]
		}

		res = append(res, nums[deque[0]])
	}

	return res
}
