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
		// Insertion sort based enqueue
		inserted := false
		if q.head.item > value {
			r.next = q.head
			q.head = r
			inserted = true
		}
		for !inserted && (q.head.next != nil) {
			if q.head.next.item > value {
				r.next = q.head.next
				q.head.next = r
				inserted = true
			}
			q.head = q.head.next
		}
		if !inserted {
			q.last.next = r.next
			q.last = r
		}
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

func bucketSort(values []int, queueQtn int) []int {
	list := make([]queue, queueQtn)
	for i := range list {
		list[i] = *new(queue)
	}
	for _, value := range values {
		key := value / 10
		list[key] = enqueue(list[key], value)
	}
	j := 0
	for i := 0; i < len(list); i++ {
		for !isEmpty(list[i]) {
			list[i], values[j] = dequeue(list[i])
			j++
		}
	}
	return values
}
