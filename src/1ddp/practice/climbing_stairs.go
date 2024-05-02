package practice

/*
Problem:
Ref: https://leetcode.com/problems/climbing-stairs/

Given:
- You are climbing a staircase of n steps
- Each time, you can either climb 1 or 2 steps

Return:
- Return the number of distinct ways you can climb

Constraints:
- n: [1, 45]
*/

/*
First approach:
I will calculate the number of options for the first step and choose each of them.
After that, I will count the number of ways to finish the rest steps

Time complexity: O(1.618^n)
*/

func climbStairs1(n int) int {
	if n < 2 {
		return 1
	}

	return climbStairs1(n-1) + climbStairs1(n-2)
}

/*
Second approach:
Calculate all the result of climb from 1 to n - 1 steps
The result = climb(n-1) + climb(n-2)

Time complexity: O(n)
Space complexity: O(n)
*/

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}

	result := make([]int, n)
	result[0] = 1
	result[1] = 1
	for i := 2; i < n; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[n-1] + result[n-2]
}
