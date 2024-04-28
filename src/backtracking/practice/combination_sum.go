package practice

/*
Problem:
Ref: https://leetcode.com/problems/combination-sum/

Given:
- An array of distinct integers (candidates)
- A target integer (target)

Return:
- A list of all distinct combinations of candidates with sum of target

Constraints:
- candidates length: [1, 30]
- candidates[i]: [2, 40]
- target: [1, 40]
*/

/*
First approach:
For each element in candidates, I find all the combinations with sum = target - element value with the array start from current candidate
Base condition:
if target < value of element => return empty
if target == value of element => return it

Time complexity: O(2^t) where t is largest size of combination
*/

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int

	if target == 1 {
		return result
	}

	for i, candidate := range candidates {
		if candidate == target {
			result = append(result, []int{target})
		} else if candidate < target {
			coms := combinationSum(candidates[i:], target-candidate)
			for _, com := range coms {
				result = append(result, append([]int{candidate}, com...))
			}
		}
	}

	return result
}
