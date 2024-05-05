package practice

/*
Problem:
Ref: https://leetcode.com/problems/maximum-product-subarray/description/

Given:
- An integer array (nums)

Return:
- The largest product of a subarray

Constraints:
- nums length: [1, 2*10^4]
- nums[i]: [-10, 10]
- The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
*/

/*
Solution:
For each element in array:
We will calculate the min and max product of a subarray ending with this element.
It will be this element or it with previous result

Time complexity: O(n)
*/

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res, curMax, curMin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(num, curMax*num)
		curMin = min(num, curMin*num)

		res = max(res, curMax)
	}

	return res
}
