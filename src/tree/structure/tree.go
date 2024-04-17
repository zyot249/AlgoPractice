package structure

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateBinaryTree(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	root := createBinarySubTree(vals, 0)
	return root
}

func createBinarySubTree(vals []int, i int) *TreeNode {
	if i >= len(vals) {
		return nil
	}

	root := &TreeNode{Val: vals[i]}
	root.Left = createBinarySubTree(vals, 2*i+1)
	root.Right = createBinarySubTree(vals, 2*i+2)
	return root
}
