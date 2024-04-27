package structure

import (
	"container/heap"
	"math"
)

/*
Problem:
Ref: https://leetcode.com/problems/find-median-from-data-stream/description/

Design:
Implement MedianFinder class:
- MedianFinder():
Initializes the MedianFinder object.
- void addNum(int num):
Adds the integer num from the data stream to the data structure.
- double findMedian():
Returns the median of all elements so far.
Answers within 10^-5 of the actual answer will be accepted.

Constraints:
- num: [-10^5, 10^5]
- There will be at least one element in the data structure before calling findMedian.
- At most 5 * 10^4 calls will be made to addNum and findMedian.

*/

/*
First approach:
In order to find the median, I need to find the kth smallest and (k+1)th largest of the stream - k = data size / 2
I will use a min heap of size k+1 and a max heap of size k to find them
*/

type MedianFinder struct {
	Count   int
	MinHeap IntMinHeap
	MaxHeap IntMaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		Count:   0,
		MinHeap: make(IntMinHeap, 0),
		MaxHeap: make(IntMaxHeap, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.Count++
	if this.MinHeap.Len() > 0 && num > this.MinHeap[0] {
		heap.Push(&this.MinHeap, num)
		minHeapLimit := int(math.Ceil(float64(this.Count) / 2))
		if this.MinHeap.Len() > minHeapLimit {
			item := heap.Pop(&this.MinHeap)
			heap.Push(&this.MaxHeap, item)
		}
	} else if this.MaxHeap.Len() > 0 && num < this.MaxHeap[0] {
		heap.Push(&this.MaxHeap, num)
		maxHeapLimit := int(math.Floor(float64(this.Count) / 2))
		if this.MaxHeap.Len() > maxHeapLimit {
			item := heap.Pop(&this.MaxHeap)
			heap.Push(&this.MinHeap, item)
		}
	} else {
		heap.Push(&this.MinHeap, num)
		minHeapLimit := int(math.Ceil(float64(this.Count) / 2))
		if this.MinHeap.Len() > minHeapLimit {
			heap.Pop(&this.MinHeap)
		}

		heap.Push(&this.MaxHeap, num)
		maxHeapLimit := int(math.Floor(float64(this.Count) / 2))
		if this.MaxHeap.Len() > maxHeapLimit {
			heap.Pop(&this.MaxHeap)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Count == 0 {
		return 0
	}

	if this.Count%2 == 0 {
		return float64(this.MinHeap[0]+this.MaxHeap[0]) / 2
	} else {
		return float64(this.MinHeap[0])
	}
}
