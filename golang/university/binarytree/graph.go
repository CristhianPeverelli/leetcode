package main

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
}

type graph struct {
	n    int
	near []linkedList
}

func newGraph(n int) *graph {
	return &graph{n, make([]linkedList, n)}
}

func main() {

}
