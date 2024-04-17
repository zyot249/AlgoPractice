package practice

import . "AlgoPractice/src/tree/structure"

/*
Problem:
Ref: https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

Given:
- A root of a binary tree

Return:
- Its maximum depth

Constraints:
- Number of node: [0, 10^4]
- Node value: [-100, 100]
*/

/*
First approach:
I will use the recursive DFS to calculate the depth of tree

Time complexity: O(n)
Space complexity: O(1)
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
