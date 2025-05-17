package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	ans := []int{}
	queue := []*TreeNode{}

	queue = append(queue, root)
	for len(queue) > 0 {
		level := len(queue)
		cnt := 1
		for level > 0 {
			front := queue[0]
			queue = queue[1:] // deleting the front element
			if cnt == 1 {
				ans = append(ans, front.Val)
				cnt++
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			level--
		}
	}
	return ans
}
