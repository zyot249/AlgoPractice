package practice

import (
	. "AlgoPractice/src/heap/structure"
	"container/heap"
)

/*
Problem:
Ref: https://leetcode.com/problems/last-stone-weight/description/

Given:
- An integer array (stones)
- stones[i] is the weight of ith stone
- At each turn, choose the 2 heaviest stones and smash them together.
Assume that their weights are x and y, and x <= y
If x == y, both stones are destroyed.
If x < y, stone of weight x is destroyed and stone of weight y has new weight of y - x
Repeat until there are at most 1 stone left

Return:
- The weight of last stone

Constraints:
- stones length: [1, 30]
- stones[i]: [1, 1000]
*/

/*
First approach:
I use a heap to find the 2 largest stones

Time complexity: O(nlogn)
Space complexity: O(n)
*/

func lastStoneWeight(stones []int) int {
	maxHeap := IntMaxHeap(stones)
	heap.Init(&maxHeap)
	for len(maxHeap) > 1 {
		stone1 := heap.Pop(&maxHeap).(int)
		stone2 := heap.Pop(&maxHeap).(int)
		if stone1 > stone2 {
			heap.Push(&maxHeap, stone1-stone2)
		} else if stone2 > stone1 {
			heap.Push(&maxHeap, stone2-stone1)
		}
	}

	if len(maxHeap) == 1 {
		return maxHeap[0]
	}

	return 0
}
