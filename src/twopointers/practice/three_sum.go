package practice

import "sort"

/*
Problem:
Ref: https://leetcode.com/problems/3sum/description/

Given:
- An array of integers (nums)

Return:
- all the triplets (nums[i], nums[j], nums[k]) (i,j,k are distinct, nums[i] + nums[j] + nums[k] = 0)

Constraints:
- arr length: [3, 3000]
- nums[i]: [-10^5, 10^5]
- The order of output is not necessary
- No duplication in result
*/

/*
First approach:
Firstly, we sort the arr in non-decreasing order.
We iterate all the elements in sorted arr.
For each element, we solve problem 2 sum in sorted array.
Note that we need to avoid handle first element with same value

Time complexity: O(n^2)
Space complexity: O(1)
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] < -nums[i] {
				l++
			} else if nums[l]+nums[r] > -nums[i] {
				r--
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})

				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}

	return result
}
