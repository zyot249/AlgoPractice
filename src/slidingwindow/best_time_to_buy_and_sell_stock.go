package slidingwindow

/*
Problem:
Ref: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

Given:
- An integer array represents prices of a stock day by day

Return:
- Maximum profit that can be received by buying at a single day and selling it at different day in future
- If it cannot get any profit, return 0

Constraints:
- Arr length: [1, 10^5]
- Arr[i]: [0, 10^4]
*/

/*
First approach:
I will iterate all the elements in arr prices.
For each price at ith day, I will find the maximum profit can be got by iterating all the prices after ith day.

Time complexity: O(n^2)
Space complexity: O(1)
*/

func maxProfit1(prices []int) int {
	maxPro := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			maxPro = max(maxPro, profit)
		}
	}

	return maxPro
}

/*
Second approach:
I use 2 pointers to iterate all the elements in arr
Initially, they start at 0 and 1 and calculate the profit
To increase profit, I will move the right pointer to next element that has greater value and calculate profit at each element
until there doesn't exist any element that is greater than current element.
After that, I will move the left pointer to find next element that has smaller value and calculate profit until it meets the right one.
Then, if the rest of arr has more than 1 element, move the left and right pointers to next position.

Time complexity: O(n)
Space complexity: O(1)
*/

func maxProfit2(prices []int) int {
	maxPro := 0
	if len(prices) <= 1 {
		return maxPro
	}

	l, r := 0, 1
	for r < len(prices) {
		profit := prices[r] - prices[l]
		maxPro = max(maxPro, profit)

		for i := r + 1; i < len(prices); i++ {
			if prices[i] >= prices[r] {
				r = i
				profit = prices[r] - prices[l]
				maxPro = max(maxPro, profit)
			}
		}

		for i := l + 1; i <= r; i++ {
			if prices[i] <= prices[l] {
				l = i
				profit = prices[r] - prices[l]
				maxPro = max(maxPro, profit)
			}
		}

		l = r + 1
		r = l + 1
	}

	return maxPro
}

/*
Solution: (Variable size sliding window technique)
I use 2 pointers to iterate all the elements in arr
Initially, they start at 0 and 1 and calculate the profit
I let the right pointer move to the left side.
If profit > max, max is re-assign
If profit <= 0, left pointer will be moved to the right pointer position
*/

func maxProfit(prices []int) int {
	maxPro := 0
	l, r := 0, 1
	for r < len(prices) {
		if prices[r] > prices[l] {
			profit := prices[r] - prices[l]
			maxPro = max(maxPro, profit)
		} else {
			l = r
		}

		r++
	}

	return maxPro
}
