package practice

/*
Problem:
Ref: https://leetcode.com/problems/product-of-array-except-self/
Given:
- An integer array: arr

Return:
- An integer array: answer
- answer[i] = product of all elements in arr except arr[i]

Constraints:
- Division operation is not allowed
- arr length: [2, 10^5]
- arr[i]: [-30, 30]
*/

/*
First approach:
Easiest, we loop all the elements in array.
For each element, we will multiply all the elements in answer array (except element at the same index with visiting element)
with value of visiting element

Time complexity: O(n^2)
Space complexity: O(n)
*/
func productExceptSelf1(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(result); i++ {
		result[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(result); j++ {
			if i == j {
				continue
			}
			result[j] *= nums[i]
		}
	}

	return result
}

/*
Second approach:
O(n^2) is too large, we need improve it to have O(n)
We need to calculate the prefix and postfix of each element in arr, in which:
Prefix means the product of all elements before visiting element
Postfix means the product of all elements after visiting element
The result will be the product of prefix and postfix for each element

[1,2,3,4]
-> [1,1,2,6]
-> [24,12,4,1]
-> [24,12,8,6]

Time complexity: O(n + n + n) ~ O(n)
Space complexity: O(n) -> O(1)
*/
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	//calculate prefix
	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	//calculate postfix
	postfix := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result
}
