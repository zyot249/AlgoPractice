package practice

import "sort"

/*
Problem:
Ref: https://leetcode.com/problems/subsets-ii/description/

Given:
- An integer array (nums)
- Array can contain duplicate numbers

Return:
- all subsets

Constraints:
- nums length: [1, 10]
- nums[i]: [-10, 10]
*/

/*
First approach:
For each element in array, I find all the subsets of the array starting after it
Base condition:
if length of array == 0 => return empty arr as an element of result

checking existence before adding to res
*/

func subsetsWithDup1(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})
	if len(nums) == 0 {
		return res
	}

	for i, num := range nums {
		sets := subsetsWithDup(nums[i+1:])
		for _, set := range sets {
			subset := append([]int{num}, set...)
			if !existSet(res, subset) {
				res = append(res, subset)
			}
		}
	}

	return res
}

func existSet(res [][]int, set []int) bool {
	for _, s := range res {
		if compareSet(s, set) {
			return true
		}
	}

	return false
}

func compareSet(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	freqA := make(map[int]int)
	freqB := make(map[int]int)

	for i := 0; i < len(a); i++ {
		freqA[a[i]]++
		freqB[b[i]]++
	}

	for k, v := range freqA {
		if v != freqB[k] {
			return false
		}
	}

	return true
}

/*
Solution:
Firstly, we sort the array.
We iterate all the elements in sorted array
We will not handle for all elements having the same value after visiting element
*/

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	res = append(res, []int{})
	if len(nums) == 0 {
		return res
	}

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}

		sets := subsetsWithDup(nums[i+1:])
		for _, set := range sets {
			res = append(res, append([]int{num}, set...))
		}
	}

	return res
}
