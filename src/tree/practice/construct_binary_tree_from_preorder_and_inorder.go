package practice

import . "AlgoPractice/src/tree/structure"

/*
Problem:
Ref: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

Given:
- 2 integer arrays - preorder and inorder

Return:
- Construct the binary tree and return its root

Constraints:
- array length: [1, 3000]
- element value: [-3000, 3000]
- arrays contain only unique values
*/

/*
First approach:
First, I look at preorder array, the 1st element will be the root of tree.
After that, I find the position of root in inorder array.
All the elements in the left of this array will be the inorder array of root left and similar with the right
The preorder arrays of root left and right children have the same size with corresponding inorder arrays and put in series after root in preorder array

Time complexity: O(n)
Space complexity: O(n)
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	for i, v := range inorder {
		if v == root.Val {
			root.Left = buildTree(preorder[1:i+1], inorder[0:i])
			root.Right = buildTree(preorder[i+1:], inorder[i+1:])
			break
		}
	}

	return root
}
