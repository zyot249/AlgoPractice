package practice

import . "AlgoPractice/src/tree/structure"

/*
Problem:
Ref: https://leetcode.com/problems/diameter-of-binary-tree/description/

Given:
- A root of a binary tree
- The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
This path may or may not pass through the root.

Return:
- The length of diameter of the binary tree

Constraints:
- node number: [1, 10^4]
- node value: [-100, 100]
*/

/*
First approach:
Length of longest path through each node is the sum of max depth of left child and max depth of right child
I will traverse the tree in post order

Time complexity: O(n)
Space complexity: O(n)
*/

func diameterOfBinaryTree(root *TreeNode) int {
	diameter, _ := longestPath(root)
	return diameter
}

func longestPath(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftDiameter, leftDepth := longestPath(root.Left)
	rightDiameter, rightDepth := longestPath(root.Right)
	return max(leftDepth+rightDepth, leftDiameter, rightDiameter), max(leftDepth, rightDepth) + 1
}
