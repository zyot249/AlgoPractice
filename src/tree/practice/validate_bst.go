package practice

import "math"
import . "AlgoPractice/src/tree/structure"

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
