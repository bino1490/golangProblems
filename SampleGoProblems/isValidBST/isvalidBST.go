package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}
	return validate(node.Left, min, &node.Val) &&
		validate(node.Right, &node.Val, max)
}

func main() {
	root := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(isValidBST(root)) // true
}
