package practice

import (
	. "AlgoPractice/src/tree/structure"
	"math"
)

/*
Problem:
Ref: https://leetcode.com/problems/balanced-binary-tree/description/

Given:
- A root of a binary tree
- A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

Return:
- True if it is height balanced and false otherwise

Constraints:
- Number of nodes: [0, 5000]
- Node value: [-10^4, 10^4]
*/

/*
First approach:
I will calculate the depth of 2 children of every node and compare them together

Time complexity: O(n)
Space complexity: O(n)
*/

func isBalanced(root *TreeNode) bool {
	balanced, _ := checkBalancedAndGetDepth(root)
	return balanced
}

func checkBalancedAndGetDepth(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	isLeftBalanced, leftDepth := checkBalancedAndGetDepth(root.Left)
	isRightBalanced, rightDepth := checkBalancedAndGetDepth(root.Right)
	isRootBalanced := math.Abs(float64(leftDepth)-float64(rightDepth)) <= 1
	return isRootBalanced && isLeftBalanced && isRightBalanced, max(leftDepth, rightDepth) + 1
}
