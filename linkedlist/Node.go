package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	header *Node
	size   int
}

func NewLinkedList() *LinkedList {
	header := &Node{
		val:  -1,
		next: nil,
	}

	return &LinkedList{header: header, size: 0}
}

func (linkedList *LinkedList) Append(val int) {
	var newNode *Node = &Node{val, nil}
	curNode := linkedList.header
	for curNode.next != nil {
		curNode = curNode.next
	}
	curNode.next = newNode
	linkedList.size++
}

func main() {
	linkedList := NewLinkedList()
	linkedList.Append(10)
	linkedList.Append(20)
	fmt.Println(linkedList.header.next)
	fmt.Println(linkedList.header.next.next)
	fmt.Println("hi hans!")
}
