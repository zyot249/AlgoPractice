package practice

/*
Problem:
Ref: https://leetcode.com/problems/coin-change/description/

Given:
- An integer array (coins) representing coins of different denominations
- An integer (amount)

Return:
- The fewest number of coins that you need to make up that amount.
- If there doesn't exist any combination => -1

Constraints:
- coins length: [1, 12]
- coins[i]: [1, 2^31 - 1]
- amount: [0, 10^4]
*/

/*
Solution:
We will calculate the min number of coin for number from 1 => amount.
To calculate the min, we choose the min value of all options

Time complexity: O(amount*n) n is the length of coins
Space complexity: O(amount)
*/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		res := -1
		for _, coin := range coins {
			if i < coin || dp[i-coin] == -1 {
				continue
			}

			if res == -1 {
				res = dp[i-coin] + 1
			} else {
				res = min(res, dp[i-coin]+1)
			}
		}

		dp[i] = res
	}

	return dp[amount]
}
