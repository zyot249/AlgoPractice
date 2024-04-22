package practice

import . "AlgoPractice/src/tree/structure"

/*
Problem:
Ref: https://leetcode.com/problems/binary-tree-maximum-path-sum/description/

Given:
- a root of binary tree
- A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them.
A node can only appear in the sequence at most once.
Note that the path does not need to pass through the root.
- The path sum of a path is the sum of the node's values in the path.

Return:
- the maximum path sum

Constraints:
- Number of nodes: [1, 3*10^4]
- Node value: [-1000, 1000]
*/

/*
First approach:
Max path sum of a tree is the max value of:
- root value
- max path sum of left child (if exist)
- max path sum of right child (if exist)
- root + max path can be extended from left
- root + max path can be extended from right
- root + max path can be extended from left and right

Time complexity: O(n)
Space complexity: O(n)
*/

func maxPathSum(root *TreeNode) int {
	result, _ := findPathSum(root)
	return result
}

func findPathSum(root *TreeNode) (int, int) {
	maxExtendPSum := root.Val
	maxPSum := root.Val

	if root.Left != nil && root.Right != nil {
		maxPSumLeft, maxExtendPSumLeft := findPathSum(root.Left)
		maxPSumRight, maxExtendPSumRight := findPathSum(root.Right)
		maxExtendPSum = max(maxExtendPSum, maxExtendPSumLeft+maxExtendPSum, maxExtendPSumRight+maxExtendPSum)
		maxPSum = max(maxPSum, maxPSumLeft, maxPSum+maxExtendPSumLeft, maxPSumRight, maxPSum+maxExtendPSumRight, maxPSum+maxExtendPSumLeft+maxExtendPSumRight)
	} else if root.Left != nil {
		maxPSumLeft, maxExtendPSumLeft := findPathSum(root.Left)
		maxExtendPSum = max(maxExtendPSum, maxExtendPSumLeft+maxExtendPSum)
		maxPSum = max(maxPSum, maxPSumLeft, maxPSum+maxExtendPSumLeft)
	} else if root.Right != nil {
		maxPSumRight, maxExtendPSumRight := findPathSum(root.Right)
		maxExtendPSum = max(maxExtendPSum, maxExtendPSumRight+maxExtendPSum)
		maxPSum = max(maxPSum, maxPSumRight, maxPSum+maxExtendPSumRight)
	}

	return maxPSum, maxExtendPSum
}
