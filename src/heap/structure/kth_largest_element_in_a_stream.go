package structure

import "container/heap"

/*
Problem:
Ref: https://leetcode.com/problems/kth-largest-element-in-a-stream/description/

Design a class to find the kth largest element in a stream.
- KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
- int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.

Constraints:
- It is the kth largest element in the sorted order, not the kth distinct element.
- k: [1, 10^4]
- nums length: [0, 10^4]
- nums value: [-10^4, 10^4]
- At most 104 calls will be made to add.
- It is guaranteed that there will be at least k elements in the array when you search for the kth element.
*/

/*
Solution:
We will use a min heap of size K
*/

type KthLargest struct {
	K       int
	MinHeap *IntMinHeap
}

func Constructor(k int, nums []int) KthLargest {
	tmp := IntMinHeap(nums)
	this := KthLargest{k, &tmp}
	heap.Init(this.MinHeap)
	for len(*this.MinHeap) > k {
		heap.Pop(this.MinHeap)
	}
	return this
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.MinHeap, val)
	for len(*this.MinHeap) > this.K {
		heap.Pop(this.MinHeap)
	}
	return (*this.MinHeap)[0]
}
