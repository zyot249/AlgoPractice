package dailytest

import "sort"

/*
Problem:
Ref: https://leetcode.com/problems/4sum/description/

Given:
- Integer array of size n (nums)
- A target

Return:
- All quadruplets (nums[a], nums[b], nums[c], nums[d]) having sum of target

Constraints:
- n: [1, 200]
- nums: [-10^9, 10^9]
- target: [-10^9, 10^9]
- a, b, c, d: [0, n]
- a, b, c, d are distinct
*/

/*
Sort the array
For each index a in array, find all the triplet (b, c, d) in the array starting from a+1
For finding all the triplet, for each index b, find all the couple (c, d) in the array starting from b+1

Time complexity: O(n^3)
Space complexity: O(n)
*/
func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 4 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			sum := target - nums[i] - nums[j]
			l, r := j+1, len(nums)-1
			for l < r {
				if nums[l]+nums[r] > sum {
					r--
				} else if nums[l]+nums[r] < sum {
					l++
				} else {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})

					l++
					for nums[l] == nums[l-1] && l < r {
						l++
					}
				}
			}
		}
	}

	return result
}
