package main

import "fmt"

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
}

type queue struct {
	linkedList
	last *listNode
}

func first(q queue) int {
	return q.head.item
}

func isEmpty(q queue) bool {
	return q.head == nil
}

func print(q queue) {
	local := q.head
	for local != nil {
		fmt.Print(local.item, " - ")
		local = local.next
	}
}

func enqueue(q queue, value int) queue {
	r := new(listNode)
	r.item = value
	r.next = nil
	if isEmpty(q) {
		q.head = r
		q.last = r
	} else {
		q.last.next = r
		q.last = r
	}
	return q
}

func dequeue(q queue) (queue, int) {
	if isEmpty(q) {
		return q, -1
	}
	dato := q.head.item
	if q.head.next == nil {
		q.head = nil
		q.last = nil
		return q, dato
	} else {
		q.head = q.head.next
		return q, dato
	}
}

/*
func main() {
	var q queue
	q = enqueue(q, 10)
	q = enqueue(q, 8)
	q = enqueue(q, 1)
	q = enqueue(q, 22)
	q = enqueue(q, 9)
	q = enqueue(q, 24)
	print(q)
	var valore int
	q, valore = dequeue(q)
	fmt.Println("\n", valore)
	print(q)
}
*/
