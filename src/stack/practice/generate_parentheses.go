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
	open := 0
	combinations := genAllCombinations(open, n, n)
	return combinations
}

func genAllCombinations(open int, remainOpen int, remainClose int) []string {
	var result []string

	if remainOpen == 0 && remainClose == 0 {
		result = append(result, "")
		return result
	}

	if remainOpen > 0 {
		combinations := genAllCombinations(open+1, remainOpen-1, remainClose)
		for _, c := range combinations {
			result = append(result, "("+c)
		}
	}

	if remainClose > 0 && open > 0 {
		combinations := genAllCombinations(open-1, remainOpen, remainClose-1)
		for _, c := range combinations {
			result = append(result, ")"+c)
		}
	}

	return result
}
