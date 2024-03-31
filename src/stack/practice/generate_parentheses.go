package practice

/*
Problem:
Ref: https://leetcode.com/problems/generate-parentheses/
Given:
- Num of pairs of parentheses

Return:
- All the combinations of well-formed parentheses

Constraints:
- Num of pairs: [1, 8]
*/

/*
First approach:
Since the num of pairs is not too big, we can prepare the results for each input.

Time complexity: O(1)
Space complexity: O(1)
*/

/*
Second approach:
We can generate all the distinct combinations formed from pairs of parentheses
We will generate combination by filling it from left to right.
Each time we fill '(', we increase count var by 1.
Each time we fill ')', we decrease count var by 1.
If count var equals to 0, don't fill ')'
*/

func generateParenthesis(n int) []string {
	combinations := genAllCombinations(n, 0, 0)
	return combinations
}

func genAllCombinations(pairNum int, left int, right int) []string {
	var result []string

	if left == pairNum && right == pairNum {
		result = append(result, "")
		return result
	}

	if left < pairNum {
		combinations := genAllCombinations(pairNum, left+1, right)
		for _, c := range combinations {
			result = append(result, "("+c)
		}
	}

	if right < left {
		combinations := genAllCombinations(pairNum, left, right+1)
		for _, c := range combinations {
			result = append(result, ")"+c)
		}
	}

	return result
}
