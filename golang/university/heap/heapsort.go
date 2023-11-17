package main

type treeNode struct {
	left  *treeNode
	right *treeNode
	item  int
}

type tree struct {
	root *treeNode
}

func main() {

}

func rearrange(heap *treeNode) *tree {
	ok := false
	for !ok {
		ok = true
		val := heap.item
		node := maxNode(heap, val)
		if node != nil {
			heap, node = node, heap
		}
	}
}
