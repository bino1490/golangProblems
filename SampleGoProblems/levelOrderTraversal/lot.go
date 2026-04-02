package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		level := []int{}

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil},
	}

	fmt.Println(levelOrder(root)) // [[1] [2 3]]
}
