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

func (linkedList *LinkedList) AddTail(val int) {
	var newNode *Node = &Node{val, nil}
	curNode := linkedList.header
	for curNode.next != nil {
		curNode = curNode.next
	}
	curNode.next = newNode
	linkedList.size++
}

func (linkedList *LinkedList) AddHead(val int) {
	var newNode *Node = &Node{val, nil}
	curNode := linkedList.header
	tmp := curNode.next
	curNode.next = newNode
	newNode.next = tmp
	linkedList.size++
}

func (linkedList *LinkedList) Print() {
	curNode := linkedList.header
	for curNode.next != nil {
		fmt.Println(curNode.next.val)
		curNode = curNode.next
	}
}

func main() {
	linkedList := NewLinkedList()
	//linkedList.AddTail(10)
	//linkedList.AddTail(20)
	//linkedList.AddTail(30)
	//linkedList.AddTail(40)
	//linkedList.Print()

	linkedList.AddHead(10)
	linkedList.AddHead(20)
	linkedList.AddHead(30)
	linkedList.AddHead(40)
	linkedList.Print()

	fmt.Println("hi hans!")
}
