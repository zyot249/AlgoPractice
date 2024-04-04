package practice

/*
Problem:
Ref: https://leetcode.com/problems/trapping-rain-water/

Given:
- An non-negative integer arr represents an elevation map
- Element at ith position represents a bar with the height is the value of element and the width is 1

Return:
- The amount of trapped water after raining

Constraints:
- Arr length: [1, 2*10^4]
- height[i]: [0, 10^5]
*/

/*
First approach:
We let 2 pointers iterate the arr from the left and right.
The pointer at the shorter bar will be shifted to find the next higher bar until 2 pointer meet each other.
We will calculate the amount of water trapped when shifting pointer.

Time complexity: O(n)
Space complexity: O(1)
*/

func trap(height []int) int {
	amount := 0
	l, r := 0, len(height)-1
	for l < r {
		if height[l] <= height[r] {
			h := height[l]

			l++
			for l < r && height[l] <= h {
				amount += h - height[l]
				l++
			}
		} else {
			h := height[r]

			r--
			for l < r && height[r] <= h {
				amount += h - height[r]
				r--
			}
		}
	}

	return amount
}
