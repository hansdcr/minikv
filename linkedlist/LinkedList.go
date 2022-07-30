package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func NewNode(data int) *Node {
	return &Node{data: data, next: nil}
}

type LinkedList struct {
	header *Node
	size   int
}

func NewLinkedList() *LinkedList {
	header := &Node{
		data: -1,
		next: nil,
	}

	return &LinkedList{
		header: header,
		size:   0,
	}
}

//nil->0->1->nil
func (linkedlist *LinkedList) Add(data int) {
	newNode := NewNode(data)
	prev := linkedlist.header

	for prev.next != nil {
		prev = prev.next
	}
	prev.next = newNode
	linkedlist.size++
}

// data = 3
//nil->nil
// nil->0->1->4->nil
func (linkedlist *LinkedList) AddBySort(data int) {
	newNode := NewNode(data)
	prev := linkedlist.header
	if prev == nil {
		linkedlist.header = newNode
		linkedlist.size++
		return
	}

	for {
		if newNode.data > prev.next.data {
			prev = prev.next
		}

		if prev.next == nil {
			prev.next = newNode
			linkedlist.size++
			return
		}

		if newNode.data == prev.next.data {
			return
		}

		if newNode.data < prev.next.data {
			tmp := prev.next
			prev.next = newNode
			newNode.next = tmp
			linkedlist.size++
			return
		}

	}
}

func main() {
	list := NewLinkedList()
	list.Add(0)
	list.Add(1)
	list.Add(4)
	list.AddBySort(3)
	list.AddBySort(1)
	list.AddBySort(5)
	fmt.Println(list)
}
