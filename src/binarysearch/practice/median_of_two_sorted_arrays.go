package practice

import "math"

/*
Problem:
Ref: https://leetcode.com/problems/median-of-two-sorted-arrays/description/

Given:
- 2 sorted integer arrays nums1 and nums2 with length of m and n respectively

Return:
- The median of 2 arrays

Constraints:
- Time complexity: O(log(m+n))
- m, n: [0, 1000]
- m+n: [1, 2000]
- nums1[i], nums2[i]: [-10^6, 10^6]
*/

/*
I can't solve it :((
Solution:
The median will split the merged array into 2 parts with the same size
So the mission is finding the left part of merged array in 2 original arrays.
We know the size of the left part, so we will use binary search in 1 of 2 original arrays to find the first part of it,
the rest will be auto calculated.
The division is satisfied when the element at the left of division in array A <= the element at the right of division in array B and vice versa.
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	half := total / 2

	arrA, arrB := nums1, nums2
	if len(arrA) > len(arrB) {
		arrA, arrB = arrB, arrA
	}

	l, r := 0, len(arrA)-1
	for {
		m := (l + r) >> 1 //used to avoid the calculation of -1/2 = 0 when l = 0, r = -1
		n := half - m - 2

		var aLeft, aRight, bLeft, bRight float64
		if m >= 0 {
			aLeft = float64(arrA[m])
		} else {
			aLeft = math.Inf(-1)
		}

		if m+1 < len(arrA) {
			aRight = float64(arrA[m+1])
		} else {
			aRight = math.Inf(1)
		}

		if n >= 0 {
			bLeft = float64(arrB[n])
		} else {
			bLeft = math.Inf(-1)
		}

		if n+1 < len(arrB) {
			bRight = float64(arrB[n+1])
		} else {
			bRight = math.Inf(1)
		}

		if aLeft <= bRight && bLeft <= aRight {
			if total%2 == 0 {
				return (max(aLeft, bLeft) + min(bRight, aRight)) / 2
			} else {
				return min(bRight, aRight)
			}
		} else if aLeft > bRight {
			r = m - 1
		} else {
			l = m + 1
		}
	}
}
