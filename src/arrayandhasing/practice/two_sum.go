package practice

import (
	"slices"
	"sort"
)

/*
Problem:
Ref: https://leetcode.com/problems/two-sum/
An array of integers
A target (int)
Find indices of 2 elements that have sum of target
Assume:
- Have exactly 1 solution
- 1 element is only used once
- result may be in any order

constraints:
- arr length: [2, 10^4]
- Element and target value: [-10^9, 10^9]
*/

/*
First approach:
Iterate all the elements of the arr and calculate the sum of it with the rest
Time complexity: O(n^2)
Space complexity: O(1)
*/
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return make([]int, 0)
}

/*
Second approach:
First, store the origin indices
Sort the arr
After that, we iterate all the elements of arr and use binary search to find the rest
Time complexity: O(nlogn + nlogn) ~ O(nlogn)
Space complexity: O(n)
*/
func twoSum2(nums []int, target int) []int {
	indices := make(map[int][]int)
	for i, v := range nums {
		is, exist := indices[v]
		if exist {
			is = append(is, i)
		} else {
			is = []int{i}
		}

		indices[v] = is
	}

	sort.Ints(nums)
	for i, v1 := range nums {
		_, exist := slices.BinarySearch(nums[i+1:], target-v1)
		if exist {
			if v1 == target-v1 {
				return indices[v1]
			} else {
				return []int{indices[v1][0], indices[target-v1][0]}
			}
		}
	}

	return make([]int, 0)
}

/*
Third approach:
First, store the origin indices
After that, we iterate over all the pairs in index map
For each key, we will find the index of the rest in map
Time complexity: O(n + n) ~ O(n)
Space complexity: O(n)

Improvement 1:
Instead of finding after storing indices of all elements of arr, we can find the rest element immediately at the first loop
Time complexity: O(n)
Space complexity: O(n)
*/
func twoSum(nums []int, target int) []int {
	indices := make(map[int]int)
	for i1, v1 := range nums {
		i2, exist := indices[target-v1]
		if exist {
			return []int{i1, i2}
		} else {
			indices[v1] = i1
		}
	}
	return make([]int, 0)
}
