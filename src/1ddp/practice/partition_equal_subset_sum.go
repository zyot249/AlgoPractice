package practice

/*
Problem:
Ref: https://leetcode.com/problems/partition-equal-subset-sum/description/

Given:
- An integer array (nums)

Return:
- True if you can partition the array into 2 subsets such that the sum of the elements in both subsets is equal
- False otherwise

Constraints:
- length of nums: [1, 200]
- nums[i]: [1, 100]
*/

/*
First approach:
Calculate the sum of all elements in array
Find a subset that has its sum of total sum / 2

Time complexity: O(n * sum(nums))
Space complexity: O(sum(nums))
*/

func canPartition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	if totalSum%2 == 1 {
		return false
	}

	targetSum := totalSum / 2
	for _, num := range nums {
		if num > targetSum {
			return false
		}
	}

	dp := make([]bool, targetSum+1)
	dp[0] = true
	for _, num := range nums {
		nextDp := make([]bool, targetSum+1)
		for sum, exist := range dp {
			if !exist {
				continue
			}

			if num+sum == targetSum {
				return true
			} else if num+sum < targetSum {
				nextDp[sum+num] = true
			}

			nextDp[sum] = true
		}

		dp = nextDp
	}

	return false
}
