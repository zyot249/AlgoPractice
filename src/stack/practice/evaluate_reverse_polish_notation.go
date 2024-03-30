package practice

import "strconv"

/*
Problem:
Ref:
- https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
- https://en.wikipedia.org/wiki/Reverse_Polish_notation

Given:
- An arr of strings that represents an arithmetic expression in Reverse Polish Notation

Return:
- The value of expression

Constraints:
- The expression is valid arithmetic
- All the calculations in 32-bit integer
- Arr length: [1, 10^4]
- arr[i]: [-200, 200] or [+, -, *, /]
*/

/*
First approach:
We will use a stack to store all visiting operands.
We iterate all the string in arr.
Each time we meet an operand, we push it to stack
If we meet an operator, we pop 2 operands from stack and do the calculation

Time complexity: O(n)
Space complexity: O(n)
*/

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		if token == "+" ||
			token == "-" ||
			token == "*" ||
			token == "/" {
			operand1 := stack[len(stack)-2]
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			value := doCalculation(token, operand1, operand2)

			stack = append(stack, value)
		} else {
			value, _ := strconv.Atoi(token)
			stack = append(stack, value)
		}
	}

	return stack[0]
}

func doCalculation(operator string, operand1 int, operand2 int) int {
	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2
	}

	return 0
}
