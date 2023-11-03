package main

import "fmt"

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
}

func newNode(value int) *listNode {
	node := new(listNode)
	node.item = value
	return node
}

func addNewNode(value int, list linkedList) linkedList {
	node := newNode(value)
	node.next = list.head
	list.head = node
	return list
}

func printList(list linkedList) {
	for list.head != nil {
		fmt.Print(list.head.item, " - ")
		list.head = list.head.next
	}
}

func searchList(value int, list linkedList) (trovato bool) {
	for list.head != nil {
		if list.head.item == value {
			trovato = true
			break
		}
		list.head = list.head.next
	}
	return
}

func removeItem(value int, list linkedList) {
	actual := list.head
	var previous *listNode = nil
	for actual != nil {
		if actual.item == value {
			if previous == nil {
				list.head = list.head.next
			} else {
				previous.next = actual.next
			}
			break
		}
		previous = actual
		actual = actual.next
	}
}

// func main() {
// 	var lista linkedList
// 	lista = addNewNode(2, lista)
// 	lista = addNewNode(8, lista)
// 	lista = addNewNode(3, lista)
// 	lista = addNewNode(6, lista)
// 	printList(lista)
// 	fmt.Println(searchList(4, lista))
// 	removeItem(2, lista)
// 	printList(lista)
// }
