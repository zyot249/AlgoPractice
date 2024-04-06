package practice

import (
	"math"
)

/*
Problem:
Ref: https://leetcode.com/problems/koko-eating-bananas/description/

Given:
- An integer arr represents the number of bananas in piles
- Koko has an eating speed of k bananas per hour
- Each hour, Koko will choose a pile of bananas to eat, and not eat another

Return:
- minimum of k such that Koko can eat all the piles of bananas in h hours

Constraints:
- num of piles: [1, 10^4]
- num of piles <= h <= 10^9
- piles[i]: [1, 10^9]
*/

/*
First approach:
We use the technique of binary search to search the speed in range of 1 to max bananas in pile
For find the min speed, when finding the speed at which Koko can eat all bananas with h hours, we continue finding in lower range

Time complexity: O(n + nlogm) where n is arr length, m is maximum value of element in arr
Space complexity: O(1)
*/

func minEatingSpeed(piles []int, h int) int {
	maxBananasInPile := 0
	for _, num := range piles {
		maxBananasInPile = max(num, maxBananasInPile)
	}

	l, r := 1, maxBananasInPile
	for l <= r {
		m := (l + r) / 2
		needHour := 0
		for _, num := range piles {
			needHour += int(math.Ceil(float64(num) / float64(m)))
		}
		if needHour <= h {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return l
}
