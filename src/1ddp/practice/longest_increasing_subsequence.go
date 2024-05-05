package practice

import "slices"

/*
Problem:
Ref: https://leetcode.com/problems/longest-increasing-subsequence/description/

Given:
- An integer array (nums)

Return:
- the length of longest strictly increasing subsequence

Constraints:
- length of nums: [1, 2500]
- nums[i]: [-10^4, 10^4]
*/

/*
First approach:
To find the longest length of subsequence starting from an element, I find the maximum of subsequence combined by this element with subsequences starting from the rest elements in array
Cache result for reusing

Time complexity: O(n^2)
Space complexity: O(n)
*/

func lengthOfLIS1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[n-1] = 1
	LIS := 1
	for i := n - 2; i >= 0; i-- {
		substringLengths := []int{1}
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				substringLengths = append(substringLengths, 1+dp[j])
			}
		}

		dp[i] = slices.Max(substringLengths)
		LIS = max(LIS, dp[i])
	}

	return LIS
}

/*
Solution:
Todo: Analyze later
*/

func lengthOfLIS(nums []int) int {
	tails := []int{nums[0]}

	for _, num := range nums {
		if num > tails[len(tails)-1] {
			tails = append(tails, num)
			continue
		}

		l, r := 0, len(tails)-1

		for l < r {
			md := l + (r-l)/2
			if tails[md] < num {
				l = md + 1
			} else {
				r = md
			}
		}
		tails[l] = num
	}
	return len(tails)
}
