package practice

/*
Problem:
Ref: https://leetcode.com/problems/house-robber-ii/description/

Given:
- An integer array (nums) representing the amount of money of each house.
- You cannot rob 2 adjacent houses on the same night because of security systems
- The first house is the neighbor of the last house

Return:
- Maximum amount of money you can rob tonight without alerting the police

Constraints:
- nums length: [1, 100]
- nums[i]: [0, 1000]
*/

/*
First approach:
I will split the problem into 2 cases:
- Rob the first house
- Don't rob the first house
The task after splitting is the house robber with houses not in cycle

Time complexity: O(2n)
Space complexity: O(2n)
*/

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return max(rob1(nums[:len(nums)-1]), rob1(nums[1:]))
}
