package practice

/*
Problem:
Ref: https://leetcode.com/problems/house-robber/description/

Given:
- An integer array (nums) representing the amount of money of each house.
- You cannot rob 2 adjacent houses on the same night because of security systems

Return:
- Maximum amount of money you can rob tonight without alerting the police

Constraints:
- nums length: [1, 100]
- nums[i]: [0, 400]
*/

/*
First approach:
I will calculate maximum amount of money can be robbed with 1 -> (n - 1) houses
If there is 1 house => money in this house
If there are 2 houses => the greater amount of money
If there are n houses => the greater one between the amount robbed from n-1 houses and the amount robbed from n-2 houses added with the amount of money in nth house

Time complexity: O(n)
Space complexity: O(n)

Improvement 1:
Instead of storing all the result from 0 -> n, we can only store the last 2 results
*/

func rob1(nums []int) int {
	//if len(nums) == 1 {
	//	return nums[0]
	//} else if len(nums) == 2 {
	//	return max(nums[0], nums[1])
	//}
	//
	//result := make([]int, len(nums)+1)
	//result[0] = nums[0]
	//result[1] = max(nums[0], nums[1])
	//for i := 2; i < len(nums); i++ {
	//	result[i] = max(result[i-1], result[i-2]+nums[i])
	//}
	//
	//return result[len(nums)-1]

	r1, r2 := 0, 0
	for _, num := range nums {
		tmp := r2
		r2 = max(r1+num, r2)
		r1 = tmp
	}

	return r2
}
