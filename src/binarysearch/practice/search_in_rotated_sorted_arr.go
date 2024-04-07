package practice

/*
Problem:
Ref: https://leetcode.com/problems/search-in-rotated-sorted-array/description/

Given:
- A sorted integer arr (nums) that is possibly rotated at unknown pivot k (1 <= k < arr length)
- This arr is in form of nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] compared to original arr
- A target

Return:
- Index of target in arr if it exists and -1 otherwise

Constraints:
- Arr length: [1, 5000]
- Arr[i], target: [-10^4, 10^4]
- All elements in arr are unique
*/

/*
First approach:
Firstly, I will find the pivot using binary search.
Then, I will use the binary search based on the pivot.

Time complexity: O(logn + logn)
Space complexity: O(1)
*/

func search1(nums []int, target int) int {
	minPos := findMinPosition(nums)

	l, r := minPos, minPos+len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		i := m % len(nums)
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}

func findMinPosition(nums []int) int {
	minRotation, maxRotation := 1, len(nums)-1
	for minRotation <= maxRotation {
		mid := (minRotation + maxRotation) / 2
		if nums[mid] < nums[mid-1] {
			return mid
		}

		if nums[mid] > nums[len(nums)-1] {
			minRotation = mid + 1
		} else {
			maxRotation = mid - 1
		}
	}

	return 0
}

/*
Second approach:
I will use binary search with following conditions:
If mid > left => left arr is sorted, right arr is sorted otherwise
If target is in range of sorted arr => continue in it and the rest otherwise
*/

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}

		if nums[m] >= nums[l] {
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return -1
}
