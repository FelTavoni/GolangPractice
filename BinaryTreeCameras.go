/*
 * Created by Felipe Tavoni on Thu May 13 2021
 *
 * Problem proposal extracted from LeetCode©
 *
 * Binary Tree Cameras
 *
 * Given a binary tree, we install cameras on the nodes of the tree.
 * Each camera at a node can monitor its parent, itself, and its immediate children.
 * Calculate the minimum number of cameras needed to monitor all nodes of the tree.
 *
 * My solution: Very tough problem...At first, the ideia is to start from the bottom of the tree. Then we go way up, and verify the
 *	child nodes. If they are '0' it means there's no camera that have eyes on them. Otherwise, if it's '-1', it means that that spot
 *	needs a camera to monitorate it's children. At last, if there's a '1' as return, a camera has been positioned, so it will force
 *  the program to ignore a superior level, because it already has a camera.
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

func minCameraCover(root *TreeNode) int {
	var camera int
	if countCamera(root, &camera) == -1 {
		camera++
	}
	return camera

}
func countCamera(node *TreeNode, camera *int) int {
	if node == nil {
		return 0
	}
	left, right := countCamera(node.Left, camera), countCamera(node.Right, camera)
	if left == 0 && right == 0 {
		return -1
	}
	if left == -1 || right == -1 {
		*camera++
		return 1
	}
	return 0
}
