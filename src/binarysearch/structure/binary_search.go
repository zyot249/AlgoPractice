package structure

/*
Problem:
Ref: https://leetcode.com/problems/binary-search/description/
Given:
- An integer arr sorted in ascending order
- A target number

Return:
- Implement binary search: if exist -> return index, otherwise return -1

Constraints:
- Arr length: [1, 10^4]
- arr[i], target: [-10^4, 10^4]
- All numbers in arr is unique
*/

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, l int, r int, target int) int {
	if l == r {
		if nums[l] == target {
			return l
		} else {
			return -1
		}
	}

	m := (l + r) / 2
	if nums[m] == target {
		return m
	} else if nums[m] < target {
		return binarySearch(nums, m+1, r, target)
	} else {
		return binarySearch(nums, l, m, target)
	}
}
