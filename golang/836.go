package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	nodes := countNodes(root)
	m := make([][]int, nodes+1)
	for i := range m {
		m[i] = make([]int, nodes+1)
	}
	adjacencyMatrix(root, m)
	var adiacenti []int
	var result []int
	for i, _ := range m[target.Val] {
		if m[i][target.Val] == 1 {
			adiacenti = append(adiacenti, i)
		}
	}
	for j := 0; j < len(adiacenti); j++ {
		for i, _ := range m[adiacenti[j]] {
			if m[i][adiacenti[j]] == 1 && i != target.Val {
				result = append(result, i)
			}
		}
	}
	return result
}

func adjacencyMatrix(root *TreeNode, matrix [][]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		matrix[root.Val][root.Left.Val] = 1
		matrix[root.Left.Val][root.Val] = 1
		adjacencyMatrix(root.Left, matrix)
	}

	if root.Right != nil {
		matrix[root.Val][root.Right.Val] = 1
		matrix[root.Right.Val][root.Val] = 1
		adjacencyMatrix(root.Right, matrix)
	}
}

func countNodes(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + countNodes(node.Left) + countNodes(node.Right)
}

// func main() {
// 	root := &TreeNode{Val: 3}
// 	root.Left = &TreeNode{Val: 5}
// 	root.Right = &TreeNode{Val: 1}
// 	root.Left.Left = &TreeNode{Val: 6}
// 	root.Left.Right = &TreeNode{Val: 2}
// 	root.Right.Left = &TreeNode{Val: 0}
// 	root.Right.Right = &TreeNode{Val: 8}
// 	root.Left.Left.Left = &TreeNode{Val: 7}
// 	root.Left.Left.Right = &TreeNode{Val: 4}
// 	fmt.Println(distanceK(root, &TreeNode{Val: 5}, 2))
// }
