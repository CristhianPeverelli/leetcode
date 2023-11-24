package main

import "fmt"

type treeNode struct {
	left  *treeNode
	right *treeNode
	item  int
}

type tree struct {
	root *treeNode
}

func newNode(val int) *treeNode {
	return &treeNode{nil, nil, val}
}

func buildHeap(tree *treeNode) *treeNode {
	if tree != nil {
		if tree.left != nil {
			buildHeap(tree.left)
		}
		if tree.right != nil {
			buildHeap(tree.right)
		}
		rearrange(tree)
	}
	return tree
}

func printPreOrder(btree *treeNode) {
	if btree != nil {
		fmt.Print(btree.item, " - ")
		printPreOrder(btree.left)
		printPreOrder(btree.right)
	} else {
		fmt.Print("X - ")
	}
}

func printPostOrder(btree *treeNode) {
	if btree != nil {
		printPostOrder(btree.left)
		printPostOrder(btree.right)
		fmt.Print(btree.item, " - ")
	}
}

func rearrange(tree *treeNode) {
	ok := false
	if tree != nil {
		for !ok {
			if tree.left == nil && tree.right == nil {
				ok = true
			} else {
				val := tree.item
				node := maxNode(tree)
				if node != nil && node.item > val {
					tree.item, node.item = node.item, tree.item
					tree = node
				} else {
					ok = true
				}
			}
		}
	}
}

func maxNode(tree *treeNode) *treeNode {
	if tree.left == nil {
		return tree.right
	} else if tree.right == nil {
		return tree.left
	} else if tree.left.item > tree.right.item {
		return tree.left
	} else {
		return tree.right
	}
}

/*
func main() {
	t := &tree{nil}
	t.root = &treeNode{nil, nil, 78}
	t.root.left = newNode(54)
	t.root.right = newNode(21)
	t.root.left.right = newNode(90)
	t.root.left.right.left = newNode(19)
	t.root.left.right.right = newNode(95)
	t.root.right.left = newNode(16)
	t.root.right.left.left = newNode(5)
	t.root.right.right = newNode(19)
	t.root.right.right.left = newNode(56)
	t.root.right.right.right = newNode(43)

	printPreOrder(t.root)
	fmt.Println()
	t.root = buildHeap(t.root)
	printPreOrder(t.root)
}
*/
