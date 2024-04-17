package practice

import . "AlgoPractice/src/tree/structure"

/*
Problem:
Ref: https://leetcode.com/problems/invert-binary-tree/description/

Given:
- A root of a binary tree

Return:
- Invert the tree and return its root

Constraints:
- number of nodes: [0, 100]
- node value: [-100, 100]
*/

/*
First approach:
I will swap the left child and right child for every node from root to leaves

Time complexity: O(n)
Space complexity: O(1)
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
