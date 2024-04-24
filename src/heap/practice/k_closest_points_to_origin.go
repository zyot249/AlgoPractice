package practice

import "container/heap"

/*
Problem:
Ref: https://leetcode.com/problems/k-closest-points-to-origin/description/

Given:
- An array of points
- points[i] = [xi, yi] that represents a point in X-Y plane
- An integer k

Return:
- the k closest points to the origin (0, 0)

Constraints:
- k <= length of points
- length of points: [1, 10^4]
- xi, yi: [-10^4, 10^4]
*/

/*
First approach:
I use a min heap to find the closest point and add it to result
Repeat until there are k points in result

Time complexity: O(nlogn)
Space complexity: O(n)
*/

type PointMinHeap [][]int

func (h *PointMinHeap) Len() int { return len(*h) }
func (h *PointMinHeap) Less(i, j int) bool {
	iDis := (*h)[i][0]*(*h)[i][0] + (*h)[i][1]*(*h)[i][1]
	jDis := (*h)[j][0]*(*h)[j][0] + (*h)[j][1]*(*h)[j][1]
	return iDis < jDis
}
func (h *PointMinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *PointMinHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *PointMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type PointMaxHeap [][]int

func (h *PointMaxHeap) Len() int { return len(*h) }
func (h *PointMaxHeap) Less(i, j int) bool {
	iDis := (*h)[i][0]*(*h)[i][0] + (*h)[i][1]*(*h)[i][1]
	jDis := (*h)[j][0]*(*h)[j][0] + (*h)[j][1]*(*h)[j][1]
	return iDis > jDis
}
func (h *PointMaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *PointMaxHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *PointMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	if k < len(points)/2 {
		var result [][]int
		minHeap := PointMinHeap(points)
		heap.Init(&minHeap)
		for len(result) < k {
			point := heap.Pop(&minHeap).([]int)
			result = append(result, point)
		}

		return result
	} else {
		maxHeap := PointMaxHeap(points)
		heap.Init(&maxHeap)
		for len(maxHeap) > k {
			heap.Pop(&maxHeap)
		}

		return maxHeap
	}
}
