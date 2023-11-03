package main

import "fmt"

type circNode struct {
	value int
	next  *circNode
	prev  *circNode
}

type circList struct {
	head *circNode
	tail *circNode
}

func addNewNodeC(value int, list circList) circList {
	node := new(circNode)
	node.value = value
	if list.tail == nil {
		list.tail = node
		list.head = node
	} else if list.tail.next == nil && list.head.next == nil {
		node.next = list.tail
		node.prev = list.tail
		list.head = node
	} else {
		node.next = list.head
		list.head.prev = node
		list.tail.next = node
		list.head = node
	}
	return list
}

func printFromZero(list circList) {
	for list.head.value != 0 {
		list.head = list.head.next
	}
	fmt.Print(list.head.value, " - ")
	list.head = list.head.next
	for list.head.value != 0 {
		fmt.Print(list.head.value, " - ")
		list.head = list.head.next
	}
}

func main() {
	var list circList
	list = addNewNodeC(2, list)
	list = addNewNodeC(8, list)
	list = addNewNodeC(3, list)
	list = addNewNodeC(9, list)
	list = addNewNodeC(6, list)
	list = addNewNodeC(0, list)
	list = addNewNodeC(1, list)
	printFromZero(list)
}
