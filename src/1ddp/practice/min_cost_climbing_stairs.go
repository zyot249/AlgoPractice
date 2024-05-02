package practice

/*
Problem:
Ref: https://leetcode.com/problems/min-cost-climbing-stairs/description/

Given:
- An integer array (cost) where cost[i] is the cost of ith step on a staircase
- You can either climb 1 or 2 steps once you pay the cost
- You can either start at the step with index 0 or 1

Return:
- Minimum cost to reach the top of the floor

Constraints:
- cost length: [2, 1000]
- cost[i]: [0, 999]
*/

/*
First approach:
I will calculate the minimum cost for climbing to 1th -> (n-1)th step
The result will be minimum of (cost to n-1 + cost at n-1) and (cost to n-2 + cost at n-2)

Time complexity: O(n)
Space complexity: O(n)
*/

func minCostClimbingStairs(cost []int) int {
	result := make([]int, len(cost)+1)
	result[0] = 0
	result[1] = 0
	for i := 2; i <= len(cost); i++ {
		result[i] = min(result[i-2]+cost[i-2], result[i-1]+cost[i-1])
	}

	return result[len(cost)]
}
