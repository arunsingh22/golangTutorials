package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{}
	levelAns := []int{}

	queue = append(queue, root)
	for len(queue) > 0 {
		qs := len(queue)
		levelAns = []int{}

		for qs > 0 {
			front := queue[0]
			queue = queue[1:] // deleting the front element

			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
			levelAns = append(levelAns, front.Val)
			qs--
		}
		ans = append(ans, levelAns)
	}
	return ans
}

func findBottomLeftValue(root *TreeNode) int {
	var ans int
	queue := []*TreeNode{}
	if root == nil {
		return ans
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		level := len(queue)
		cnt := 1
		for level > 0 {
			front := queue[0]
			queue = queue[1:] // deleting the front element
			if cnt == 1 {
				ans = front.Val
				cnt++
			}

			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}
	}
	return ans
}
