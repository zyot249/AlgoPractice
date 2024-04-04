package practice

/*
Problem:
Ref: https://leetcode.com/problems/container-with-most-water/

Given:
- An integer arr (height)
- Element at ith position in height arr is the height of the vertical line at ith position
- Assume that 2 vertical lines formed with x-axis a container

Return:
- The maximum amount of water can be stored in formed container

Constraints:
- arr length: [2, 10^5]
- height[i]: [0, 10^4]
*/

/*
First approach:
We will loop through all the containers formed with the width from 1 -> n - 1

Time complexity: O(n^2)
Space complexity: O(1)
*/

func maxArea1(height []int) int {
	maximumArea := 0
	for w := 1; w < len(height); w++ {
		for i := 0; i < len(height)-w; i++ {
			area := w * min(height[i], height[i+w])
			maximumArea = max(area, maximumArea)
		}
	}

	return maximumArea
}

/*
Second approach:
We let 2 pointers iterate the arr from left and right at the same time.
When a pointer goes, the width of container will decrease.
So, if we want to find a larger container, we must let the pointer at smaller line runs and finds a higher line

Time complexity: O(n)
Space complexity: O(1)
*/

func maxArea(height []int) int {
	maximumArea := 0
	l, r := 0, len(height)-1
	for l < r {
		if height[l] <= height[r] {
			h := height[l]
			area := h * (r - l)
			maximumArea = max(area, maximumArea)

			for l < r && height[l] <= h {
				l++
			}
		} else {
			h := height[r]
			area := h * (r - l)
			maximumArea = max(area, maximumArea)

			for l < r && height[r] <= h {
				r--
			}
		}
	}

	return maximumArea
}
