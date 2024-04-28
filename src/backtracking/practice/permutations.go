package practice

/*
Problem:
Ref: https://leetcode.com/problems/permutations/description/

Given:
- An array of distinct integers

Return:
- All possible permutations

Constraints:
- nums length: [1, 6]
- nums[i]: [-10, 10]
*/

/*
First approach:
For each element in array, I find all the permutations of the rest elements in array
Base condition:
if length of array = 1 => return it

Time complexity: O(n!)
*/

func permute(nums []int) [][]int {
	var result [][]int
	if len(nums) == 1 {
		result = append(result, []int{nums[0]})
		return result
	}

	for i, num := range nums {
		var rest []int
		for j := 0; j < len(nums); j++ {
			if j != i {
				rest = append(rest, nums[j])
			}
		}

		perms := permute(rest)
		for _, perm := range perms {
			result = append(result, append([]int{num}, perm...))
		}
	}

	return result
}
