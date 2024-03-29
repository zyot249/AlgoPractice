package practice

import stack2 "github.com/golang-collections/collections/stack"

/*
Problem:
Ref: https://leetcode.com/problems/valid-parentheses/
Given:
- A string contains only '(', ')', '{', '}', '[' and ']'

Return:
- input is valid or not

Constraints:
- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
- Every close bracket has a corresponding open bracket of the same type.
- string length: [1, 10^4]
*/

/*
First approach:
We will visit all the characters in string from start to end.
When we meet an open bracket, push it to stack.
When we meet a close bracket, pull a bracket from stack and compare. If there aren't exist any open bracket or pulled bracket is not same type -> invalid
*/
func isValid(s string) bool {
	stack := stack2.Stack{}
	for _, c := range s {
		if c == '{' || c == '(' || c == '[' {
			stack.Push(c)
		} else {
			open := stack.Pop()
			if open == nil ||
				(c == '}' && open != '{') ||
				(c == ')' && open != '(') ||
				(c == ']' && open != '[') {
				return false
			}
		}
	}

	if stack.Len() > 0 {
		return false
	}

	return true
}
