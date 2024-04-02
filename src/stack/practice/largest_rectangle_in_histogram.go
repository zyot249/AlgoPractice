package practice

/*
Problem:
Ref: https://leetcode.com/problems/largest-rectangle-in-histogram/description/
Given:
- An arr of integers representing histogram's bar heights
- The width of histogram's bar is 1

Return:
- The area of largest rectangle in histogram

Constraints:
- arr length: [0, 10^4]
- height: [0, 10^5]
*/

/*
First approach:
Firstly, we iterate all the elements in arr.
For each element, we find the first elements that have smaller value than it at both left and right side and store their distance.
The largest rectangle formed at an element is established from the first smaller element at the left to the one at the right

Improvement 1:
When we find the first left element that has smaller value for all the elements in array,
the one that stands directly before it will be the first right element that has smaller value.

The largest rectangle formed at an element is established from the first smaller element at the left to the one at the right

Time complexity: O(n)
Space complexity: O(n)
*/

type Height struct {
	Val   int
	Index int
}

type IntStack struct {
	Stack []int
}

func (s *IntStack) push(val int) {
	s.Stack = append(s.Stack, val)
}

func (s *IntStack) pop() int {
	if len(s.Stack) == 0 {
		return -1
	}

	result := s.Stack[len(s.Stack)-1]
	s.Stack = s.Stack[:len(s.Stack)-1]
	return result
}

func (s *IntStack) peek() int {
	if len(s.Stack) == 0 {
		return -1
	}

	return s.Stack[len(s.Stack)-1]
}

func largestRectangleArea(heights []int) int {
	//var stack []Height
	//var rightSmallers = make([]int, len(heights))
	//for i := 0; i < len(heights); i++ {
	//	if len(stack) > 0 {
	//		for len(stack) > 0 {
	//			top := stack[len(stack)-1]
	//			if top.Val <= heights[i] {
	//				break
	//			}
	//
	//			rightSmallers[top.Index] = i - top.Index - 1
	//			stack = stack[:len(stack)-1]
	//		}
	//	}
	//
	//	stack = append(stack, Height{Val: heights[i], Index: i})
	//}
	//
	//for len(stack) > 0 {
	//	top := stack[len(stack)-1]
	//	rightSmallers[top.Index] = len(heights) - 1 - top.Index
	//	stack = stack[:len(stack)-1]
	//}
	//
	//var leftSmallers = make([]int, len(heights))
	//for i := len(heights) - 1; i >= 0; i-- {
	//	if len(stack) > 0 {
	//		for len(stack) > 0 {
	//			top := stack[len(stack)-1]
	//			if top.Val <= heights[i] {
	//				break
	//			}
	//
	//			leftSmallers[top.Index] = top.Index - i - 1
	//			stack = stack[:len(stack)-1]
	//		}
	//	}
	//
	//	stack = append(stack, Height{Val: heights[i], Index: i})
	//}
	//
	//for len(stack) > 0 {
	//	top := stack[len(stack)-1]
	//	leftSmallers[top.Index] = top.Index
	//	stack = stack[:len(stack)-1]
	//}
	//
	//maxArea := 0
	//for i, h := range heights {
	//	area := h * (1 + leftSmallers[i] + rightSmallers[i])
	//	if area > maxArea {
	//		maxArea = area
	//	}
	//}

	//Improvement
	var stack IntStack
	maxArea := 0
	for i, h := range heights {
		for stack.peek() != -1 && h <= heights[stack.peek()] {
			recH := heights[stack.pop()]
			recW := i - stack.peek() - 1
			maxArea = max(maxArea, recW*recH)
		}

		stack.push(i)
	}

	for stack.peek() != -1 {
		recH := heights[stack.pop()]
		recW := len(heights) - stack.peek() - 1
		maxArea = max(maxArea, recW*recH)
	}

	return maxArea
}
