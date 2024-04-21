package practice

import "math"
import . "AlgoPractice/src/tree/structure"

/*
Problem:
Ref: https://leetcode.com/problems/validate-binary-search-tree/description/

Given:
- a root of binary tree

Return:
- True if it is a binary search tree and false otherwise

Constraints:
- number of nodes: [1, 10^4]
- node val: [-2^31, 2^31 - 1]
*/

/*
First approach:
When I traverse the tree in preorder, I can find the range of possible value of next node.
Value of a node must be in range

Time complexity: O(n)
Space complexity: O(n)
*/

func isValidBST(root *TreeNode) bool {
	return checkBST(root, math.Inf(-1), math.Inf(1))
}

func checkBST(root *TreeNode, minVal float64, maxVal float64) bool {
	if root == nil {
		return true
	}

	if float64(root.Val) <= minVal || float64(root.Val) >= maxVal {
		return false
	}

	return checkBST(root.Left, minVal, float64(root.Val)) && checkBST(root.Right, float64(root.Val), maxVal)
}
