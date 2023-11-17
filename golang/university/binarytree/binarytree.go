package main

import (
	"fmt"
)

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

func buildTreeRec(values []int, index int) (binary *treeNode) {
	if index > len(values)-1 {
		return nil
	}
	binary = newNode(values[index])
	binary.left = buildTreeRec(values, (2*index)+1)
	binary.right = buildTreeRec(values, (2*index)+2)
	return binary
}

func buildTree(values []int) (binary *tree) {
	binary = new(tree)
	binary.root = buildTreeRec(values, 0)
	return
}

func printPreOrder(btree *treeNode) {
	if btree != nil {
		fmt.Print(btree.item, " - ")
		printPreOrder(btree.left)
		printPreOrder(btree.right)
	}
}

func printInOrder(btree *treeNode) {
	if btree != nil {
		printInOrder(btree.left)
		fmt.Print(btree.item, " - ")
		printInOrder(btree.right)
	}
}

func printPostOrder(btree *treeNode) {
	if btree != nil {
		printPostOrder(btree.left)
		printPostOrder(btree.right)
		fmt.Print(btree.item, " - ")
	}
}

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
	printInOrder(t.root)
	fmt.Println()
	printPostOrder(t.root)
	fmt.Println()

	values := []int{69, 89, 28, 39, 66, 44, 12, 2, 71}
	tree := buildTree(values)
	printPreOrder(tree.root)
}
