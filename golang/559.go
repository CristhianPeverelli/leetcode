package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if len(root.Children) > 0 {
		max := 0
		for _, child := range root.Children {
			if maxDepth(child) > max {
				max = maxDepth(child)
			}
		}
		return 1 + max
	}
	return 1
}
