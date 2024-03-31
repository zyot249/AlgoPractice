package practice

/*
Problem:
Ref: https://leetcode.com/problems/daily-temperatures/
Given:
- An array of integers that represents the daily temperatures

Return:
- The array answer such that answer[i] is the number of days we need to wait to be warmer

Constraints:
- The arr length: [1, 10^5]
- Temperature: [30, 100]
*/

/*
First approach:
We can loop all the elements in arr
For each element, we will loop from its index to the end of arr to find out the element with higher value

Time complexity: O(n^2)
Space complexity: O(1)
*/

func dailyTemperatures1(temperatures []int) []int {
	result := make([]int, len(temperatures))
	for i, tem := range temperatures {
		waitDays := 0
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > tem {
				waitDays = j - i
				break
			}
		}
		result[i] = waitDays
	}

	return result
}

/*
Second approach:
We need to decrease the time complexity.

We loop the arr from the end to the start.
We will use a stack to store the temperature value and index of it
For each element, we will pop from stack to find out the warmer day.
If we can't find warmer day until the stack becomes empty, the result will be 0
Finally, we push this element with its index to stack

Time complexity: O(2n) ~ O(n)
Space complexity: O(n)
*/

type StackValue struct {
	val   int
	index int
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	var stack []StackValue
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 {
			top := stack[len(stack)-1]

			if top.val > temperatures[i] {
				result[i] = top.index - i
				break
			} else {
				stack = stack[:len(stack)-1]
			}
		}

		if len(stack) == 0 {
			result[i] = 0
		}

		stack = append(stack, StackValue{val: temperatures[i], index: i})
	}

	return result
}
