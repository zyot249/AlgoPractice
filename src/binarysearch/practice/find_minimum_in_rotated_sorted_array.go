package practice

/*
Problem:
Ref: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/

Given:
- An integer array is sorted in ascending order
- It also was rotated some times
- Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]]

Return:
- The minimum element

Constraints:
- Time complexity: O(logn)
- Arr length: [1, 5000]
- Arr[i]: [-5000, 5000]
- All the elements in arr is unique
*/

/*
First approach:
I'm thinking about trying to find the number of rotations
I can use the binary search to find the number of rotations in range of [1, n]
With the number of rotations, I can easily find the minimum of arr.
So, the problem here is how I can determine that the number of rotations that is checking is right number.
I think that if the number of rotations (k) is less than the length of arr, arr[k] < arr[k-1].
If the element at position k is greater than the element at the end of arr, the number of rotations (k) must be greater and otherwise

Time complexity: O(logn)
Space complexity: O(1)
*/

func findMin(nums []int) int {
	minRotation, maxRotation := 1, len(nums)-1
	for minRotation <= maxRotation {
		mid := (minRotation + maxRotation) / 2
		if nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[mid] > nums[len(nums)-1] {
			minRotation = mid + 1
		} else {
			maxRotation = mid - 1
		}
	}

	return nums[0]
}
