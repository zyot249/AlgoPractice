package practice

import "sort"

/*
Problem:
Ref: https://leetcode.com/problems/combination-sum-ii/description/

Given:
- An integer array (candidates)
- An integer (target)

Return:
- All unique combinations that have sum of target

Constraints:
- candidates length: [1, 100]
- candidates[i]: [1, 50]
- target: [1, 30]
- each number in candidates must be used once in combinations
*/

/*
First approach:
I will find all the subsets of candidates having sum of target
Firstly, I sort the candidates.
Next, I iterate the sorted arr.
For each candidate:
if candidate == target => return
if candidate > target => find all combinations having sum of (target - candidate) in the subarray starting after current candidate

For candidates having the same value, I will handle only the first one.

Time complexity: O(nlogn + 2^t) where t is the longest length of combinations
*/

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int

	for i, candidate := range candidates {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}

		if candidate == target {
			ans = append(ans, []int{candidate})
		} else if candidate < target {
			coms := combinationSum2(candidates[i+1:], target-candidate)
			for _, com := range coms {
				ans = append(ans, append([]int{candidate}, com...))
			}
		}
	}

	return ans
}
