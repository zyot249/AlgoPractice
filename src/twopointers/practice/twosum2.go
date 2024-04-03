package practice

import "slices"

/*
Problem:
Ref: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

Given:
- An 1-indexed arr of integers (numbers)
- numbers is sorted in non-decreasing order

Return:
- Indices of 2 elements whose sum equals to target.

Constraints:
- 1 <= index 1 < index 2 < length of arr
- Arr length: [2, 3 * 10^4]
- numbers[i] and target: [-1000, 1000]
- There is exactly one solution
- The space complexity must be O(1)
*/

/*
First approach:
We will iterate all the elements in arr.
For each element, we will calculate the value of the rest and search it in the arr using binary search

Time complexity: O(nlogn)
Space complexity: O(1)
*/

func twoSum1(numbers []int, target int) []int {
	for i1, v1 := range numbers {
		v2 := target - v1
		i2, exist := slices.BinarySearch(numbers[i1+1:], v2)
		if exist {
			return []int{i1 + 1, i1 + 1 + i2 + 1}
		}
	}

	return make([]int, 0)
}

/*
Second approach:
We will use 2 pointers for iterating elements in arr, 1 starts from the left and the rest starts from the right.
For each element visiting by left pointer, we will try to find the other by using right pointer
Since the arr is sorted, left and right pointers will always go in 1-direction

Time complexity: O(n)
Space complexity: O(1)
*/

func twoSum(numbers []int, target int) []int {
	i1, i2 := 0, len(numbers)-1
	for i1 < i2 {
		if numbers[i1]+numbers[i2] == target {
			return []int{i1 + 1, i2 + 1}
		} else if numbers[i1]+numbers[i2] < target {
			i1++
		} else {
			i2--
		}
	}

	return make([]int, 0)
}
