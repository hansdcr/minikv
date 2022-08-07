package skiplist

import (
	"fmt"
	"strconv"
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

func (linkedlist *LinkedList) AddBySort(data int) {
	newNode := NewNode(data)
	prev := linkedlist.header

	for prev.next != nil {
		if newNode.data > prev.next.data {
			prev = prev.next
		} else if newNode.data == prev.next.data {
			return
		} else if newNode.data < prev.next.data {
			tmp := prev.next
			prev.next = newNode
			newNode.next = tmp
			linkedlist.size++
			return
		}
	}

	// 尾节点为空的情况
	prev.next = newNode
	linkedlist.size++
}

func (linkedlist *LinkedList) Search(data int) bool {
	header := linkedlist.header
	for header.next != nil {
		if header.next.data == data {
			return true
		}
		header = header.next
	}
	return false
}

func (linkedlist *LinkedList) Print() {
	header := linkedlist.header
	str := "nil"
	for header.next != nil {
		str = str + "--->" + strconv.Itoa(header.next.data)
		header = header.next
	}
	fmt.Printf(str)
}

func (linkedlist *LinkedList) Delete(data int) {
	prev := linkedlist.header

	// 0 --> 1---> 2 --->5
	for prev.next != nil {
		if prev.next.data == data {
			tmp := prev.next.next
			prev.next = tmp
			linkedlist.size--
			return
		}
		prev = prev.next
	}
}
