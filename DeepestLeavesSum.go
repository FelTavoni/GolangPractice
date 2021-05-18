/*
 * Created by Felipe Tavoni on Mon May 17 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Deepest Leaves Sum
 *
 * Given the root of a binary tree, return the sum of values of its deepest leaves.
 *
 * My solution: We depth-first search the binary tree. If we hit an end, check it's height. If is the biggest found, reset the sum.
 *	Otherwise, sum the value. By using global variables, we an make a better control of variables.
 *
 * (Not included a main function...)
 *
 */

package GolangPractice

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum, height, maxHeight int

func deepestLeavesSum(root *TreeNode) int {
	// Setting the variables to zero for a possible reuse.
	sum, height, maxHeight = 0, 0, 0
	// Call de dfs sum.
	auxDeepestLeavesSum(root)
	return sum
}

func auxDeepestLeavesSum(node *TreeNode) {
	// If the node is an end, reduce the height and go back.
	if node == nil {
		height--
		return
	}
	// Going to a lower level.
	height++
	auxDeepestLeavesSum(node.Left)
	// If it's height is higher, update it and reset sum.
	if height > maxHeight {
		sum = node.Val
		maxHeight = height
		// Otherwise, if it's the height we're used to, increase sum.
	} else if height == maxHeight {
		sum += node.Val
	}
	// Same for the right children.
	height++
	auxDeepestLeavesSum(node.Right)
	if height > maxHeight {
		sum = node.Val
		maxHeight = height
	}
	// Goes back.
	height--
}
