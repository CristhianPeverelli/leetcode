package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func pop(list linkedList) int {
	value := list.head.item
	list.head = list.head.next
	return value
}

func valutate(exp string) int {
	values := strings.Split(exp, " ")
	var pila linkedList
	for _, c := range values {
		if c >= "0" && c <= "9" {
			n, _ := strconv.Atoi(c)
			pila = addNewNode(n, pila)
		} else {
			first := pop(pila)
			second := pop(pila)
			switch c {
			case "+":
				pila = addNewNode((first + second), pila)
			case "-":
				pila = addNewNode((first - second), pila)
			case "*":
				pila = addNewNode((first * second), pila)
			case "/":
				pila = addNewNode((first / second), pila)
			default:
				fmt.Println("Input error.")
			}
		}
	}
	return pila.head.item
}

func convert(exp string) (res string) {
	values := strings.Split(exp, " ")
	var pila linkedList
	for _, c := range values {
		if c >= "0" && c <= "9" {
			res += c + " "
		} else if c == ")" {
			op := pop(pila)
			res += string(op) + " "
		} else if c != ")" {
			val := 0
			switch c {
			case "+":
				val += 43
			case "-":
				val += 45
			case "*":
				val += 42
			case "/":
				val += 47
			default:
				fmt.Println("Input error.")
			}
			pila = addNewNode(val, pila)
		}
	}
	return
}

func main() {
	result := valutate("5 3 - 2 *")
	fmt.Println(result)
	result2 := convert("( ( 5 - 3 ) * 2 )")
	fmt.Println(result2)
}
